# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

httpecho is a simple HTTP web server designed for testing purposes. It provides two main functionalities:

- Echo HTTP request headers back to clients (sorted alphabetically)
- Simulate slow/long-running HTTP responses with configurable delays

Built with Go 1.23 and uses Go modules for dependency management (though currently has no external dependencies).

## Development Commands

**Build and run:**

```bash
make build  # Build Docker image tagged for ghcr.io
make run    # Run container locally on port 8080
make push   # Push image to GitHub Container Registry (ghcr.io)
```

**Direct Docker usage:**

```bash
docker run -it -p 8080:8080 ghcr.io/ianneub/httpecho:latest
```

**CI/CD:**

- GitHub Actions workflow automatically builds and pushes images to ghcr.io on:
  - Pushes to `master` branch (tagged as `latest`)
  - New version tags (e.g., `v1.0.0`)
- Pull requests trigger builds but don't push images

## Code Architecture

The codebase uses a handler-based architecture with separate files for different endpoints:

- **[server.go](server.go)** - Main entry point that registers HTTP handlers and starts the server on port 8080
- **[echo.go](echo.go)** - Implements `echoHandler()` for the root endpoint (`/`), which logs requests and returns all HTTP headers sorted alphabetically
- **[long.go](long.go)** - Implements `longHandler()` for the `/long` endpoint, which simulates slow responses with a configurable timeout (query parameter `s` in seconds, default 300s/5min)

**Handler registration flow:**

```
main() → Register handlers → Start HTTP server
  ├─ / (catch-all) → echoHandler
  └─ /long → longHandler
```

## Key Design Patterns

- **Standard library only**: No external dependencies beyond Go standard library
- **Stateless**: Pure request/response model with no persistence or state
- **Docker-focused**: Uses multi-stage Docker build (Go 1.23-alpine for building, alpine:latest for runtime) resulting in ~15MB images
- **Separation of concerns**: One file per handler for clear organization
- **Developer-friendly logging**: All requests logged to console with method, URL, and remote address

## Build Process

The Dockerfile uses a multi-stage build:

1. **Build stage**: Compiles the Go binary using `golang:1.23-alpine`
2. **Runtime stage**: Copies only the binary to a minimal `alpine:latest` image

This approach results in small, secure images (~15MB) compared to including the full Go toolchain.
