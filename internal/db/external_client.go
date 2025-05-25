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

func (db *DB) UpsertExtClient(extClient *models.ExtClient) error {
	// Check if the ext client exists with any version
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM ext_clients WHERE id = $1)
	`, extClient.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if ext client exists: %w", err)
	}

	// If ext client exists, get the current version
	if exists {
		var currentExtClient models.ExtClient
		err := db.Get(&currentExtClient, `
			SELECT * FROM ext_clients 
			WHERE id = $1 AND is_current = true
		`, extClient.ID)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("failed to get current ext client: %w", err)
		}

		// If we found a current version, check for changes
		if err == nil {
			// Ext client exists, check if there are meaningful changes
			if extClientsEqual(currentExtClient, *extClient) {
				// No changes, nothing to do
				logrus.Debugf("No changes for ext client %s, skipping update", extClient.ID)
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
			extClient.Version = nextVersion
			extClient.LastModified = time.Now()
			_, err = tx.NamedExec(`
				INSERT INTO ext_clients (
					id, version, network_id, name, address, address6, public_key,
					enabled, is_current, last_modified, created_at, data
				) VALUES (
					:id, :version, :network_id, :name, :address, :address6, :public_key,
					:enabled, true, :last_modified, NOW(), :data
				)
			`, extClient)

			if err != nil {
				return fmt.Errorf("failed to insert new ext client version: %w", err)
			}

			logrus.Infof("Updated ext client %s to version %d", extClient.ID, nextVersion)
			return tx.Commit()
		}
	}

	// Ext client doesn't exist or no current version, create it
	extClient.Version = 1
	extClient.LastModified = time.Now()

	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.NamedExec(`
		INSERT INTO ext_clients (
			id, version, network_id, name, address, address6, public_key,
			enabled, is_current, last_modified, created_at, data
		) VALUES (
			:id, :version, :network_id, :name, :address, :address6, :public_key,
			:enabled, true, :last_modified, NOW(), :data
		)
	`, extClient)

	if err != nil {
		return fmt.Errorf("failed to insert first ext client version: %w", err)
	}

	logrus.Infof("Created new ext client %s", extClient.ID)
	return tx.Commit()
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
