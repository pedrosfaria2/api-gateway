# API Gateway

A modern, high-performance API Gateway built with Go, designed with Clean Architecture principles. This project aims to provide a flexible, modular, and extensible solution for API management and routing.

## Key Features

- **Clean Architecture**: Clear separation of concerns with a layered architecture
- **Dynamic Configuration**: Easy configuration via YAML/JSON without requiring recompilation
- **High Performance**: Built for high concurrency and low latency
- **Extensible Plugin System**: Easy to add new middlewares and functionalities
- **Monitoring & Metrics**: Built-in support for metrics and monitoring
- **Modern Development**: Docker support, hot-reloading, and comprehensive testing

## Architecture

The project follows Clean Architecture principles with distinct layers:

- **Core**: Contains the central business logic and domain rules
- **Infrastructure**: Implements external tool integrations
- **Plugins**: Modular system for extending functionality
- **Configuration**: Manages system settings and routing rules

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Make

### Development Setup

1. Install dependencies:
```bash
make install-deps
make setup
make setup-hooks
```

2. Run locally:
```bash
make run-dev  # Run with hot-reload
# or
make run      # Run without hot-reload
```

3. Run with Docker:
```bash
make run-docker
# or
make run-compose  # Run with Docker Compose
```

## Testing

```bash
make test            # Run tests
make test-coverage   # Run tests with coverage
```

## Code Quality

```bash
make lint     # Run linters
make format   # Format code
```


## Project Goals

1. **Extensibility**: Easy addition of new middlewares and backends without core changes
2. **Modularity**: Clear separation of concerns and responsibilities
3. **Performance**: Optimized for high throughput and low latency
4. **Developer Experience**: Simple configuration and deployment
5. **Reliability**: Comprehensive testing and monitoring

## Core Components

- Dynamic routing with configurable backends
- Middleware chain for request/response processing
- Plugin system for authentication, caching, and more
- Metrics and monitoring integration
- Request/response transformation capabilities
- Rate limiting and circuit breaking
