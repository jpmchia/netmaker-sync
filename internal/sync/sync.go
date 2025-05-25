// sync.go
package sync

import (
	"context"
	"netmaker-sync/internal/api"
	"netmaker-sync/internal/db"
	"netmaker-sync/internal/models"
	"time"

	"github.com/sirupsen/logrus"
)

// timePtr returns a pointer to the given time
func timePtr(t time.Time) *time.Time {
	return &t
}

// Service handles syncing data from Netmaker API to the database
type Service struct {
	apiClient *api.Client
	db        *db.DB
}

// New creates a new sync service
func New(apiClient *api.Client, db *db.DB) *Service {
	return &Service{
		apiClient: apiClient,
		db:        db,
	}
}

// SyncAll syncs all resources from Netmaker API to the database
func (s *Service) SyncAll(ctx context.Context, includeAcls bool) error {
	// Start with networks
	if err := s.SyncNetworks(ctx); err != nil {
		return err
	}

	// Get all networks to sync their resources
	networks, err := s.db.GetNetworks()
	if err != nil {
		return err
	}

	// For each network, sync nodes, ext clients, DNS entries, and ACLs
	for _, network := range networks {
		if err := s.SyncNodes(ctx, network.ID); err != nil {
			logrus.Errorf("Failed to sync nodes for network %s: %v", network.ID, err)
		}

		if err := s.SyncExtClients(ctx, network.ID); err != nil {
			logrus.Errorf("Failed to sync ext clients for network %s: %v", network.ID, err)
		}

		if err := s.SyncDNSEntries(ctx, network.ID); err != nil {
			logrus.Errorf("Failed to sync DNS entries for network %s: %v", network.ID, err)
		}

		if includeAcls {
			if err := s.SyncACLs(ctx, network.ID); err != nil {
				logrus.Errorf("Failed to sync ACLs for network %s: %v", network.ID, err)
			}
		}
	}

	// Sync hosts
	if err := s.SyncHosts(ctx); err != nil {
		logrus.Errorf("Failed to sync hosts: %v", err)
	}

	return nil
}

// SyncNetworks syncs networks from Netmaker API to the database
func (s *Service) SyncNetworks(ctx context.Context) error {
	// Record sync start
	syncHistory := &models.SyncHistory{
		ResourceType: models.ResourceTypeNetwork,
		Status:       models.SyncStatusPending,
		StartedAt:    time.Now(),
	}
	if err := s.db.CreateSyncHistory(syncHistory); err != nil {
		return err
	}

	// Get networks from API
	networks, err := s.apiClient.GetNetworks()
	if err != nil {
		// Record sync failure
		syncHistory.Status = models.SyncStatusFailed
		syncHistory.Message = err.Error()
		syncHistory.CompletedAt = timePtr(time.Now())
		s.db.UpdateSyncHistory(syncHistory)
		return err
	}

	// Upsert networks to database
	for _, network := range networks {
		if err := s.db.UpsertNetwork(&network); err != nil {
			logrus.Errorf("Failed to upsert network %s: %v", network.ID, err)
		}
	}

	// Record sync completion
	syncHistory.Status = models.SyncStatusCompleted
	syncHistory.CompletedAt = timePtr(time.Now())
	return s.db.UpdateSyncHistory(syncHistory)
}

// SyncNodes syncs nodes for a network from Netmaker API to the database
func (s *Service) SyncNodes(ctx context.Context, networkID string) error {
	// Record sync start
	syncHistory := &models.SyncHistory{
		ResourceType: models.ResourceTypeNode,
		Status:       models.SyncStatusPending,
		StartedAt:    time.Now(),
	}
	if err := s.db.CreateSyncHistory(syncHistory); err != nil {
		return err
	}

	// Get nodes from API
	nodes, err := s.apiClient.GetNodes(networkID)
	if err != nil {
		// Record sync failure
		syncHistory.Status = models.SyncStatusFailed
		syncHistory.Message = err.Error()
		syncHistory.CompletedAt = timePtr(time.Now())
		s.db.UpdateSyncHistory(syncHistory)
		return err
	}

	// Upsert nodes to database
	for _, node := range nodes {
		if err := s.db.UpsertNode(&node); err != nil {
			logrus.Errorf("Failed to upsert node %s: %v", node.ID, err)
		}
	}

	// Record sync completion
	syncHistory.Status = models.SyncStatusCompleted
	syncHistory.CompletedAt = timePtr(time.Now())
	return s.db.UpdateSyncHistory(syncHistory)
}

// SyncExtClients syncs external clients for a network from Netmaker API to the database
func (s *Service) SyncExtClients(ctx context.Context, networkID string) error {
	// Record sync start
	syncHistory := &models.SyncHistory{
		ResourceType: models.ResourceTypeExtClient,
		Status:       models.SyncStatusPending,
		StartedAt:    time.Now(),
	}
	if err := s.db.CreateSyncHistory(syncHistory); err != nil {
		return err
	}

	// Get external clients from API
	extClients, err := s.apiClient.GetExtClients(networkID)
	if err != nil {
		// Record sync failure
		syncHistory.Status = models.SyncStatusFailed
		syncHistory.Message = err.Error()
		syncHistory.CompletedAt = timePtr(time.Now())
		s.db.UpdateSyncHistory(syncHistory)
		return err
	}

	// Upsert external clients to database
	for _, extClient := range extClients {
		if err := s.db.UpsertExtClient(&extClient); err != nil {
			logrus.Errorf("Failed to upsert external client %s: %v", extClient.ID, err)
		}
	}

	// Record sync completion
	syncHistory.Status = models.SyncStatusCompleted
	syncHistory.CompletedAt = timePtr(time.Now())
	return s.db.UpdateSyncHistory(syncHistory)
}

// SyncDNSEntries syncs DNS entries for a network from Netmaker API to the database
func (s *Service) SyncDNSEntries(ctx context.Context, networkID string) error {
	// Record sync start
	syncHistory := &models.SyncHistory{
		ResourceType: models.ResourceTypeDNS,
		Status:       models.SyncStatusPending,
		StartedAt:    time.Now(),
	}
	if err := s.db.CreateSyncHistory(syncHistory); err != nil {
		return err
	}

	// Get DNS entries from API
	dnsEntries, err := s.apiClient.GetDNSEntries(networkID)
	if err != nil {
		// Record sync failure
		syncHistory.Status = models.SyncStatusFailed
		syncHistory.Message = err.Error()
		syncHistory.CompletedAt = timePtr(time.Now())
		s.db.UpdateSyncHistory(syncHistory)
		return err
	}

	// Upsert DNS entries to database
	for _, dnsEntry := range dnsEntries {
		if err := s.db.UpsertDNSEntry(&dnsEntry); err != nil {
			logrus.Errorf("Failed to upsert DNS entry %s: %v", dnsEntry.Name, err)
		}
	}

	// Record sync completion
	syncHistory.Status = models.SyncStatusCompleted
	syncHistory.CompletedAt = timePtr(time.Now())
	return s.db.UpdateSyncHistory(syncHistory)
}

// SyncACLs syncs ACLs for a network from Netmaker API to the database
func (s *Service) SyncACLs(ctx context.Context, networkID string) error {
	// Record sync start
	syncHistory := &models.SyncHistory{
		ResourceType: models.ResourceTypeACL,
		Status:       models.SyncStatusPending,
		StartedAt:    time.Now(),
	}
	if err := s.db.CreateSyncHistory(syncHistory); err != nil {
		return err
	}

	// Get ACLs from API
	acls, err := s.apiClient.GetACLs(networkID)
	if err != nil {
		// Record sync failure
		syncHistory.Status = models.SyncStatusFailed
		syncHistory.Message = err.Error()
		syncHistory.CompletedAt = timePtr(time.Now())
		s.db.UpdateSyncHistory(syncHistory)
		return err
	}

	// Upsert ACLs to database
	if err := s.db.UpsertACLs(networkID, acls); err != nil {
		logrus.Errorf("Failed to upsert ACLs for network %s: %v", networkID, err)
	}

	// Record sync completion
	syncHistory.Status = models.SyncStatusCompleted
	syncHistory.CompletedAt = timePtr(time.Now())
	return s.db.UpdateSyncHistory(syncHistory)
}

// SyncHosts syncs hosts from Netmaker API to the database
func (s *Service) SyncHosts(ctx context.Context) error {
	// Record sync start
	syncHistory := &models.SyncHistory{
		ResourceType: models.ResourceTypeHost,
		Status:       models.SyncStatusPending,
		StartedAt:    time.Now(),
	}
	if err := s.db.CreateSyncHistory(syncHistory); err != nil {
		return err
	}

	// Get hosts from API
	hosts, err := s.apiClient.GetHosts()
	if err != nil {
		// Record sync failure
		syncHistory.Status = models.SyncStatusFailed
		syncHistory.Message = err.Error()
		syncHistory.CompletedAt = timePtr(time.Now())
		s.db.UpdateSyncHistory(syncHistory)
		return err
	}

	// Upsert hosts to database
	for _, host := range hosts {
		if err := s.db.UpsertHost(&host); err != nil {
			logrus.Errorf("Failed to upsert host %s: %v", host.ID, err)
		}
	}

	// Record sync completion
	syncHistory.Status = models.SyncStatusCompleted
	syncHistory.CompletedAt = timePtr(time.Now())
	return s.db.UpdateSyncHistory(syncHistory)
}

// GetNetworks retrieves all networks from the database
func (s *Service) GetNetworks(ctx context.Context) ([]models.Network, error) {
	return s.db.GetNetworks()
}

// GetNetwork retrieves a specific network from the database
func (s *Service) GetNetwork(ctx context.Context, networkID string) (*models.Network, error) {
	return s.db.GetNetwork(networkID)
}
