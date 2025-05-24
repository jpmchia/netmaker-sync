package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	NetmakerAPI NetmakerAPIConfig
	Database    DatabaseConfig
	Sync        SyncConfig
	API         APIConfig
}

// NetmakerAPIConfig holds Netmaker API specific configuration
type NetmakerAPIConfig struct {
	URL string
	Key string
}

// DatabaseConfig holds database specific configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// SyncConfig holds synchronization specific configuration
type SyncConfig struct {
	Interval time.Duration
}

// APIConfig holds API server specific configuration
type APIConfig struct {
	Host string
	Port int
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	viper.SetDefault("NETMAKER_API_URL", "https://api.netmaker.example.com")
	viper.SetDefault("NETMAKER_API_KEY", "")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_NAME", "netmaker_sync")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("SYNC_INTERVAL", "5m")
	viper.SetDefault("API_HOST", "0.0.0.0")
	viper.SetDefault("API_PORT", 8080)

	viper.AutomaticEnv()

	syncInterval, err := time.ParseDuration(viper.GetString("SYNC_INTERVAL"))
	if err != nil {
		logrus.Warnf("Invalid SYNC_INTERVAL: %s, using default 5m", viper.GetString("SYNC_INTERVAL"))
		syncInterval = 5 * time.Minute
	}

	return &Config{
		NetmakerAPI: NetmakerAPIConfig{
			URL: viper.GetString("NETMAKER_API_URL"),
			Key: viper.GetString("NETMAKER_API_KEY"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			Name:     viper.GetString("DB_NAME"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
		},
		Sync: SyncConfig{
			Interval: syncInterval,
		},
		API: APIConfig{
			Host: viper.GetString("API_HOST"),
			Port: viper.GetInt("API_PORT"),
		},
	}, nil
}
