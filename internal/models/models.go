package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// JSONB is a wrapper around map[string]interface{} to implement the driver.Valuer and sql.Scanner interfaces
type JSONB map[string]interface{}

// Value implements the driver.Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan implements the sql.Scanner interface
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	
	return json.Unmarshal(bytes, j)
}

// Network represents a Netmaker network
type Network struct {
	ID                     string    `json:"id" db:"id"`
	Name                   string    `json:"name" db:"name"`
	AddressRange           string    `json:"addressrange" db:"address_range"`
	AddressRange6          string    `json:"addressrange6" db:"address_range6"`
	LocalRange             string    `json:"localrange" db:"local_range"`
	IsDualStack            bool      `json:"isdualstack" db:"is_dual_stack"`
	IsIPv4                 bool      `json:"isipv4" db:"is_ipv4"`
	IsIPv6                 bool      `json:"isipv6" db:"is_ipv6"`
	IsLocal                bool      `json:"islocal" db:"is_local"`
	DefaultAccessControl   string    `json:"defaultacl" db:"default_access_control"`
	DefaultUDPHolePunching bool      `json:"defaultudphp" db:"default_udp_hole_punching"`
	DefaultExtClientDNS    string    `json:"defaultextclientdns" db:"default_ext_client_dns"`
	DefaultMTU             int       `json:"defaultmtu" db:"default_mtu"`
	DefaultKeepalive       int       `json:"defaultkeepalive" db:"default_keepalive"`
	DefaultInterface       string    `json:"defaultinterface" db:"default_interface"`
	NodeLimit              int       `json:"nodelimit" db:"node_limit"`
	LastModified           time.Time `json:"lastmodified" db:"last_modified"`
	CreatedAt              time.Time `json:"created_at" db:"created_at"`
	Data                   JSONB     `json:"data" db:"data"`
}

// Node represents a Netmaker node
type Node struct {
	ID               string    `json:"id" db:"id"`
	NetworkID        string    `json:"network" db:"network_id"`
	Name             string    `json:"name" db:"name"`
	Address          string    `json:"address" db:"address"`
	Address6         string    `json:"address6" db:"address6"`
	PublicKey        string    `json:"publickey" db:"public_key"`
	Endpoint         string    `json:"endpoint" db:"endpoint"`
	IsEgressGateway  bool      `json:"isegressgateway" db:"is_egress_gateway"`
	IsIngressGateway bool      `json:"isingressgateway" db:"is_ingress_gateway"`
	IsRelay          bool      `json:"isrelay" db:"is_relay"`
	Connected        bool      `json:"connected" db:"connected"`
	LastModified     time.Time `json:"lastmodified" db:"last_modified"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	Data             JSONB     `json:"data" db:"data"`
}

// ExtClient represents a Netmaker external client
type ExtClient struct {
	ID           string    `json:"clientid" db:"id"`
	NetworkID    string    `json:"network" db:"network_id"`
	Name         string    `json:"name" db:"name"`
	Address      string    `json:"address" db:"address"`
	Address6     string    `json:"address6" db:"address6"`
	PublicKey    string    `json:"publickey" db:"public_key"`
	Enabled      bool      `json:"enabled" db:"enabled"`
	LastModified time.Time `json:"lastmodified" db:"last_modified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	Data         JSONB     `json:"data" db:"data"`
}

// DNSEntry represents a Netmaker DNS entry
type DNSEntry struct {
	ID           int       `json:"id" db:"id"`
	NetworkID    string    `json:"network" db:"network_id"`
	Name         string    `json:"name" db:"name"`
	Address      string    `json:"address" db:"address"`
	Address6     string    `json:"address6" db:"address6"`
	LastModified time.Time `json:"lastmodified" db:"last_modified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// Host represents a Netmaker host
type Host struct {
	ID                  string    `json:"id" db:"id"`
	Name                string    `json:"name" db:"name"`
	EndpointIP          string    `json:"endpointip" db:"endpoint_ip"`
	EndpointIPv6        string    `json:"endpointipv6" db:"endpoint_ipv6"`
	PublicKey           string    `json:"publickey" db:"public_key"`
	ListenPort          int       `json:"listenport" db:"listen_port"`
	MTU                 int       `json:"mtu" db:"mtu"`
	PersistentKeepalive int       `json:"persistentkeepalive" db:"persistent_keepalive"`
	LastModified        time.Time `json:"lastmodified" db:"last_modified"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	Data                JSONB     `json:"data" db:"data"`
}

// ACL represents a Netmaker ACL
type ACL struct {
	ID           int       `json:"id" db:"id"`
	NetworkID    string    `json:"network" db:"network_id"`
	NodeID       string    `json:"nodeid" db:"node_id"`
	Data         JSONB     `json:"data" db:"data"`
	LastModified time.Time `json:"lastmodified" db:"last_modified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// SyncHistory represents a record of a sync operation
type SyncHistory struct {
	ID          int        `json:"id" db:"id"`
	ResourceType string     `json:"resource_type" db:"resource_type"`
	Status      string     `json:"status" db:"status"`
	Message     string     `json:"message" db:"message"`
	StartedAt   time.Time  `json:"started_at" db:"started_at"`
	CompletedAt *time.Time `json:"completed_at" db:"completed_at"`
}

// SyncStatus constants
const (
	SyncStatusPending   = "pending"
	SyncStatusCompleted = "completed"
	SyncStatusFailed    = "failed"
)

// ResourceType constants
const (
	ResourceTypeNetwork   = "network"
	ResourceTypeNode      = "node"
	ResourceTypeExtClient = "ext_client"
	ResourceTypeDNS       = "dns"
	ResourceTypeHost      = "host"
	ResourceTypeACL       = "acl"
)
