// service/server.go
package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"netmaker-sync/internal/sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Server represents the HTTP API server
type Server struct {
	router      *chi.Mux
	syncService *sync.Service
}

// New creates a new HTTP API server
func New(syncService *sync.Service) *Server {
	s := &Server{
		router:      chi.NewRouter(),
		syncService: syncService,
	}

	s.setupRoutes()
	return s
}

// setupRoutes sets up the HTTP routes
func (s *Server) setupRoutes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// CORS middleware
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// API routes
	s.router.Route("/api", func(r chi.Router) {
		// Sync routes
		r.Route("/sync", func(r chi.Router) {
			r.Post("/", s.handleSyncAll)
			r.Post("/networks", s.handleSyncNetworks)
			r.Post("/networks/{networkID}/nodes", s.handleSyncNodes)
			// More sync routes
		})

		// Data routes
		r.Route("/data", func(r chi.Router) {
			r.Get("/networks", s.handleGetNetworks)
			r.Get("/networks/{networkID}", s.handleGetNetwork)
			// More data routes
		})
	})
}

// Start starts the HTTP server
func (s *Server) Start(host string, port int) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), s.router)
}

// Handler implementations

// handleSyncAll handles a request to sync all resources
func (s *Server) handleSyncAll(w http.ResponseWriter, r *http.Request) {
	err := s.syncService.SyncAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success","message":"Sync completed successfully"}`)) 
}

// handleSyncNetworks handles a request to sync networks
func (s *Server) handleSyncNetworks(w http.ResponseWriter, r *http.Request) {
	err := s.syncService.SyncNetworks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success","message":"Networks sync completed successfully"}`)) 
}

// handleSyncNodes handles a request to sync nodes for a specific network
func (s *Server) handleSyncNodes(w http.ResponseWriter, r *http.Request) {
	networkID := chi.URLParam(r, "networkID")
	if networkID == "" {
		http.Error(w, "Network ID is required", http.StatusBadRequest)
		return
	}

	err := s.syncService.SyncNodes(r.Context(), networkID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success","message":"Nodes sync completed successfully"}`)) 
}

// handleGetNetworks handles a request to get all networks
func (s *Server) handleGetNetworks(w http.ResponseWriter, r *http.Request) {
	networks, err := s.syncService.GetNetworks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(networks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

// handleGetNetwork handles a request to get a specific network
func (s *Server) handleGetNetwork(w http.ResponseWriter, r *http.Request) {
	networkID := chi.URLParam(r, "networkID")
	if networkID == "" {
		http.Error(w, "Network ID is required", http.StatusBadRequest)
		return
	}

	network, err := s.syncService.GetNetwork(r.Context(), networkID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(network)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
