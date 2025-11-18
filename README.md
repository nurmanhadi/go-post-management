# Post Management

A Go-based microservice responsible for managing posts, likes, and comments within a distributed microservices architecture.

## Tech Stack

- **Language:** Go 1.21+
- **Database:** PostgreSQL 15+
- **Cache:** Memcached
- **Message Broker:** LavinMQ / RabbitMQ
- **Containerization:** Docker & Docker Compose
- **Orchestration:** Kubernetes

## Prerequisites

- Go 1.21 or higher
- Docker & Docker Compose
- PostgreSQL 15+
- Memcached
- LavinMQ or RabbitMQ

## Quick Start

### 1. Clone the Repository

```bash
git https://github.com/nurmanhadi/go-post-management.git
cd post-management
```

### 2. Configure Environment Variables

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5434
DB_USERNAME=post
DB_PASSWORD=post
DB_NAME=post_management

# Cache Configuration
CACHE_HOST=localhost
CACHE_PORT=11213

# Message Broker Configuration
BROKER_HOST=localhost
BROKER_PORT=5672
BROKER_USERNAME=guest
BROKER_PASSWORD=guest
BROKER_VHOST=someone

# api user
API_USER=http://localhost:3001
```

### 3. Start Services with Docker Compose

```bash
docker-compose up -d
```

### 4. Run Database Migrations

```bash
make migrate-up
```

## Development

### Run Locally

```bash
go run cmd/main.go
```

### Build Binary

```bash
go build -o bin/user-service cmd/main.go
```

## Api Documentation
[Here!](docs/api.md)

## Security Considerations

- Enable TLS/SSL for production deployments
- Use HTTPS for all API endpoints
- Validate and sanitize all user inputs
- Implement rate limiting on authentication endpoints
- Use environment variables for sensitive configuration (API keys, database credentials)
- Implement proper authentication and authorization mechanisms
- Enable CORS only for trusted domains in production
- Keep dependencies updated regularly

## License

This project is licensed under the MIT License.

## Author

**Nurman Hadi**  
Backend Developer (Golang, Microservices)  
GitHub: [@nurmanhadi](https://github.com/nurmanhadi)