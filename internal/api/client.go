package api

import (
	"fmt"
	"netmaker-sync/internal/config"
	"netmaker-sync/internal/models"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Client is a wrapper around the Netmaker API client
type Client struct {
	client *resty.Client
	config *config.NetmakerAPIConfig
}

// New creates a new Netmaker API client
func New(cfg *config.NetmakerAPIConfig) *Client {
	client := resty.New()
	client.SetBaseURL(cfg.URL)
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.Key))
	client.SetTimeout(30 * time.Second)
	client.SetRetryCount(3)
	client.SetRetryWaitTime(5 * time.Second)
	client.SetRetryMaxWaitTime(20 * time.Second)

	return &Client{
		client: client,
		config: cfg,
	}
}

// GetNetworks retrieves all networks from the Netmaker API
func (c *Client) GetNetworks() ([]models.Network, error) {
	var networks []models.Network
	resp, err := c.client.R().
		SetResult(&networks).
		Get("/api/networks")

	if err != nil {
		return nil, fmt.Errorf("failed to get networks: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get networks: %s", resp.String())
	}

	logrus.Infof("Retrieved %d networks from Netmaker API", len(networks))
	return networks, nil
}

// GetNodes retrieves all nodes for a network from the Netmaker API
func (c *Client) GetNodes(networkID string) ([]models.Node, error) {
	var nodes []models.Node
	resp, err := c.client.R().
		SetResult(&nodes).
		Get(fmt.Sprintf("/api/nodes/%s", networkID))

	if err != nil {
		return nil, fmt.Errorf("failed to get nodes for network %s: %w", networkID, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get nodes for network %s: %s", networkID, resp.String())
	}

	logrus.Infof("Retrieved %d nodes for network %s from Netmaker API", len(nodes), networkID)
	return nodes, nil
}

// GetExtClients retrieves all external clients for a network from the Netmaker API
func (c *Client) GetExtClients(networkID string) ([]models.ExtClient, error) {
	var extClients []models.ExtClient
	resp, err := c.client.R().
		SetResult(&extClients).
		Get(fmt.Sprintf("/api/extclients/%s", networkID))

	if err != nil {
		return nil, fmt.Errorf("failed to get external clients for network %s: %w", networkID, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get external clients for network %s: %s", networkID, resp.String())
	}

	logrus.Infof("Retrieved %d external clients for network %s from Netmaker API", len(extClients), networkID)
	return extClients, nil
}

// GetDNSEntries retrieves all DNS entries for a network from the Netmaker API
func (c *Client) GetDNSEntries(networkID string) ([]models.DNSEntry, error) {
	var dnsEntries []models.DNSEntry
	resp, err := c.client.R().
		SetResult(&dnsEntries).
		Get(fmt.Sprintf("/api/dns/adm/%s", networkID))

	if err != nil {
		return nil, fmt.Errorf("failed to get DNS entries for network %s: %w", networkID, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get DNS entries for network %s: %s", networkID, resp.String())
	}

	logrus.Infof("Retrieved %d DNS entries for network %s from Netmaker API", len(dnsEntries), networkID)
	return dnsEntries, nil
}

// GetHosts retrieves all hosts from the Netmaker API
func (c *Client) GetHosts() ([]models.Host, error) {
	var hosts []models.Host
	resp, err := c.client.R().
		SetResult(&hosts).
		Get("/api/hosts")

	if err != nil {
		return nil, fmt.Errorf("failed to get hosts: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get hosts: %s", resp.String())
	}

	logrus.Infof("Retrieved %d hosts from Netmaker API", len(hosts))
	return hosts, nil
}

// GetACLs retrieves all ACLs for a network from the Netmaker API
func (c *Client) GetACLs(networkID string) (map[string]map[string]int, error) {
	var acls map[string]map[string]int
	resp, err := c.client.R().
		SetResult(&acls).
		Get(fmt.Sprintf("/api/networks/%s/acls", networkID))

	if err != nil {
		return nil, fmt.Errorf("failed to get ACLs for network %s: %w", networkID, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get ACLs for network %s: %s", networkID, resp.String())
	}

	logrus.Infof("Retrieved ACLs for network %s from Netmaker API", networkID)
	return acls, nil
}
