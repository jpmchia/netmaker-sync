// db/operations.go
package db

import (
	"database/sql"
	"errors"
	"fmt"
	"netmaker-sync/internal/models"
	"reflect"
	"time"
)

func (db *DB) UpsertNetwork(network *models.Network) error {
	// Get the current version from the database
	var currentNetwork models.Network
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM networks 
			WHERE id = $1 AND is_current = true
		)
	`, network.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if network exists: %w", err)
	}

	if exists {
		// Get the current network data
		err = db.Get(&currentNetwork, `
			SELECT * FROM networks 
			WHERE id = $1 AND is_current = true
		`, network.ID)

		if err != nil {
			return fmt.Errorf("failed to get current network: %w", err)
		}

		// Check if there are meaningful changes
		if networksEqual(currentNetwork, *network) {
			// No changes, nothing to do
			return nil
		}

		// Start a transaction
		tx, err := db.Beginx()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		// Set the current version to not current
		_, err = tx.Exec(`
			UPDATE networks 
			SET is_current = false 
			WHERE id = $1 AND is_current = true
		`, network.ID)

		if err != nil {
			return fmt.Errorf("failed to update current network: %w", err)
		}

		// Get the next version number
		var nextVersion int
		err = tx.Get(&nextVersion, `
			SELECT COALESCE(MAX(version), 0) + 1 FROM networks WHERE id = $1
		`, network.ID)

		if err != nil {
			return fmt.Errorf("failed to get next version: %w", err)
		}

		// Insert the new version
		network.LastModified = time.Now()
		_, err = tx.NamedExec(`
			INSERT INTO networks (
				id, version, name, address_range, address_range6, local_range,
				is_dual_stack, is_ipv4, is_ipv6, is_local, default_access_control,
				default_udp_hole_punching, default_ext_client_dns, default_mtu,
				default_keepalive, default_interface, node_limit, is_current,
				last_modified, created_at, data
			) VALUES (
				:id, :version, :name, :address_range, :address_range6, :local_range,
				:is_dual_stack, :is_ipv4, :is_ipv6, :is_local, :default_access_control,
				:default_udp_hole_punching, :default_ext_client_dns, :default_mtu,
				:default_keepalive, :default_interface, :node_limit, true,
				:last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":                        network.ID,
			"version":                   nextVersion,
			"name":                      network.Name,
			"address_range":             network.AddressRange,
			"address_range6":            network.AddressRange6,
			"local_range":               network.LocalRange,
			"is_dual_stack":             network.IsDualStack,
			"is_ipv4":                   network.IsIPv4,
			"is_ipv6":                   network.IsIPv6,
			"is_local":                  network.IsLocal,
			"default_access_control":    network.DefaultAccessControl,
			"default_udp_hole_punching": network.DefaultUDPHolePunching,
			"default_ext_client_dns":    network.DefaultExtClientDNS,
			"default_mtu":               network.DefaultMTU,
			"default_keepalive":         network.DefaultKeepalive,
			"default_interface":         network.DefaultInterface,
			"node_limit":                network.NodeLimit,
			"last_modified":             network.LastModified,
			"data":                      network.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert new network version: %w", err)
		}

		return tx.Commit()
	} else {
		// First version of this network
		network.LastModified = time.Now()
		_, err = db.NamedExec(`
			INSERT INTO networks (
				id, version, name, address_range, address_range6, local_range,
				is_dual_stack, is_ipv4, is_ipv6, is_local, default_access_control,
				default_udp_hole_punching, default_ext_client_dns, default_mtu,
				default_keepalive, default_interface, node_limit, is_current,
				last_modified, created_at, data
			) VALUES (
				:id, 1, :name, :address_range, :address_range6, :local_range,
				:is_dual_stack, :is_ipv4, :is_ipv6, :is_local, :default_access_control,
				:default_udp_hole_punching, :default_ext_client_dns, :default_mtu,
				:default_keepalive, :default_interface, :node_limit, true,
				:last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":                        network.ID,
			"name":                      network.Name,
			"address_range":             network.AddressRange,
			"address_range6":            network.AddressRange6,
			"local_range":               network.LocalRange,
			"is_dual_stack":             network.IsDualStack,
			"is_ipv4":                   network.IsIPv4,
			"is_ipv6":                   network.IsIPv6,
			"is_local":                  network.IsLocal,
			"default_access_control":    network.DefaultAccessControl,
			"default_udp_hole_punching": network.DefaultUDPHolePunching,
			"default_ext_client_dns":    network.DefaultExtClientDNS,
			"default_mtu":               network.DefaultMTU,
			"default_keepalive":         network.DefaultKeepalive,
			"default_interface":         network.DefaultInterface,
			"node_limit":                network.NodeLimit,
			"last_modified":             network.LastModified,
			"data":                      network.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert first network version: %w", err)
		}

		return nil
	}
}

// networksEqual compares two networks to determine if there are meaningful changes
func networksEqual(a, b models.Network) bool {
	// Compare relevant fields, ignoring metadata like LastModified
	return a.Name == b.Name &&
		a.AddressRange == b.AddressRange &&
		a.AddressRange6 == b.AddressRange6 &&
		a.LocalRange == b.LocalRange &&
		a.IsDualStack == b.IsDualStack &&
		a.IsIPv4 == b.IsIPv4 &&
		a.IsIPv6 == b.IsIPv6 &&
		a.IsLocal == b.IsLocal &&
		a.DefaultAccessControl == b.DefaultAccessControl &&
		a.DefaultUDPHolePunching == b.DefaultUDPHolePunching &&
		a.DefaultExtClientDNS == b.DefaultExtClientDNS &&
		a.DefaultMTU == b.DefaultMTU &&
		a.DefaultKeepalive == b.DefaultKeepalive &&
		a.DefaultInterface == b.DefaultInterface &&
		a.NodeLimit == b.NodeLimit &&
		reflect.DeepEqual(a.Data, b.Data)
}

func (db *DB) GetNetworks() ([]models.Network, error) {
	// Get only the current versions of networks
	var networks []models.Network
	err := db.Select(&networks, `
		SELECT * FROM networks 
		WHERE is_current = true
	`)
	return networks, err
}

// GetNetwork retrieves a specific network by ID
func (db *DB) GetNetwork(networkID string) (*models.Network, error) {
	var network models.Network
	err := db.Get(&network, `
		SELECT * FROM networks 
		WHERE id = $1 AND is_current = true
	`, networkID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("network not found: %s", networkID)
		}
		return nil, fmt.Errorf("failed to get network: %w", err)
	}
	return &network, nil
}

// Similar methods for other resources
func (db *DB) UpsertNode(node *models.Node) error {
	// Get the current version from the database
	var currentNode models.Node
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM nodes 
			WHERE id = $1 AND is_current = true
		)
	`, node.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if node exists: %w", err)
	}

	if exists {
		// Get the current node data
		err = db.Get(&currentNode, `
			SELECT * FROM nodes 
			WHERE id = $1 AND is_current = true
		`, node.ID)

		if err != nil {
			return fmt.Errorf("failed to get current node: %w", err)
		}

		// Check if there are meaningful changes
		if nodesEqual(currentNode, *node) {
			// No changes, nothing to do
			return nil
		}

		// Start a transaction
		tx, err := db.Beginx()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		// Set the current version to not current
		_, err = tx.Exec(`
			UPDATE nodes 
			SET is_current = false 
			WHERE id = $1 AND is_current = true
		`, node.ID)

		if err != nil {
			return fmt.Errorf("failed to update current node: %w", err)
		}

		// Get the next version number
		var nextVersion int
		err = tx.Get(&nextVersion, `
			SELECT COALESCE(MAX(version), 0) + 1 FROM nodes WHERE id = $1
		`, node.ID)

		if err != nil {
			return fmt.Errorf("failed to get next version: %w", err)
		}

		// Insert the new version
		node.LastModified = time.Now()
		_, err = tx.NamedExec(`
			INSERT INTO nodes (
				id, version, network_id, name, address, address6, public_key,
				endpoint, is_egress_gateway, is_ingress_gateway, is_relay,
				connected, is_current, last_modified, created_at, data
			) VALUES (
				:id, :version, :network_id, :name, :address, :address6, :public_key,
				:endpoint, :is_egress_gateway, :is_ingress_gateway, :is_relay,
				:connected, true, :last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":                 node.ID,
			"version":            nextVersion,
			"network_id":         node.NetworkID,
			"name":               node.Name,
			"address":            node.Address,
			"address6":           node.Address6,
			"public_key":         node.PublicKey,
			"endpoint":           node.Endpoint,
			"is_egress_gateway":  node.IsEgressGateway,
			"is_ingress_gateway": node.IsIngressGateway,
			"is_relay":           node.IsRelay,
			"connected":          node.Connected,
			"last_modified":      node.LastModified,
			"data":               node.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert new node version: %w", err)
		}

		return tx.Commit()
	} else {
		// First version of this node
		node.LastModified = time.Now()
		_, err = db.NamedExec(`
			INSERT INTO nodes (
				id, version, network_id, name, address, address6, public_key,
				endpoint, is_egress_gateway, is_ingress_gateway, is_relay,
				connected, is_current, last_modified, created_at, data
			) VALUES (
				:id, 1, :network_id, :name, :address, :address6, :public_key,
				:endpoint, :is_egress_gateway, :is_ingress_gateway, :is_relay,
				:connected, true, :last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":                 node.ID,
			"network_id":         node.NetworkID,
			"name":               node.Name,
			"address":            node.Address,
			"address6":           node.Address6,
			"public_key":         node.PublicKey,
			"endpoint":           node.Endpoint,
			"is_egress_gateway":  node.IsEgressGateway,
			"is_ingress_gateway": node.IsIngressGateway,
			"is_relay":           node.IsRelay,
			"connected":          node.Connected,
			"last_modified":      node.LastModified,
			"data":               node.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert first node version: %w", err)
		}

		return nil
	}
}

// nodesEqual compares two nodes to determine if there are meaningful changes
func nodesEqual(a, b models.Node) bool {
	// Compare relevant fields, ignoring metadata like LastModified
	return a.NetworkID == b.NetworkID &&
		a.Name == b.Name &&
		a.Address == b.Address &&
		a.Address6 == b.Address6 &&
		a.PublicKey == b.PublicKey &&
		a.Endpoint == b.Endpoint &&
		a.IsEgressGateway == b.IsEgressGateway &&
		a.IsIngressGateway == b.IsIngressGateway &&
		a.IsRelay == b.IsRelay &&
		a.Connected == b.Connected &&
		reflect.DeepEqual(a.Data, b.Data)
}

func (db *DB) GetNodes(networkID string) ([]models.Node, error) {
	// Get only the current versions of nodes for a network
	var nodes []models.Node
	err := db.Select(&nodes, `
		SELECT * FROM nodes 
		WHERE network_id = $1 AND is_current = true
	`, networkID)
	return nodes, err
}

// GetNodeHistory retrieves the version history of a node
func (db *DB) GetNodeHistory(nodeID string) ([]models.Node, error) {
	var nodes []models.Node
	err := db.Select(&nodes, `
		SELECT * FROM nodes 
		WHERE id = $1
		ORDER BY version DESC
	`, nodeID)
	return nodes, err
}

func (db *DB) UpsertExtClient(extClient *models.ExtClient) error {
	// Get the current version from the database
	var currentExtClient models.ExtClient
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM ext_clients 
			WHERE id = $1 AND is_current = true
		)
	`, extClient.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if ext client exists: %w", err)
	}

	if exists {
		// Get the current ext client data
		err = db.Get(&currentExtClient, `
			SELECT * FROM ext_clients 
			WHERE id = $1 AND is_current = true
		`, extClient.ID)

		if err != nil {
			return fmt.Errorf("failed to get current ext client: %w", err)
		}

		// Check if there are meaningful changes
		if extClientsEqual(currentExtClient, *extClient) {
			// No changes, nothing to do
			return nil
		}

		// Start a transaction
		tx, err := db.Beginx()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		// Set the current version to not current
		_, err = tx.Exec(`
			UPDATE ext_clients 
			SET is_current = false 
			WHERE id = $1 AND is_current = true
		`, extClient.ID)

		if err != nil {
			return fmt.Errorf("failed to update current ext client: %w", err)
		}

		// Get the next version number
		var nextVersion int
		err = tx.Get(&nextVersion, `
			SELECT COALESCE(MAX(version), 0) + 1 FROM ext_clients WHERE id = $1
		`, extClient.ID)

		if err != nil {
			return fmt.Errorf("failed to get next version: %w", err)
		}

		// Insert the new version
		extClient.LastModified = time.Now()
		_, err = tx.NamedExec(`
			INSERT INTO ext_clients (
				id, version, network_id, name, address, address6, public_key,
				enabled, is_current, last_modified, created_at, data
			) VALUES (
				:id, :version, :network_id, :name, :address, :address6, :public_key,
				:enabled, true, :last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":            extClient.ID,
			"version":       nextVersion,
			"network_id":    extClient.NetworkID,
			"name":          extClient.Name,
			"address":       extClient.Address,
			"address6":      extClient.Address6,
			"public_key":    extClient.PublicKey,
			"enabled":       extClient.Enabled,
			"last_modified": extClient.LastModified,
			"data":          extClient.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert new ext client version: %w", err)
		}

		return tx.Commit()
	} else {
		// First version of this ext client
		extClient.LastModified = time.Now()
		_, err = db.NamedExec(`
			INSERT INTO ext_clients (
				id, version, network_id, name, address, address6, public_key,
				enabled, is_current, last_modified, created_at, data
			) VALUES (
				:id, 1, :network_id, :name, :address, :address6, :public_key,
				:enabled, true, :last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":            extClient.ID,
			"network_id":    extClient.NetworkID,
			"name":          extClient.Name,
			"address":       extClient.Address,
			"address6":      extClient.Address6,
			"public_key":    extClient.PublicKey,
			"enabled":       extClient.Enabled,
			"last_modified": extClient.LastModified,
			"data":          extClient.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert first ext client version: %w", err)
		}

		return nil
	}
}

// extClientsEqual compares two external clients to determine if there are meaningful changes
func extClientsEqual(a, b models.ExtClient) bool {
	// Compare relevant fields, ignoring metadata like LastModified
	return a.NetworkID == b.NetworkID &&
		a.Name == b.Name &&
		a.Address == b.Address &&
		a.Address6 == b.Address6 &&
		a.PublicKey == b.PublicKey &&
		a.Enabled == b.Enabled &&
		reflect.DeepEqual(a.Data, b.Data)
}

func (db *DB) GetExtClients(networkID string) ([]models.ExtClient, error) {
	// Get only the current versions of external clients for a network
	var extClients []models.ExtClient
	err := db.Select(&extClients, `
		SELECT * FROM ext_clients 
		WHERE network_id = $1 AND is_current = true
	`, networkID)
	return extClients, err
}

// GetExtClientHistory retrieves the version history of an external client
func (db *DB) GetExtClientHistory(extClientID string) ([]models.ExtClient, error) {
	var extClients []models.ExtClient
	err := db.Select(&extClients, `
		SELECT * FROM ext_clients 
		WHERE id = $1
		ORDER BY version DESC
	`, extClientID)
	return extClients, err
}

func (db *DB) UpsertDNSEntry(dnsEntry *models.DNSEntry) error {
	// Get the current version from the database
	var currentDNSEntry models.DNSEntry
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM dns_entries 
			WHERE id = $1 AND is_current = true
		)
	`, dnsEntry.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if DNS entry exists: %w", err)
	}

	if exists {
		// Get the current DNS entry data
		err = db.Get(&currentDNSEntry, `
			SELECT * FROM dns_entries 
			WHERE id = $1 AND is_current = true
		`, dnsEntry.ID)

		if err != nil {
			return fmt.Errorf("failed to get current DNS entry: %w", err)
		}

		// Check if there are meaningful changes
		if dnsEntriesEqual(currentDNSEntry, *dnsEntry) {
			// No changes, nothing to do
			return nil
		}

		// Start a transaction
		tx, err := db.Beginx()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		// Set the current version to not current
		_, err = tx.Exec(`
			UPDATE dns_entries 
			SET is_current = false 
			WHERE id = $1 AND is_current = true
		`, dnsEntry.ID)

		if err != nil {
			return fmt.Errorf("failed to update current DNS entry: %w", err)
		}

		// Get the next version number
		var nextVersion int
		err = tx.Get(&nextVersion, `
			SELECT COALESCE(MAX(version), 0) + 1 FROM dns_entries WHERE id = $1
		`, dnsEntry.ID)

		if err != nil {
			return fmt.Errorf("failed to get next version: %w", err)
		}

		// Insert the new version
		dnsEntry.LastModified = time.Now()
		_, err = tx.NamedExec(`
			INSERT INTO dns_entries (
				id, version, name, network_id, address, address6,
				is_current, last_modified, created_at
			) VALUES (
				:id, :version, :name, :network_id, :address, :address6,
				true, :last_modified, NOW()
			)
		`, map[string]interface{}{
			"id":            dnsEntry.ID,
			"version":       nextVersion,
			"name":          dnsEntry.Name,
			"network_id":    dnsEntry.NetworkID,
			"address":       dnsEntry.Address,
			"address6":      dnsEntry.Address6,
			"last_modified": dnsEntry.LastModified,
		})

		if err != nil {
			return fmt.Errorf("failed to insert new DNS entry version: %w", err)
		}

		return tx.Commit()
	} else {
		// First version of this DNS entry
		dnsEntry.LastModified = time.Now()
		_, err = db.NamedExec(`
			INSERT INTO dns_entries (
				id, version, name, network_id, address, address6,
				is_current, last_modified, created_at
			) VALUES (
				:id, 1, :name, :network_id, :address, :address6,
				true, :last_modified, NOW()
			)
		`, map[string]interface{}{
			"id":            dnsEntry.ID,
			"name":          dnsEntry.Name,
			"network_id":    dnsEntry.NetworkID,
			"address":       dnsEntry.Address,
			"address6":      dnsEntry.Address6,
			"last_modified": dnsEntry.LastModified,
		})

		if err != nil {
			return fmt.Errorf("failed to insert first DNS entry version: %w", err)
		}

		return nil
	}
}

// dnsEntriesEqual compares two DNS entries to determine if there are meaningful changes
func dnsEntriesEqual(a, b models.DNSEntry) bool {
	// Compare relevant fields, ignoring metadata like LastModified
	return a.Name == b.Name &&
		a.NetworkID == b.NetworkID &&
		a.Address == b.Address &&
		a.Address6 == b.Address6
}

func (db *DB) GetDNSEntries(networkID string) ([]models.DNSEntry, error) {
	// Get only the current versions of DNS entries for a network
	var dnsEntries []models.DNSEntry
	err := db.Select(&dnsEntries, `
		SELECT * FROM dns_entries 
		WHERE network_id = $1 AND is_current = true
	`, networkID)
	return dnsEntries, err
}

// GetDNSEntryHistory retrieves the version history of a DNS entry
func (db *DB) GetDNSEntryHistory(dnsEntryID string) ([]models.DNSEntry, error) {
	var dnsEntries []models.DNSEntry
	err := db.Select(&dnsEntries, `
		SELECT * FROM dns_entries 
		WHERE id = $1
		ORDER BY version DESC
	`, dnsEntryID)
	return dnsEntries, err
}

func (db *DB) UpsertHost(host *models.Host) error {
	// Get the current version from the database
	var currentHost models.Host
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM hosts 
			WHERE id = $1 AND is_current = true
		)
	`, host.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if host exists: %w", err)
	}

	if exists {
		// Get the current host data
		err = db.Get(&currentHost, `
			SELECT * FROM hosts 
			WHERE id = $1 AND is_current = true
		`, host.ID)

		if err != nil {
			return fmt.Errorf("failed to get current host: %w", err)
		}

		// Check if there are meaningful changes
		if hostsEqual(currentHost, *host) {
			// No changes, nothing to do
			return nil
		}

		// Start a transaction
		tx, err := db.Beginx()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		// Set the current version to not current
		_, err = tx.Exec(`
			UPDATE hosts 
			SET is_current = false 
			WHERE id = $1 AND is_current = true
		`, host.ID)

		if err != nil {
			return fmt.Errorf("failed to update current host: %w", err)
		}

		// Get the next version number
		var nextVersion int
		err = tx.Get(&nextVersion, `
			SELECT COALESCE(MAX(version), 0) + 1 FROM hosts WHERE id = $1
		`, host.ID)

		if err != nil {
			return fmt.Errorf("failed to get next version: %w", err)
		}

		// Insert the new version
		host.LastModified = time.Now()
		_, err = tx.NamedExec(`
			INSERT INTO hosts (
				id, version, name, endpoint_ip, endpoint_ipv6, public_key,
				listen_port, mtu, persistent_keepalive, is_current,
				last_modified, created_at, data
			) VALUES (
				:id, :version, :name, :endpoint_ip, :endpoint_ipv6, :public_key,
				:listen_port, :mtu, :persistent_keepalive, true,
				:last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":                   host.ID,
			"version":              nextVersion,
			"name":                 host.Name,
			"endpoint_ip":          host.EndpointIP,
			"endpoint_ipv6":        host.EndpointIPv6,
			"public_key":           host.PublicKey,
			"listen_port":          host.ListenPort,
			"mtu":                  host.MTU,
			"persistent_keepalive": host.PersistentKeepalive,
			"last_modified":        host.LastModified,
			"data":                 host.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert new host version: %w", err)
		}

		return tx.Commit()
	} else {
		// First version of this host
		host.LastModified = time.Now()
		_, err = db.NamedExec(`
			INSERT INTO hosts (
				id, version, name, endpoint_ip, endpoint_ipv6, public_key,
				listen_port, mtu, persistent_keepalive, is_current,
				last_modified, created_at, data
			) VALUES (
				:id, 1, :name, :endpoint_ip, :endpoint_ipv6, :public_key,
				:listen_port, :mtu, :persistent_keepalive, true,
				:last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":                   host.ID,
			"name":                 host.Name,
			"endpoint_ip":          host.EndpointIP,
			"endpoint_ipv6":        host.EndpointIPv6,
			"public_key":           host.PublicKey,
			"listen_port":          host.ListenPort,
			"mtu":                  host.MTU,
			"persistent_keepalive": host.PersistentKeepalive,
			"last_modified":        host.LastModified,
			"data":                 host.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert first host version: %w", err)
		}

		return nil
	}
}

// hostsEqual compares two hosts to determine if there are meaningful changes
func hostsEqual(a, b models.Host) bool {
	// Compare relevant fields, ignoring metadata like LastModified
	return a.Name == b.Name &&
		a.EndpointIP == b.EndpointIP &&
		a.EndpointIPv6 == b.EndpointIPv6 &&
		a.PublicKey == b.PublicKey &&
		a.ListenPort == b.ListenPort &&
		a.MTU == b.MTU &&
		a.PersistentKeepalive == b.PersistentKeepalive &&
		reflect.DeepEqual(a.Data, b.Data)
}

func (db *DB) GetHosts() ([]models.Host, error) {
	// Get only the current versions of hosts
	var hosts []models.Host
	err := db.Select(&hosts, `
		SELECT * FROM hosts 
		WHERE is_current = true
	`)
	return hosts, err
}

// GetHostHistory retrieves the version history of a host
func (db *DB) GetHostHistory(hostID string) ([]models.Host, error) {
	var hosts []models.Host
	err := db.Select(&hosts, `
		SELECT * FROM hosts 
		WHERE id = $1
		ORDER BY version DESC
	`, hostID)
	return hosts, err
}

func (db *DB) UpsertACL(acl *models.ACL) error {
	// Get the current version from the database
	var currentACL models.ACL
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM acls 
			WHERE id = $1 AND is_current = true
		)
	`, acl.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if ACL exists: %w", err)
	}

	if exists {
		// Get the current ACL data
		err = db.Get(&currentACL, `
			SELECT * FROM acls 
			WHERE id = $1 AND is_current = true
		`, acl.ID)

		if err != nil {
			return fmt.Errorf("failed to get current ACL: %w", err)
		}

		// Check if there are meaningful changes
		if aclsEqual(currentACL, *acl) {
			// No changes, nothing to do
			return nil
		}

		// Start a transaction
		tx, err := db.Beginx()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		// Set the current version to not current
		_, err = tx.Exec(`
			UPDATE acls 
			SET is_current = false 
			WHERE id = $1 AND is_current = true
		`, acl.ID)

		if err != nil {
			return fmt.Errorf("failed to update current ACL: %w", err)
		}

		// Get the next version number
		var nextVersion int
		err = tx.Get(&nextVersion, `
			SELECT COALESCE(MAX(version), 0) + 1 FROM acls WHERE id = $1
		`, acl.ID)

		if err != nil {
			return fmt.Errorf("failed to get next version: %w", err)
		}

		// Insert the new version
		acl.LastModified = time.Now()
		_, err = tx.NamedExec(`
			INSERT INTO acls (
				id, version, network_id, node_id, is_current, 
				last_modified, created_at, data
			) VALUES (
				:id, :version, :network_id, :node_id, true, 
				:last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":            acl.ID,
			"version":       nextVersion,
			"network_id":    acl.NetworkID,
			"node_id":       acl.NodeID,
			"last_modified": acl.LastModified,
			"data":          acl.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert new ACL version: %w", err)
		}

		return tx.Commit()
	} else {
		// First version of this ACL
		acl.LastModified = time.Now()
		_, err = db.NamedExec(`
			INSERT INTO acls (
				id, version, network_id, node_id, is_current, 
				last_modified, created_at, data
			) VALUES (
				:id, 1, :network_id, :node_id, true, 
				:last_modified, NOW(), :data
			)
		`, map[string]interface{}{
			"id":            acl.ID,
			"network_id":    acl.NetworkID,
			"node_id":       acl.NodeID,
			"last_modified": acl.LastModified,
			"data":          acl.Data,
		})

		if err != nil {
			return fmt.Errorf("failed to insert first ACL version: %w", err)
		}

		return nil
	}
}

// aclsEqual compares two ACLs to determine if there are meaningful changes
func aclsEqual(a, b models.ACL) bool {
	// Compare relevant fields, ignoring metadata like LastModified
	return a.NetworkID == b.NetworkID &&
		a.NodeID == b.NodeID &&
		reflect.DeepEqual(a.Data, b.Data)
}

func (db *DB) GetACLs(networkID string) ([]models.ACL, error) {
	// Get only the current versions of ACLs for a network
	var acls []models.ACL
	err := db.Select(&acls, `
		SELECT * FROM acls 
		WHERE network_id = $1 AND is_current = true
	`, networkID)
	return acls, err
}

// GetACLHistory retrieves the version history of an ACL
func (db *DB) GetACLHistory(aclID int) ([]models.ACL, error) {
	var acls []models.ACL
	err := db.Select(&acls, `
		SELECT * FROM acls 
		WHERE id = $1
		ORDER BY version DESC
	`, aclID)
	return acls, err
}

func (db *DB) UpsertACLs(networkID string, aclsMap map[string]map[string]int) error {
	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Get existing ACLs for this network to avoid ID conflicts
	existingACLs := make(map[string]int) // Map of source_node-dest_node to ACL ID
	var existingACLsList []models.ACL
	err = tx.Select(&existingACLsList, `
		SELECT * FROM acls 
		WHERE network_id = $1 AND is_current = true
	`, networkID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to get existing ACLs: %w", err)
	}

	// Build a map of existing ACLs by source and dest nodes
	for _, acl := range existingACLsList {
		if acl.Data != nil {
			sourceNode, _ := acl.Data["source_node"].(string)
			destNode, _ := acl.Data["dest_node"].(string)
			if sourceNode != "" && destNode != "" {
				key := sourceNode + "-" + destNode
				existingACLs[key] = acl.ID
			}
		}
	}

	// Process each ACL from the map
	for sourceNode, destMap := range aclsMap {
		for destNode, allowed := range destMap {
			// Store the ACL data in JSONB
			aclData := models.JSONB{
				"source_node": sourceNode,
				"dest_node":   destNode,
				"is_allowed":  allowed == 1,
			}

			// Check if this ACL already exists
			key := sourceNode + "-" + destNode
			aclID, exists := existingACLs[key]

			if exists {
				// Get the current ACL data
				var currentACL models.ACL
				err = tx.Get(&currentACL, `
					SELECT * FROM acls 
					WHERE id = $1 AND is_current = true
				`, aclID)

				if err != nil {
					return fmt.Errorf("failed to get current ACL: %w", err)
				}

				// Check if there are meaningful changes
				isAllowed, _ := currentACL.Data["is_allowed"].(bool)
				if (allowed == 1) == isAllowed {
					// No changes, nothing to do for this ACL
					continue
				}

				// Set the current version to not current
				_, err = tx.Exec(`
					UPDATE acls 
					SET is_current = false 
					WHERE id = $1 AND is_current = true
				`, aclID)

				if err != nil {
					return fmt.Errorf("failed to update current ACL: %w", err)
				}

				// Get the next version number
				var nextVersion int
				err = tx.Get(&nextVersion, `
					SELECT COALESCE(MAX(version), 0) + 1 FROM acls WHERE id = $1
				`, aclID)

				if err != nil {
					return fmt.Errorf("failed to get next version: %w", err)
				}

				// Insert the new version
				_, err = tx.NamedExec(`
					INSERT INTO acls (
						id, version, network_id, node_id, is_current, 
						last_modified, created_at, data
					) VALUES (
						:id, :version, :network_id, :node_id, true, 
						:last_modified, NOW(), :data
					)
				`, map[string]interface{}{
					"id":            aclID,
					"version":       nextVersion,
					"network_id":    networkID,
					"node_id":       sourceNode,
					"last_modified": time.Now(),
					"data":          aclData,
				})

				if err != nil {
					return fmt.Errorf("failed to insert new ACL version: %w", err)
				}
			} else {
				// First version of this ACL - get a new ID
				var newID int
				err = tx.Get(&newID, `SELECT COALESCE(MAX(id), 0) + 1 FROM acls`)
				if err != nil {
					return fmt.Errorf("failed to get new ACL ID: %w", err)
				}

				// Insert the new ACL
				_, err = tx.NamedExec(`
					INSERT INTO acls (
						id, version, network_id, node_id, is_current, 
						last_modified, created_at, data
					) VALUES (
						:id, 1, :network_id, :node_id, true, 
						:last_modified, NOW(), :data
					)
				`, map[string]interface{}{
					"id":            newID,
					"network_id":    networkID,
					"node_id":       sourceNode,
					"last_modified": time.Now(),
					"data":          aclData,
				})

				if err != nil {
					return fmt.Errorf("failed to insert first ACL version: %w", err)
				}
			}
		}
	}

	return tx.Commit()
}

func (db *DB) CreateSyncHistory(syncHistory *models.SyncHistory) error {
	// Insert a new sync history record
	query := `
		INSERT INTO sync_history (
			resource_type, status, message, started_at, completed_at
		) VALUES (
			$1, $2, $3, $4, $5
		)
		RETURNING id
	`
	return db.QueryRow(query,
		syncHistory.ResourceType,
		syncHistory.Status,
		syncHistory.Message,
		syncHistory.StartedAt,
		syncHistory.CompletedAt).Scan(&syncHistory.ID)
}

func (db *DB) UpdateSyncHistory(syncHistory *models.SyncHistory) error {
	// Update an existing sync history record
	query := `
		UPDATE sync_history
		SET status = $1, message = $2, completed_at = $3
		WHERE id = $4
	`
	_, err := db.Exec(query,
		syncHistory.Status,
		syncHistory.Message,
		syncHistory.CompletedAt,
		syncHistory.ID)

	if err != nil {
		return fmt.Errorf("failed to update sync history: %w", err)
	}

	return nil
}
