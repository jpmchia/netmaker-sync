package db

import (
	"fmt"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
)

// GenericUpsert provides a simplified approach to versioned record management
// It follows these rules:
// 1. Check if the record in the DB matches the record from the API - if so, do not update
// 2. If there are differences, insert a new row with the results from the API and increment the version number
//
// Parameters:
// - db: The database connection
// - tableName: The name of the table to operate on
// - idField: The name of the ID field in the table (e.g., "id" or "network_id")
// - idValue: The value of the ID field
// - currentRecord: Pointer to a struct containing the current record from the database (if it exists)
// - newRecord: Pointer to a struct containing the new record from the API
// - equalsFn: Function that compares the current and new records to determine if they are equal
// - getVersionFn: Function that returns the version field from a record
// - setVersionFn: Function that sets the version field on a record
// - setLastModifiedFn: Function that sets the LastModified field on a record
// - insertFn: Function that inserts the record into the database
//
// Returns:
// - bool: true if a new version was created, false if no changes were made
// - error: any error that occurred during the operation
func (db *DB) GenericUpsert(
	tableName string,
	idField string,
	idValue interface{},
	currentRecord interface{},
	newRecord interface{},
	equalsFn func(current, new interface{}) bool,
	getVersionFn func(record interface{}) int,
	setVersionFn func(record interface{}, version int),
	setLastModifiedFn func(record interface{}, time time.Time),
	insertFn func(tx interface{}, record interface{}) error,
) (bool, error) {
	// Check if a current version exists
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)", tableName, idField)
	err := db.QueryRow(query, idValue).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if record exists: %w", err)
	}

	// If no record exists, insert the first version
	if !exists {
		// Set version to 1 and last modified to now
		setVersionFn(newRecord, 1)
		setLastModifiedFn(newRecord, time.Now())

		// Start a transaction
		tx, err := db.Beginx()
		if err != nil {
			return false, fmt.Errorf("failed to begin transaction: %w", err)
		}
		defer tx.Rollback()

		// Insert the new record
		if err := insertFn(tx, newRecord); err != nil {
			return false, fmt.Errorf("failed to insert first version: %w", err)
		}

		// Commit the transaction
		if err := tx.Commit(); err != nil {
			return false, fmt.Errorf("failed to commit transaction: %w", err)
		}

		logrus.Infof("Created new record in %s with %s = %v", tableName, idField, idValue)
		return true, nil
	}

	// If the current record is nil or not of the expected type, we can't compare
	if currentRecord == nil || reflect.ValueOf(currentRecord).IsNil() {
		return false, fmt.Errorf("current record is nil, cannot compare")
	}

	// Check if there are meaningful changes
	if equalsFn(currentRecord, newRecord) {
		// No changes, nothing to do
		logrus.Debugf("No changes for record in %s with %s = %v, skipping update", tableName, idField, idValue)
		return false, nil
	}

	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Set the current version to not current
	updateQuery := fmt.Sprintf("UPDATE %s SET is_current = false WHERE %s = $1 AND is_current = true", tableName, idField)
	_, err = tx.Exec(updateQuery, idValue)
	if err != nil {
		return false, fmt.Errorf("failed to update current record: %w", err)
	}

	// Get the next version number
	nextVersion := getVersionFn(currentRecord) + 1

	// Set version and last modified
	setVersionFn(newRecord, nextVersion)
	setLastModifiedFn(newRecord, time.Now())

	// Insert the new version
	if err := insertFn(tx, newRecord); err != nil {
		return false, fmt.Errorf("failed to insert new version: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return false, fmt.Errorf("failed to commit transaction: %w", err)
	}

	logrus.Infof("Updated record in %s with %s = %v to version %d", tableName, idField, idValue, nextVersion)
	return true, nil
}
