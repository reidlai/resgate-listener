# Quickstart: Local Development Infrastructure

This guide explains how to use Docker Compose to set up your local development environment with NATS and Resgate.

## Prerequisites

- Docker
- Docker Compose

## Start the Infrastructure

Run the following command from the project root:

```bash
docker-compose -f deploy/docker-compose/docker-compose.localdev.yaml up -d
```

This starts:
- **NATS**: Listening on port `4222` (client) and `4223` (HTTP stats).
- **Resgate**: Listening on port `8080`.

## Verify Connectivity

Once the services are healthy, run the example application:

```bash
go run example/main.go start
```

The application should log `Starting dummy server...` and successfully connect to the NATS instance running in the Docker stack.

## Stopping the Services

To stop and remove the containers:

```bash
docker-compose -f deploy/docker-compose/docker-compose.localdev.yaml down
```
