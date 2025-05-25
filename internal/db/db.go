package db

import (
	"fmt"
	"netmaker-sync/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
)

// DB is a wrapper around sqlx.DB
type DB struct {
	*sqlx.DB
}

// New creates a new database connection
func New(cfg *config.DatabaseConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logrus.Info("Connected to database")
	return &DB{db}, nil
}

// Initialize creates the necessary tables if they don't exist
func (db *DB) Initialize() error {
	// Create the schema_migrations table to track migrations
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INTEGER PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// Run migrations
	return db.runMigrations()
}

// runMigrations runs all database migrations
func (db *DB) runMigrations() error {
	migrations := []struct {
		version int
		up      string
	}{
		{
			version: 1,
			up: `
				CREATE TABLE IF NOT EXISTS networks (
					id TEXT NOT NULL,
					version INTEGER NOT NULL,
					name TEXT NOT NULL,
					address_range TEXT,
					address_range6 TEXT,
					local_range TEXT,
					is_dual_stack BOOLEAN NOT NULL DEFAULT FALSE,
					is_ipv4 BOOLEAN NOT NULL DEFAULT TRUE,
					is_ipv6 BOOLEAN NOT NULL DEFAULT FALSE,
					is_local BOOLEAN NOT NULL DEFAULT FALSE,
					default_access_control TEXT,
					default_udp_hole_punching BOOLEAN NOT NULL DEFAULT TRUE,
					default_ext_client_dns TEXT,
					default_mtu INTEGER,
					default_keepalive INTEGER,
					default_interface TEXT,
					node_limit INTEGER,
					is_current BOOLEAN NOT NULL DEFAULT TRUE,
					last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					data JSONB,
					PRIMARY KEY (id, version)
				)
			`,
		},
		{
			version: 2,
			up: `
				CREATE TABLE IF NOT EXISTS nodes (
					id TEXT NOT NULL,
					version INTEGER NOT NULL,
					network_id TEXT NOT NULL,
					name TEXT NOT NULL,
					address TEXT,
					address6 TEXT,
					public_key TEXT,
					endpoint TEXT,
					is_egress_gateway BOOLEAN NOT NULL DEFAULT FALSE,
					is_ingress_gateway BOOLEAN NOT NULL DEFAULT FALSE,
					is_relay BOOLEAN NOT NULL DEFAULT FALSE,
					connected BOOLEAN NOT NULL DEFAULT FALSE,
					is_current BOOLEAN NOT NULL DEFAULT TRUE,
					last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					data JSONB,
					PRIMARY KEY (id, version),
					UNIQUE(network_id, name, version)
				)
			`,
		},
		{
			version: 3,
			up: `
				CREATE TABLE IF NOT EXISTS ext_clients (
					id TEXT NOT NULL,
					version INTEGER NOT NULL,
					network_id TEXT NOT NULL,
					name TEXT NOT NULL,
					address TEXT,
					address6 TEXT,
					public_key TEXT,
					enabled BOOLEAN NOT NULL DEFAULT TRUE,
					is_current BOOLEAN NOT NULL DEFAULT TRUE,
					last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					data JSONB,
					PRIMARY KEY (id, version),
					UNIQUE(network_id, name, version)
				)
			`,
		},
		{
			version: 4,
			up: `
				-- Add a unique constraint on networks.id to support foreign key references
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM pg_constraint 
						WHERE conname = 'networks_id_unique' AND conrelid = 'networks'::regclass
					) THEN
						ALTER TABLE networks ADD CONSTRAINT networks_id_unique UNIQUE (id);
					END IF;
				END
				$$;
				
				CREATE TABLE IF NOT EXISTS dns_entries (
					id SERIAL PRIMARY KEY,
					network_id TEXT NOT NULL REFERENCES networks(id) ON DELETE CASCADE,
					name TEXT NOT NULL,
					address TEXT,
					address6 TEXT,
					last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					UNIQUE(network_id, name)
				)
			`,
		},
		{
			version: 5,
			up: `
				CREATE TABLE IF NOT EXISTS hosts (
					id TEXT NOT NULL,
					version INTEGER NOT NULL,
					name TEXT NOT NULL,
					endpoint_ip TEXT,
					endpoint_ipv6 TEXT,
					public_key TEXT,
					listen_port INTEGER,
					mtu INTEGER,
					persistent_keepalive INTEGER,
					is_current BOOLEAN NOT NULL DEFAULT TRUE,
					last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					data JSONB,
					PRIMARY KEY (id, version)
				)
			`,
		},
		{
			version: 6,
			up: `
				-- Add a unique constraint on nodes.id to support foreign key references
				DO $$
				BEGIN
					IF NOT EXISTS (
						SELECT 1 FROM pg_constraint 
						WHERE conname = 'nodes_id_unique' AND conrelid = 'nodes'::regclass
					) THEN
						ALTER TABLE nodes ADD CONSTRAINT nodes_id_unique UNIQUE (id);
					END IF;
				END
				$$;
				
				CREATE TABLE IF NOT EXISTS acls (
					id SERIAL PRIMARY KEY,
					network_id TEXT NOT NULL REFERENCES networks(id) ON DELETE CASCADE,
					node_id TEXT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
					data JSONB NOT NULL,
					last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					UNIQUE(network_id, node_id)
				)
			`,
		},
		{
			version: 7,
			up: `
				CREATE TABLE IF NOT EXISTS sync_history (
					id SERIAL PRIMARY KEY,
					resource_type TEXT NOT NULL,
					status TEXT NOT NULL,
					message TEXT,
					started_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
					completed_at TIMESTAMP WITH TIME ZONE
				)
			`,
		},
		{
			version: 8,
			up: `
				-- Add version and is_current columns to the acls table
				ALTER TABLE acls ADD COLUMN IF NOT EXISTS version INTEGER NOT NULL DEFAULT 1;
				ALTER TABLE acls ADD COLUMN IF NOT EXISTS is_current BOOLEAN NOT NULL DEFAULT TRUE;
				-- Update the primary key to include version
				ALTER TABLE acls DROP CONSTRAINT IF EXISTS acls_pkey;
				ALTER TABLE acls ADD PRIMARY KEY (id, version);
			`,
		},
		{
			version: 9,
			up: `
				-- Drop all unique constraints on the acls table except for the primary key
				ALTER TABLE acls DROP CONSTRAINT IF EXISTS acls_network_id_node_id_key;
				ALTER TABLE acls DROP CONSTRAINT IF EXISTS acls_network_id_node_id_version_key;
			`,
		},
		{
			version: 10,
			up: `
				-- Truncate the acls table to start fresh
				TRUNCATE TABLE acls;
			`,
		},
	}

	// Check the current migration version
	var currentVersion int
	err := db.Get(&currentVersion, `
		SELECT COALESCE(MAX(version), 0) FROM schema_migrations
	`)
	if err != nil {
		return fmt.Errorf("failed to get current migration version: %w", err)
	}

	// Run migrations that haven't been applied yet
	for _, migration := range migrations {
		if migration.version > currentVersion {
			logrus.Infof("Running migration %d", migration.version)

			tx, err := db.Beginx()
			if err != nil {
				return fmt.Errorf("failed to begin transaction for migration %d: %w", migration.version, err)
			}

			_, err = tx.Exec(migration.up)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to run migration %d: %w", migration.version, err)
			}

			_, err = tx.Exec(`
				INSERT INTO schema_migrations (version) VALUES ($1)
			`, migration.version)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to record migration %d: %w", migration.version, err)
			}

			err = tx.Commit()
			if err != nil {
				return fmt.Errorf("failed to commit migration %d: %w", migration.version, err)
			}

			logrus.Infof("Migration %d completed", migration.version)
		}
	}

	return nil
}
