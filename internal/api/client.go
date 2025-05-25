package api

import (
	"context"
	"encoding/json"
	"fmt"
	"netmaker-sync/internal/config"
	"netmaker-sync/internal/models"
	"netmaker-sync/swagger"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// Client is a wrapper around the Netmaker API client
type Client struct {
	restClient    *resty.Client
	swaggerClient *swagger.APIClient
	config        *config.NetmakerAPIConfig
	ctx           context.Context
}

// New creates a new Netmaker API client
func New(cfg *config.NetmakerAPIConfig, loggingCfg *config.LoggingConfig) *Client {
	// Set up REST client for debugging
	restClient := resty.New()
	restClient.SetBaseURL(cfg.URL)
	restClient.SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.Key))
	restClient.SetTimeout(30 * time.Second)
	restClient.SetRetryCount(3)
	restClient.SetRetryWaitTime(5 * time.Second)
	restClient.SetRetryMaxWaitTime(20 * time.Second)

	// Only enable Resty debug mode if we're in debug log level AND disable_resty_debug is false
	isDebug := logrus.GetLevel() <= logrus.DebugLevel && !loggingCfg.DisableRestyDebug
	restClient.SetDebug(isDebug)

	// Log request and response details
	restClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		logrus.Debugf("API Request: %s %s", req.Method, req.URL)
		logrus.Debugf("API Request Headers: %v", req.Header)
		return nil
	})

	restClient.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		logrus.Debugf("API Response Status: %s", resp.Status())
		logrus.Debugf("API Response Headers: %v", resp.Header())
		logrus.Debugf("API Response Body: %s", resp.String())
		return nil
	})

	// Set up Swagger client
	swaggerConfig := swagger.NewConfiguration()
	swaggerConfig.BasePath = cfg.URL

	// Create an OAuth2 token source for authentication
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Key},
	)

	// Create an HTTP client with the token source
	oauth2Client := oauth2.NewClient(context.Background(), tokenSource)
	swaggerConfig.HTTPClient = oauth2Client

	// Create Swagger client
	swaggerClient := swagger.NewAPIClient(swaggerConfig)

	logrus.Infof("Initialized Swagger API client with base URL: %s", cfg.URL)

	return &Client{
		restClient:    restClient,
		swaggerClient: swaggerClient,
		config:        cfg,
		ctx:           context.Background(),
	}
}

// GetNetworks retrieves all networks from the Netmaker API
func (c *Client) GetNetworks() ([]models.Network, error) {
	logrus.Info("Retrieving networks from Netmaker API")

	// Use the REST client directly since we're having issues with the Swagger client
	// We'll use a map to parse the raw JSON first
	resp, err := c.restClient.R().Get("/api/networks")
	if err != nil {
		logrus.Errorf("REST API error: %v", err)
		return nil, fmt.Errorf("failed to get networks: %w", err)
	}

	if resp.IsError() {
		logrus.Errorf("REST API error response: %s", resp.String())
		return nil, fmt.Errorf("failed to get networks: %s", resp.Status())
	}

	logrus.Debugf("REST API response status: %s", resp.Status())

	// Parse the JSON response
	var swaggerNetworks []swagger.Network
	err = json.Unmarshal(resp.Body(), &swaggerNetworks)
	if err != nil {
		logrus.Errorf("Failed to parse network response: %v", err)
		return nil, fmt.Errorf("failed to parse network response: %w", err)
	}

	logrus.Debugf("Retrieved %d networks from REST API", len(swaggerNetworks))

	// Convert Swagger models to our internal models
	networks := make([]models.Network, len(swaggerNetworks))
	for i, swaggerNetwork := range swaggerNetworks {
		// Map fields from swagger.Network to our models.Network
		networks[i] = models.Network{
			ID:                     swaggerNetwork.Netid,
			Version:                1,                    // Default to version 1 for new networks
			Name:                   swaggerNetwork.Netid, // No display name in the model
			AddressRange:           swaggerNetwork.Addressrange,
			AddressRange6:          swaggerNetwork.Addressrange6,
			LocalRange:             "", // Not available in swagger model
			IsDualStack:            swaggerNetwork.Isipv4 != "" && swaggerNetwork.Isipv6 != "",
			IsIPv4:                 swaggerNetwork.Isipv4 == "yes",
			IsIPv6:                 swaggerNetwork.Isipv6 == "yes",
			IsLocal:                false, // Not available in swagger model
			DefaultAccessControl:   swaggerNetwork.Defaultacl,
			DefaultUDPHolePunching: swaggerNetwork.Defaultudpholepunch == "yes",
			DefaultExtClientDNS:    "", // Not available in swagger model
			DefaultMTU:             int(swaggerNetwork.Defaultmtu),
			DefaultKeepalive:       int(swaggerNetwork.Defaultkeepalive),
			DefaultInterface:       swaggerNetwork.Defaultinterface,
			NodeLimit:              int(swaggerNetwork.Nodelimit),
			IsCurrent:              true,
			LastModified:           time.Now(),
			CreatedAt:              time.Now(),
			Data:                   models.JSONB{}, // Initialize empty JSONB
		}
		logrus.Debugf("Converted network: %s", networks[i].ID)
	}

	logrus.Infof("Retrieved and converted %d networks from Netmaker API", len(networks))
	return networks, nil
}

// GetNodes retrieves all nodes for a network from the Netmaker API
func (c *Client) GetNodes(networkID string) ([]models.Node, error) {
	logrus.Debugf("Getting nodes for network %s using Swagger client", networkID)

	// Use the Swagger client to get nodes
	swaggerNodes, resp, err := c.swaggerClient.NodesApi.GetNetworkNodes(c.ctx, networkID)
	if err != nil {
		logrus.Errorf("Swagger API error getting nodes: %v", err)
		return nil, fmt.Errorf("failed to get nodes for network %s: %w", networkID, err)
	}

	logrus.Debugf("Swagger API response status for nodes: %s", resp.Status)
	logrus.Debugf("Retrieved %d nodes from Swagger API", len(swaggerNodes))

	// Convert to our internal model
	nodes := make([]models.Node, len(swaggerNodes))
	for i, swaggerNode := range swaggerNodes {
		// Map fields from swagger.ApiNode to our models.Node
		// This is a simplified mapping, you may need to adjust based on your models
		nodes[i] = models.Node{
			ID:               swaggerNode.Id,
			NetworkID:        swaggerNode.Network,
			Name:             swaggerNode.Id, // Using ID as name since there's no name field
			Address:          swaggerNode.Address,
			Address6:         swaggerNode.Address6,
			PublicKey:        "", // Not available in ApiNode
			Endpoint:         "", // Not available in ApiNode
			IsEgressGateway:  swaggerNode.Isegressgateway,
			IsIngressGateway: swaggerNode.Isingressgateway,
			IsRelay:          swaggerNode.Isrelay,
			Connected:        swaggerNode.Connected,
			// Add other fields as needed
		}
		logrus.Debugf("Converted node: %s", nodes[i].ID)
	}

	logrus.Infof("Retrieved and converted %d nodes for network %s from Netmaker API", len(nodes), networkID)
	return nodes, nil
}

// GetExtClients retrieves all external clients for a network from the Netmaker API
func (c *Client) GetExtClients(networkID string) ([]models.ExtClient, error) {
	logrus.Debugf("Getting ext clients for network %s using Swagger client", networkID)

	// Use the Swagger client
	swaggerExtClients, resp, err := c.swaggerClient.ExtClientApi.GetNetworkExtClients(c.ctx, networkID)
	if err != nil {
		logrus.Errorf("Swagger API error getting ext clients: %v", err)
		return nil, fmt.Errorf("failed to get external clients for network %s: %w", networkID, err)
	}

	logrus.Debugf("Swagger API response status for ext clients: %s", resp.Status)
	logrus.Debugf("Retrieved %d ext clients from Swagger API", len(swaggerExtClients))

	// Convert to our internal model
	extClients := make([]models.ExtClient, len(swaggerExtClients))
	for i, swaggerExtClient := range swaggerExtClients {
		// Map fields from swagger.ExtClient to our models.ExtClient
		extClients[i] = models.ExtClient{
			ID:        swaggerExtClient.Clientid,
			NetworkID: swaggerExtClient.Network,
			Name:      swaggerExtClient.Clientid, // Using clientid as name since there's no name field
			Address:   swaggerExtClient.Address,
			Address6:  swaggerExtClient.Address6,
			PublicKey: swaggerExtClient.Publickey,
			Enabled:   swaggerExtClient.Enabled,
			// Add other fields as needed
		}
		logrus.Debugf("Converted ext client: %s", extClients[i].ID)
	}

	logrus.Infof("Retrieved and converted %d external clients for network %s from Netmaker API", len(extClients), networkID)
	return extClients, nil
}

// GetDNSEntries retrieves all DNS entries for a network from the Netmaker API
func (c *Client) GetDNSEntries(networkID string) ([]models.DNSEntry, error) {
	logrus.Infof("Retrieving DNS entries for network %s from Netmaker API", networkID)

	// Note: The Swagger client doesn't have a direct method for retrieving DNS entries
	// so we'll use the REST client directly
	var dnsEntries []models.DNSEntry
	resp, err := c.restClient.R().SetResult(&dnsEntries).Get(fmt.Sprintf("/api/dns/%s", networkID))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DNS entries: %w", err)
	}

	logrus.Debugf("REST API response status for DNS entries: %s", resp.Status())
	logrus.Debugf("Retrieved %d DNS entries from REST API", len(dnsEntries))
	// No need to convert as we're using the REST client directly

	logrus.Infof("Retrieved and converted %d DNS entries for network %s from Netmaker API", len(dnsEntries), networkID)
	return dnsEntries, nil
}

// GetACLs retrieves all ACLs for a network from the Netmaker API
func (c *Client) GetACLs(networkID string) (map[string]map[string]int, error) {
	logrus.Infof("Retrieving ACLs for network %s from Netmaker API", networkID)

	// Note: The Swagger client doesn't have a direct method for retrieving ACLs
	// so we'll use the REST client directly
	var aclsMap map[string]map[string]int
	resp, err := c.restClient.R().SetResult(&aclsMap).Get(fmt.Sprintf("/api/networks/%s/acls", networkID))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ACLs: %w", err)
	}

	logrus.Debugf("REST API response status for ACLs: %s", resp.Status())
	logrus.Debugf("Retrieved ACL map from REST API")

	logrus.Infof("Retrieved ACL map for network %s from Netmaker API", networkID)
	return aclsMap, nil
}

// GetHosts retrieves all hosts from the Netmaker API
func (c *Client) GetHosts() ([]models.Host, error) {
	logrus.Info("Retrieving hosts from Netmaker API")

	// Use the Swagger client to get hosts
	swaggerHosts, resp, err := c.swaggerClient.HostsApi.GetHosts(c.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve hosts: %w", err)
	}

	logrus.Debugf("Swagger API response status for hosts: %d", resp.StatusCode)
	logrus.Debugf("Retrieved %d hosts from Swagger API", len(swaggerHosts))

	// Convert Swagger hosts to our model
	hosts := make([]models.Host, len(swaggerHosts))
	for i, swaggerHost := range swaggerHosts {
		hosts[i] = models.Host{
			ID:                  swaggerHost.Id,
			Version:             1, // Default to version 1 for new hosts
			Name:                swaggerHost.Name,
			EndpointIP:          swaggerHost.Endpointip,
			EndpointIPv6:        swaggerHost.Endpointipv6,
			PublicKey:           swaggerHost.Publickey,
			ListenPort:          int(swaggerHost.Listenport),
			MTU:                 int(swaggerHost.Mtu),
			PersistentKeepalive: int(swaggerHost.Persistentkeepalive),
			IsCurrent:           true,
			LastModified:        time.Now(),
			CreatedAt:           time.Now(),
			Data:                models.JSONB{},
		}
	}

	logrus.Infof("Retrieved and converted %d hosts from Netmaker API", len(hosts))
	return hosts, nil
}
