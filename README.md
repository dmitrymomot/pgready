# pgready

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/dmitrymomot/pgready)](https://github.com/dmitrymomot/pgready)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/pgready)](https://goreportcard.com/report/github.com/dmitrymomot/pgready)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/pgready.svg)](https://pkg.go.dev/github.com/dmitrymomot/pgready)
[![License](https://img.shields.io/github/license/dmitrymomot/pgready)](https://github.com/dmitrymomot/pgready/blob/main/LICENSE)

A simple command-line utility to check if a PostgreSQL database is ready to accept connections. This tool is particularly useful in containerized environments and deployment scripts where you need to ensure the database is available before starting your application.

## Features

-   Configurable connection retry attempts
-   Adjustable retry interval
-   Supports connection string via command line flag or environment variable
-   Returns appropriate exit codes (0 for success, 1 for failure)

## Installation

```bash
go install github.com/yourusername/pgready@latest
```

## Usage

### Basic Usage

```bash
# Using command line flags
pgready --db-url "postgres://user:password@localhost:5432/dbname"

# Using environment variable
export DATABASE_URL="postgres://user:password@localhost:5432/dbname"
pgready
```

### Command Line Options

-   `--db-url`: PostgreSQL connection string (can also be set via DATABASE_URL environment variable)
-   `--retry-attempts`: Number of connection retry attempts (default: 20)
-   `--retry-interval`: Interval between retry attempts (default: 1s)

### Examples

```bash
# Custom retry attempts and interval
pgready --db-url "postgres://user:password@localhost:5432/dbname" --retry-attempts 30 --retry-interval 2s

# Wait for PostgreSQL in a Docker Compose setup
pgready --db-url "postgres://postgres:postgres@db:5432/myapp?sslmode=disable" --retry-interval 5s

# Use in a shell script
if pgready --db-url "postgres://localhost/app"; then
    echo "Database is ready"
    ./start-application
else
    echo "Database failed to become ready"
    exit 1
fi
```

## Exit Codes

-   `0`: PostgreSQL is ready for connections
-   `1`: Failed to connect to PostgreSQL within the given retry attempts or invalid configuration

## License

This project is licensed under the [Apache 2.0](LICENSE) - see the LICENSE file for details.
