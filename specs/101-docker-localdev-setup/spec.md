# Feature Specification: Local Development Infrastructure Setup

**Feature Branch**: `101-docker-localdev-setup`  
**Created**: 2026-04-12  
**Status**: Draft  
**Input**: User description: "Create local development stack using Docker Compose for NATS JetStream and Resgate."

## Clarifications

### Session 2026-04-12
- Q: Which Resgate and NATS versions should be used? → A: Latest Stable (NATS 2.10, Resgate 1.7).
- Q: How should "Healthy" be defined for the services? → A: All services in the compose stack must pass their respective health checks.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Local Development Service Provisioning (Priority: P1)

As a developer, I want to quickly start a local environment containing NATS (with JetStream) and Resgate so that I can run and test the `resgate-listener` example application without manual service installation.

**Why this priority**: Essential for the "getting started" experience and local verification of feature changes. It unblocks the developer from the "no servers available" error.

**Independent Test**: Execution of `go run example/main.go start` succeeds in connecting to NATS after the infrastructure is started via Docker Compose.

**Acceptance Scenarios**:

1. **Given** a fresh project checkout, **When** I start the services using `docker-compose -f deploy/docker-compose/docker-compose.localdev.yaml up -d`, **Then** a NATS server with JetStream is available at `localhost:4222`.
2. **Given** the infrastructure is running, **When** I run `go run example/main.go start`, **Then** the application logs "Starting dummy server... nats.url=nats://localhost:4222" followed by a successful execution (no connection error).
3. **Given** the infrastructure is active, **When** I check for Resgate, **Then** the Resgate service is also running and correctly connected to the internal NATS instance.

---

### Edge Cases

- **Service Readiness**: What happens if Resgate starts before NATS is ready? (Logic: Compose `depends_on` with `service_healthy` handles this).
- **Port Conflicts**: How does the system handle if port 4222 is already in use on the host? (Standard Docker port mapping error handled by user).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST provide a `docker-compose.localdev.yaml` in the `./deploy/docker-compose/` folder.
- **FR-002**: The stack MUST include a NATS server (v2.10+) with JetStream enabled (`-js` flag).
- **FR-003**: The stack MUST include a Resgate service (v1.7+) configured to use the internal NATS server.
- **FR-004**: The services MUST be accessible from the host machine (localhost).
- **FR-005**: All services MUST include `healthcheck` definitions.

### Key Entities

- **Developer Environment**: The local machine where the Go code is executed.
- **Orchestration Stack**: The set of containers managed by Docker Compose.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Zero "nats: no servers available for connection" errors recorded when the compose stack is active.
- **SC-002**: The entire stack reaches "Healthy" status in under 10 seconds.
- **SC-003**: The entire setup is reproducible using a single `docker-compose` command.

## Assumptions

- **Default Ports**: NATS will listen on default port 4222; Resgate will listen on its default port 8080.
- **Host Execution**: The Go application is run on the host OS, not inside a container.
- **No Persistence Required**: Transient data (in-memory JetStream) is acceptable.
