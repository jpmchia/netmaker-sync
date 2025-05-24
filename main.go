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

func main() {
	rootCmd := &cobra.Command{
		Use:   "netmaker-sync",
		Short: "Netmaker API synchronization daemon",
	}

	rootCmd.AddCommand(serveCmd())

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func serveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the Netmaker sync daemon",
		Run: func(cmd *cobra.Command, args []string) {
			// Load config
			cfg, err := config.Load()
			if err != nil {
				logrus.Fatal(err)
			}

			// Initialize database
			database, err := db.New(&cfg.Database)
			if err != nil {
				logrus.Fatal(err)
			}

			if err := database.Initialize(); err != nil {
				logrus.Fatal(err)
			}

			// Initialize API client
			apiClient := api.New(&cfg.NetmakerAPI)

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

	return cmd
}
