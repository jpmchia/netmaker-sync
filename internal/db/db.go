package db

import (
	"fmt"
	"netmaker-sync/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
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
	return db.setupSchema()
}

// setupSchema sets up the database schema
func (db *DB) setupSchema() error {
	// // Drop existing tables if they exist to start fresh
	// _, err := db.Exec(`
	// 	DROP TABLE IF EXISTS acls CASCADE;
	// 	DROP TABLE IF EXISTS dns_entries CASCADE;
	// 	DROP TABLE IF EXISTS hosts CASCADE;
	// 	DROP TABLE IF EXISTS ext_clients CASCADE;
	// 	DROP TABLE IF EXISTS nodes CASCADE;
	// 	DROP TABLE IF EXISTS networks CASCADE;
	// 	DROP TABLE IF EXISTS sync_history CASCADE;
	// 	DROP TABLE IF EXISTS schema_migrations CASCADE;
	// `)

	// if err != nil {
	// 	return fmt.Errorf("failed to drop existing tables: %w", err)
	// }

	// Create the schema_migrations table
	_, err = db.Exec(`
		CREATE TABLE schema_migrations (
			version INTEGER PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// Create the networks table
	_, err = db.Exec(`
		CREATE TABLE networks (
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
		);

		-- Add a unique constraint on networks.id for foreign key references
		ALTER TABLE networks ADD CONSTRAINT networks_id_unique UNIQUE (id);
	`)
	if err != nil {
		return fmt.Errorf("failed to create networks table: %w", err)
	}

	// Create the nodes table with a unique constraint on id for foreign key references
	_, err = db.Exec(`
		CREATE TABLE nodes (
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
		);
		
		-- Add a unique constraint on nodes.id for foreign key references
		ALTER TABLE nodes ADD CONSTRAINT nodes_id_unique UNIQUE (id);
	`)
	if err != nil {
		return fmt.Errorf("failed to create nodes table: %w", err)
	}

	// Create the ext_clients table
	_, err = db.Exec(`
		CREATE TABLE ext_clients (
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
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create ext_clients table: %w", err)
	}

	// Create the dns_entries table
	_, err = db.Exec(`
		CREATE TABLE dns_entries (
			id TEXT NOT NULL,
			version INTEGER NOT NULL,
			network_id TEXT NOT NULL REFERENCES networks(id) ON DELETE CASCADE,
			name TEXT NOT NULL,
			address TEXT,
			address6 TEXT,
			is_current BOOLEAN NOT NULL DEFAULT TRUE,
			last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			PRIMARY KEY (id, version),
			UNIQUE(network_id, name, version)
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create dns_entries table: %w", err)
	}

	// Create the hosts table
	_, err = db.Exec(`
		CREATE TABLE hosts (
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
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create hosts table: %w", err)
	}

	// Create the acls table with versioning
	_, err = db.Exec(`
		CREATE TABLE acls (
			id INTEGER NOT NULL,
			version INTEGER NOT NULL,
			network_id TEXT NOT NULL REFERENCES networks(id) ON DELETE CASCADE,
			node_id TEXT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
			is_current BOOLEAN NOT NULL DEFAULT TRUE,
			data JSONB NOT NULL,
			last_modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			PRIMARY KEY (id, version)
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create acls table: %w", err)
	}

	// Create the sync_history table
	_, err = db.Exec(`
		CREATE TABLE sync_history (
			id SERIAL PRIMARY KEY,
			resource_type TEXT NOT NULL,
			status TEXT NOT NULL,
			message TEXT,
			started_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			completed_at TIMESTAMP WITH TIME ZONE
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create sync_history table: %w", err)
	}

	// Record that we've applied all migrations
	_, err = db.Exec(`INSERT INTO schema_migrations (version) VALUES (1)`)
	if err != nil {
		return fmt.Errorf("failed to record migration: %w", err)
	}

	logrus.Info("Database schema setup complete")
	return nil
}

// migration represents a database migration
type migration struct {
	version int
	up      string
	down    string
}

// runMigrations runs any pending migrations
func (db *DB) runMigrations(migrations []migration) error {
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
