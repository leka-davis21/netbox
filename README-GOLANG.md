# Golang Gin API Server

A simple single-file web API server built with Go and the Gin framework.

## Features

- 🚀 Lightweight single-file application
- 📦 Built with [Gin Web Framework](https://github.com/gin-gonic/gin)
- ✅ GitHub Actions CI/CD pipeline
- 🔧 Three endpoints ready to use

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

Clone the repository:

```bash
git clone https://github.com/leka-davis21/netbox.git
cd netbox
```

Install dependencies:

```bash
go mod download
```

## Usage

### Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Building the Application

```bash
go build -o app main.go
./app
```

## API Endpoints

### 1. Root Endpoint
```bash
GET /
```

Response:
```json
{
  "message": "Welcome to Golang Gin API",
  "status": "running"
}
```

### 2. Health Check
```bash
GET /health
```

Response:
```json
{
  "status": "healthy"
}
```

### 3. Hello Endpoint
```bash
GET /api/hello/:name
```

Example:
```bash
curl http://localhost:8080/api/hello/World
```

Response:
```json
{
  "message": "Hello, World!"
}
```

## Testing

Run the tests:

```bash
go test -v ./...
```

## GitHub Actions

This project includes a CI/CD pipeline that:
- Builds the application
- Runs tests with race detection
- Performs static code analysis
- Runs linting checks

The workflow is triggered on:
- Push to `main`, `master`, `develop`, or `copilot/*` branches
- Pull requests to `main`, `master`, or `develop` branches

## Project Structure

```
.
├── main.go                      # Single-file application
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies checksums
├── README-GOLANG.md            # This file
└── .github/
    └── workflows/
        └── golang.yml          # GitHub Actions workflow
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE.txt](LICENSE.txt) file for details.
