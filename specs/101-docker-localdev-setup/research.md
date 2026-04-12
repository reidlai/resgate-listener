# Research: Local Development Infrastructure

## Decision: Multi-service Dockerfile with Go-based Supervisor

### Rationale
To meet the project constitution requirements (Distroless, Non-root, Healthchecks) while fulfilling the user's request for a single `Dockerfile.localdev`, a custom Go-based supervisor binary is the most compliant and robust solution. Standard shell scripts or process managers like `supervisord` require a shell (Alpine/Debian), which violates the Distroless requirement.

### Alternatives Considered
1. **Alpine-based Dockerfile**: Rejected for production-like compliance (constitution rule 52). While easier for dev, maintaining the same security posture in dev tools is a project principle.
2. **Docker Compose**: Recommended as a parallel option. It is the industrial standard for multiple services. I will provide `docker-compose.localdev.yaml` alongside the Dockerfile.

## Findings

### NATS JetStream
Enabled via the `-js` flag. To listen on all interfaces (required for host access), `-a 0.0.0.0` should be used.

### Resgate Configuration
Resgate can be configured via command-line arguments:
`--nats nats://localhost:4222`
Since both run in the same container, they share the loopback interface (`localhost`).

### Local Dev Supervisor
A minimal Go program will:
1. Start `nats-server -js`.
2. Wait for NATS to be ready (optional but good).
3. Start `resgate --nats nats://localhost:4222`.
4. Forward signals (SIGTERM/SIGINT) to both processes.
5. Exit if either process fails.

## Implementation Details

### Multi-stage Build
- **Stage 1 (NATS)**: `FROM nats:latest AS nats`
- **Stage 2 (Resgate)**: `FROM resgateio/resgate:latest AS resgate`
- **Stage 3 (Supervisor Build)**: `FROM golang:alpine AS builder` (Builds the supervisor)
- **Stage 4 (Final)**: `FROM gcr.io/distroless/static`
  - Copy `nats-server` from stage 1.
  - Copy `resgate` from stage 2.
  - Copy `supervisor` from stage 3.
  - Set `USER nonroot`.
  - Set `ENTRYPOINT ["/supervisor"]`.
