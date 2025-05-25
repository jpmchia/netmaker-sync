// cmd/netmaker-sync/main.go
package main

import (
	"context"
	"netmaker-sync/internal/api"
	"netmaker-sync/internal/config"
	"netmaker-sync/internal/db"
	"netmaker-sync/internal/service"
	"netmaker-sync/internal/sync"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setLogLevel sets the log level based on the provided string
func setLogLevel(level string) {
	switch level {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warnf("Invalid log level '%s', defaulting to 'info'", level)
	}
	logrus.Infof("Log level set to '%s'", logrus.GetLevel().String())
}

func main() {
	// Set up default logging
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	
	// Default to info level until we load the config
	logrus.SetLevel(logrus.InfoLevel)

	rootCmd := &cobra.Command{
		Use:   "netmaker-sync",
		Short: "Netmaker API synchronization daemon",
	}

	serveCommand := serveCmd()
	rootCmd.AddCommand(serveCommand)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func serveCmd() *cobra.Command {
	var logLevel string
	
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the Netmaker sync daemon",
		Run: func(cmd *cobra.Command, args []string) {
			// Load config
			cfg, err := config.Load()
			if err != nil {
				logrus.Fatal(err)
			}
			
			// Set log level from config if not overridden by flag
			if !cmd.Flags().Changed("log-level") {
				logLevel = cfg.Logging.Level
			}
			setLogLevel(logLevel)

			// Initialize database
			database, err := db.New(&cfg.Database)
			if err != nil {
				logrus.Fatal(err)
			}

			if err := database.Initialize(); err != nil {
				logrus.Fatal(err)
			}

			// Initialize API client
			apiClient := api.New(&cfg.NetmakerAPI, &cfg.Logging)

			// Initialize sync service
			syncService := sync.New(apiClient, database)

			// Initialize HTTP server
			server := service.New(syncService)

			// Initialize cron scheduler
			c := cron.New()
			_, err = c.AddFunc("@every "+cfg.Sync.Interval.String(), func() {
				if err := syncService.SyncAll(context.Background()); err != nil {
					logrus.Errorf("Scheduled sync failed: %v", err)
				}
			})
			if err != nil {
				logrus.Fatal(err)
			}
			c.Start()

			// Start HTTP server in a goroutine
			go func() {
				logrus.Infof("Starting HTTP server on %s:%d", cfg.API.Host, cfg.API.Port)
				if err := server.Start(cfg.API.Host, cfg.API.Port); err != nil {
					logrus.Fatal(err)
				}
			}()

			// Wait for interrupt signal
			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
			<-sigCh

			// Shutdown gracefully
			logrus.Info("Shutting down...")
			c.Stop()

			// Allow some time for ongoing operations to complete
			time.Sleep(1 * time.Second)
		},
	}

	cmd.Flags().StringVarP(&logLevel, "log-level", "l", "info", "Set the log level (trace, debug, info, warn, error, fatal)")
	return cmd
}
