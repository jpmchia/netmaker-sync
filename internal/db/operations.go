package db

import (
	"fmt"
	"netmaker-sync/internal/models"
)

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
