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
	Logging     LoggingConfig
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
	Interval    time.Duration
	IncludeAcls bool
}

// APIConfig holds API server specific configuration
type APIConfig struct {
	Host string
	Port int
}

// LoggingConfig holds logging specific configuration
type LoggingConfig struct {
	Level             string
	DisableRestyDebug bool
}

// Load loads the configuration from environment variables and config file
func Load() (*Config, error) {
	// Load .env file if it exists
	envErr := godotenv.Load()
	if envErr != nil {
		logrus.Infof(".env file not loaded: %v", envErr)
	} else {
		logrus.Info(".env file loaded successfully")
	}

	// Set up viper to handle both YAML config and environment variables
	// Try to load config file from multiple locations
	viper.SetConfigName("config")              // name of config file (without extension)
	viper.SetConfigType("yaml")                // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")                   // look for config in the working directory
	viper.AddConfigPath("..")                  // look for config in parent directory
	viper.AddConfigPath("/etc/netmaker-sync/") // path to look for the config file in

	// Set default values
	viper.SetDefault("netmaker_api.url", "https://api.netmaker.example.com")
	viper.SetDefault("netmaker_api.key", "")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.name", "netmaker_sync")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("sync.interval", "5m")
	viper.SetDefault("api.host", "0.0.0.0")
	viper.SetDefault("api.port", 8080)
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.disable_resty_debug", true)

	// Map environment variables to viper keys
	viper.BindEnv("netmaker_api.url", "NETMAKER_API_URL")
	viper.BindEnv("netmaker_api.key", "NETMAKER_API_KEY")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("sync.interval", "SYNC_INTERVAL")
	viper.BindEnv("api.host", "API_HOST")
	viper.BindEnv("api.port", "API_PORT")
	viper.BindEnv("logging.level", "LOG_LEVEL")
	viper.BindEnv("logging.disable_resty_debug", "DISABLE_RESTY_DEBUG")

	// Enable environment variables
	viper.AutomaticEnv()

	// Try to read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("No config file found, using environment variables and defaults")
		} else {
			logrus.Warnf("Error reading config file: %v", err)
		}
	} else {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
	}

	syncInterval, err := time.ParseDuration(viper.GetString("sync.interval"))
	if err != nil {
		logrus.Warnf("Invalid sync.interval: %s, using default 5m", viper.GetString("sync.interval"))
		syncInterval = 5 * time.Minute
	}

	// Log the configuration values for debugging
	logrus.Debugf("Configuration loaded: netmaker_api.url=%s, database.host=%s, database.name=%s",
		viper.GetString("netmaker_api.url"),
		viper.GetString("database.host"),
		viper.GetString("database.name"))

	return &Config{
		NetmakerAPI: NetmakerAPIConfig{
			URL: viper.GetString("netmaker_api.url"),
			Key: viper.GetString("netmaker_api.key"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			Name:     viper.GetString("database.name"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
		},
		Sync: SyncConfig{
			Interval:    syncInterval,
			IncludeAcls: viper.GetBool("sync.include_acls"),
		},
		API: APIConfig{
			Host: viper.GetString("api.host"),
			Port: viper.GetInt("api.port"),
		},
		Logging: LoggingConfig{
			Level:             viper.GetString("logging.level"),
			DisableRestyDebug: viper.GetBool("logging.disable_resty_debug"),
		},
	}, nil
}
