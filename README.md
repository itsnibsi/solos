# Solos Project

Solos is a scalable, versioned, properly protected and dockerized JSON backend API with a React frontend, designed for easy local development and production deployments.

## Tech Stack

### Backend

- Go 1.22
- Fiber framework
- Viper for configuration
- Swagger for API documentation
- golangci-lint for linting

### Frontend

- React 18.3
- TypeScript 5.22
- Vite 5.3.4
- ESLint and Prettier for linting and formatting

### Development Tools

- Docker and Docker Compose
- Air for hot reloading (Go backend)
- Make for running common commands

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.22
- Node.js and npm
- Make

### Setup

1. Clone the repository:

```
git clone https://github.com/itsnibsi/solos.git
cd solos
```

2. Copy the example environment files:

```
cp .env.example .env
cp client/.env.example client/.env
```

3. Build and run the development environment:

```
make run-dev
```

This will start both the backend and frontend in development mode with hot reloading.

## Available Commands

- `make run-dev`: Start the development environment
- `make lint-backend`: Run golangci-lint on the backend code
- `make lint-frontend`: Run ESLint on the frontend code
- `make test-backend`: Run backend tests
- `make generate-swagger`: Generate Swagger documentation
- `make build-dev`: Build the development Docker images

## Project Structure

```
solos/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── handlers/
│   ├── config/
│   ├── errors/
│   └── server/
├── client/
│   ├── src/
│   ├── public/
│   └── package.json
├── docs/
├── .vscode/
├── docker-compose.yml
├── docker-compose.dev.yml
├── Dockerfile
├── Dockerfile.dev
├── Makefile
├── .env
└── README.md
```

## API Documentation

After running `make generate-swagger`, you can find the Swagger documentation at `/docs/swagger.json`. You can use tools like Swagger UI to visualize and interact with the API documentation.

## Contributing

TBD.

## License

TBD.
