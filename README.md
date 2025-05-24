# NetmakerSync

A daemon service that efficiently syncs data from the Netmaker API to a PostgreSQL database, with a focus on change-based versioning and historical data tracking.

## Overview

NetmakerSync is designed to maintain a synchronized copy of Netmaker network data in a PostgreSQL database. It implements an efficient change-based versioning system that only stores meaningful changes while maintaining a complete history of all resources. This approach optimizes storage and provides valuable historical data for analysis and auditing.

## Features

- **Change-Based Versioning**: Only stores meaningful changes to resources
- **Historical Data**: Maintains a complete history of all resources
- **Sync History Tracking**: Records all sync operations with timestamps and status
- **RESTful API**: Provides HTTP endpoints to trigger syncs and retrieve data
- **Scheduled Syncs**: Automatically syncs data at configurable intervals

## Supported Resources

NetmakerSync can sync and version the following Netmaker resources:

- Networks
- Nodes
- External Clients
- DNS Entries
- Hosts
- Access Control Lists (ACLs)

## Requirements

- Go 1.24 or higher
- PostgreSQL 12 or higher
- Netmaker API access

## Installation

### From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/jpmchia/netmaker-sync.git
   cd netmaker-sync
   ```

2. Build the binary:
   ```bash
   go build
   ```

3. Run the binary:
   ```bash
   ./netmaker-sync serve
   ```

## Configuration

NetmakerSync uses a configuration file to specify connection details for the Netmaker API and PostgreSQL database. Create a `config.yaml` file in the same directory as the binary:

```yaml
netmaker_api:
  base_url: "https://your-netmaker-server.com/api"
  username: "your-username"
  password: "your-password"

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "postgres"
  name: "netmaker_sync"

api:
  host: "0.0.0.0"
  port: 8080

sync:
  interval: "5m"  # Sync interval in Go duration format (e.g., 1h, 30m, 5m)
```

## API Endpoints

NetmakerSync provides the following API endpoints:

- `POST /api/sync`: Sync all resources
- `POST /api/sync/networks`: Sync only networks
- `POST /api/sync/networks/{networkID}/nodes`: Sync nodes for a specific network
- `GET /api/data/networks`: Get all networks
- `GET /api/data/networks/{networkID}`: Get a specific network

## Database Schema

NetmakerSync creates the following tables in the PostgreSQL database:

- `networks`: Stores network data with versioning
- `nodes`: Stores node data with versioning
- `ext_clients`: Stores external client data with versioning
- `dns_entries`: Stores DNS entry data with versioning
- `hosts`: Stores host data with versioning
- `acls`: Stores ACL data with versioning
- `sync_history`: Tracks sync operations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
