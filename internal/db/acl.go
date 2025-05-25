// db/acl.go
package db

import (
	"database/sql"
	"errors"
	"fmt"
	"netmaker-sync/internal/models"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (db *DB) UpsertACL(acl *models.ACL) error {
	// Get the current ACL if it exists
	var currentACL models.ACL
	err := db.Get(&currentACL, `
		SELECT * FROM acls 
		WHERE id = $1 AND is_current = true
	`, acl.ID)

	// Define the functions needed for the generic upsert
	// Function to check if ACLs are equal
	equalsFn := func(current, new interface{}) bool {
		currentACL := current.(*models.ACL)
		newACL := new.(*models.ACL)
		return aclsEqual(*currentACL, *newACL)
	}

	// Function to get the version from an ACL
	getVersionFn := func(record interface{}) int {
		return record.(*models.ACL).Version
	}

	// Function to set the version on an ACL
	setVersionFn := func(record interface{}, version int) {
		record.(*models.ACL).Version = version
	}

	// Function to set the last modified time on an ACL
	setLastModifiedFn := func(record interface{}, lastModified time.Time) {
		record.(*models.ACL).LastModified = lastModified
	}

	// Function to insert an ACL into the database
	insertFn := func(tx interface{}, record interface{}) error {
		_, err := tx.(*sqlx.Tx).NamedExec(`
			INSERT INTO acls (
				id, version, network_id, node_id, is_current, 
				last_modified, created_at, data
			) VALUES (
				:id, :version, :network_id, :node_id, true, 
				:last_modified, NOW(), :data
			)
		`, record)
		return err
	}

	// If we couldn't find the current ACL, pass nil for the current record
	var currentACLPtr *models.ACL
	if err == nil {
		currentACLPtr = &currentACL
	} else if !errors.Is(err, sql.ErrNoRows) {
		// If there was an error other than not finding the record, return it
		return fmt.Errorf("failed to get current ACL: %w", err)
	}

	// Call the generic upsert function
	_, err = db.GenericUpsert(
		"acls",
		"id",
		acl.ID,
		currentACLPtr,
		acl,
		equalsFn,
		getVersionFn,
		setVersionFn,
		setLastModifiedFn,
		insertFn,
	)

	return err
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
	// First, delete all existing ACLs for this network without a transaction
	// This is safer than trying to do everything in a single transaction
	_, err := db.Exec(`DELETE FROM acls WHERE network_id = $1`, networkID)
	if err != nil {
		return fmt.Errorf("failed to delete existing ACLs: %w", err)
	}

	// Get the next ID to use for new ACLs
	var nextID int
	err = db.Get(&nextID, `SELECT COALESCE(MAX(id), 0) + 1 FROM acls`)
	if err != nil {
		return fmt.Errorf("failed to get next ACL ID: %w", err)
	}

	// Get all nodes for this network to check if they exist
	var nodes []models.Node
	err = db.Select(&nodes, `
		SELECT id, name FROM nodes 
		WHERE network_id = $1 AND is_current = true
	`, networkID)
	if err != nil {
		return fmt.Errorf("failed to get nodes for network: %w", err)
	}

	// Create a map of node names to node IDs for quick lookup
	nodeMap := make(map[string]string)
	for _, node := range nodes {
		nodeMap[node.Name] = node.ID
	}

	// Process each ACL from the map
	successCount := 0
	failureCount := 0
	for sourceNode, destMap := range aclsMap {
		// Check if the source node exists in our database
		sourceID, exists := nodeMap[sourceNode]
		if !exists {
			logrus.Warnf("Source node '%s' not found in database, skipping ACLs", sourceNode)
			failureCount += len(destMap)
			continue
		}

		for destNode, allowed := range destMap {
			// Create a new ACL object using the node ID instead of the node name
			acl := &models.ACL{
				ID:        nextID,
				NetworkID: networkID,
				NodeID:    sourceID, // Use the actual node ID from the database
				Data: models.JSONB{
					"source_node": sourceNode,
					"dest_node":   destNode,
					"is_allowed":  allowed == 1,
				},
			}

			// Use the UpsertACL function which now uses the generic approach
			err := db.UpsertACL(acl)
			if err != nil {
				logrus.Warnf("Failed to upsert ACL for source %s and dest %s: %v", sourceNode, destNode, err)
				failureCount++
				continue
			}

			logrus.Infof("Created new record in acls with id = %d", nextID)
			// Increment the ID for the next ACL
			nextID++
			successCount++
		}
	}

	logrus.Infof("Inserted %d ACLs for network %s (failed: %d)", successCount, networkID, failureCount)
	return nil
}
