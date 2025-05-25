package db

import (
	"database/sql"
	"errors"
	"fmt"
	"netmaker-sync/internal/models"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
)

// UpsertNetwork inserts or updates a network in the database
func (db *DB) UpsertNetwork(network *models.Network) error {
	// Check if the network exists with any version
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM networks WHERE id = $1)
	`, network.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if network exists: %w", err)
	}

	// If network exists, get the current version
	if exists {
		var currentNetwork models.Network
		err := db.Get(&currentNetwork, `
			SELECT * FROM networks 
			WHERE id = $1 AND is_current = true
		`, network.ID)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("failed to get current network: %w", err)
		}

		// If we found a current version, check for changes
		if err == nil {
			// Network exists, check if there are meaningful changes
			if networksEqual(currentNetwork, *network) {
				// No changes, nothing to do
				logrus.Debugf("No changes for network %s, skipping update", network.ID)
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
			network.Version = nextVersion
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
			`, network)

			if err != nil {
				return fmt.Errorf("failed to insert new network version: %w", err)
			}

			logrus.Infof("Updated network %s to version %d", network.ID, nextVersion)
			return tx.Commit()
		}
	}

	// Network doesn't exist or no current version, create it
	network.Version = 1
	network.LastModified = time.Now()

	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

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
	`, network)

	if err != nil {
		return fmt.Errorf("failed to insert first network version: %w", err)
	}

	logrus.Infof("Created new network %s", network.ID)
	return tx.Commit()
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

// GetNetworks retrieves all current networks from the database
func (db *DB) GetNetworks() ([]models.Network, error) {
	// Get only the current versions of networks
	var networks []models.Network
	err := db.Select(&networks, `
		SELECT * FROM networks 
		WHERE is_current = true
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get networks: %w", err)
	}
	return networks, nil
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
