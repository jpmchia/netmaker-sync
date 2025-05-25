package db

import (
	"database/sql"
	"errors"
	"fmt"
	"netmaker-sync/internal/models"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx"
)

// UpsertNode inserts or updates a node in the database
func (db *DB) UpsertNode(node *models.Node) error {
	// Get the current node if it exists
	var currentNode models.Node
	err := db.Get(&currentNode, `
		SELECT * FROM nodes 
		WHERE id = $1 AND is_current = true
	`, node.ID)

	// Define the functions needed for the generic upsert
	// Function to check if nodes are equal
	equalsFn := func(current, new interface{}) bool {
		currentNode := current.(*models.Node)
		newNode := new.(*models.Node)
		return nodesEqual(*currentNode, *newNode)
	}

	// Function to get the version from a node
	getVersionFn := func(record interface{}) int {
		return record.(*models.Node).Version
	}

	// Function to set the version on a node
	setVersionFn := func(record interface{}, version int) {
		record.(*models.Node).Version = version
	}

	// Function to set the last modified time on a node
	setLastModifiedFn := func(record interface{}, lastModified time.Time) {
		record.(*models.Node).LastModified = lastModified
	}

	// Function to insert a node into the database
	insertFn := func(tx interface{}, record interface{}) error {
		_, err := tx.(*sqlx.Tx).NamedExec(`
			INSERT INTO nodes (
				id, version, network_id, name, address, address6, public_key,
				endpoint, is_egress_gateway, is_ingress_gateway, is_relay,
				connected, is_current, last_modified, created_at, data
			) VALUES (
				:id, :version, :network_id, :name, :address, :address6, :public_key,
				:endpoint, :is_egress_gateway, :is_ingress_gateway, :is_relay,
				:connected, true, :last_modified, NOW(), :data
			)
		`, record)
		return err
	}

	// If we couldn't find the current node, pass nil for the current record
	var currentNodePtr *models.Node
	if err == nil {
		currentNodePtr = &currentNode
	} else if !errors.Is(err, sql.ErrNoRows) {
		// If there was an error other than not finding the record, return it
		return fmt.Errorf("failed to get current node: %w", err)
	}

	// Call the generic upsert function
	_, err = db.GenericUpsert(
		"nodes",
		"id",
		node.ID,
		currentNodePtr,
		node,
		equalsFn,
		getVersionFn,
		setVersionFn,
		setLastModifiedFn,
		insertFn,
	)

	return err
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
