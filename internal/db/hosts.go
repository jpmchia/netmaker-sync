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

func (db *DB) UpsertHost(host *models.Host) error {
	// Check if the host exists with any version
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM hosts WHERE id = $1)
	`, host.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if host exists: %w", err)
	}

	// If host exists, get the current version
	if exists {
		var currentHost models.Host
		err := db.Get(&currentHost, `
			SELECT * FROM hosts 
			WHERE id = $1 AND is_current = true
		`, host.ID)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("failed to get current host: %w", err)
		}

		// If we found a current version, check for changes
		if err == nil {
			// Host exists, check if there are meaningful changes
			if hostsEqual(currentHost, *host) {
				// No changes, nothing to do
				logrus.Debugf("No changes for host %s, skipping update", host.ID)
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
			host.Version = nextVersion
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
			`, host)

			if err != nil {
				return fmt.Errorf("failed to insert new host version: %w", err)
			}

			logrus.Infof("Updated host %s to version %d", host.ID, nextVersion)
			return tx.Commit()
		}
	}

	// Host doesn't exist or no current version, create it
	host.Version = 1
	host.LastModified = time.Now()

	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

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
	`, host)

	if err != nil {
		return fmt.Errorf("failed to insert first host version: %w", err)
	}

	logrus.Infof("Created new host %s", host.ID)
	return tx.Commit()
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
