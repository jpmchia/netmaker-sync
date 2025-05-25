package db

import (
	"database/sql"
	"errors"
	"fmt"
	"netmaker-sync/internal/models"
	"time"

	"github.com/sirupsen/logrus"
)

func (db *DB) UpsertDNSEntry(dnsEntry *models.DNSEntry) error {
	// Check if the DNS entry exists with any version
	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM dns_entries WHERE id = $1)
	`, dnsEntry.ID).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if DNS entry exists: %w", err)
	}

	// If DNS entry exists, get the current version
	if exists {
		var currentDNSEntry models.DNSEntry
		err := db.Get(&currentDNSEntry, `
			SELECT * FROM dns_entries 
			WHERE id = $1 AND is_current = true
		`, dnsEntry.ID)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("failed to get current DNS entry: %w", err)
		}

		// If we found a current version, check for changes
		if err == nil {
			// DNS entry exists, check if there are meaningful changes
			if dnsEntriesEqual(currentDNSEntry, *dnsEntry) {
				// No changes, nothing to do
				logrus.Debugf("No changes for DNS entry %s, skipping update", dnsEntry.ID)
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

			logrus.Infof("Updated DNS entry %s to version %d", dnsEntry.ID, nextVersion)
			return tx.Commit()
		}
	}

	// DNS entry doesn't exist or no current version, create it
	dnsEntry.LastModified = time.Now()

	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.NamedExec(`
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

	logrus.Infof("Created new DNS entry %s", dnsEntry.ID)
	return tx.Commit()
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
