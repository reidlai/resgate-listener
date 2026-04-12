# Implementation Plan: Local Development Infrastructure (Compose-Only)

**Branch**: `101-docker-localdev-setup` | **Date**: 2026-04-12 | **Spec**: [spec.md](file:///Users/reidlai/GitLocal/resgate-listener/specs/101-docker-localdev-setup/spec.md)

## Summary

Provision a local development environment using Docker Compose that integrates NATS (with JetStream) and Resgate. This allows the `resgate-listener` application to connect to the required services immediately upon startup.

## User Stories

- **As a** developer, **I want to** run a single Docker Compose command **so that** I have a functional NATS/Resgate stack ready for local testing.
  - **Technical Delivery**: Create a `docker-compose.localdev.yaml` in `deploy/docker-compose/`.
  - **Acceptance**: `go run example/main.go start` connects to NATS at `localhost:4222` without error.

## Technical Context

**Primary Dependencies**: NATS 2.10, Resgate 1.7
**Storage**: N/A (In-memory JetStream)
**Target Platform**: Docker Compose
**Performance Goals**: < 10s startup time.

## Project Structure

### Documentation (this feature)

```text
specs/101-docker-localdev-setup/
├── spec.md              # Requirements
├── plan.md              # This file
├── quickstart.md        # Usage instructions
└── tasks.md             # Execution tracking
```

### Source Code

```text
deploy/
└── docker-compose/
    └── docker-compose.localdev.yaml # Primary infrastructure
```

## Proposed Changes

### [Component: Infrastructure]

#### [NEW] [docker-compose.localdev.yaml](file:///Users/reidlai/GitLocal/resgate-listener/deploy/docker-compose/docker-compose.localdev.yaml)
- Standard NATS 2.10 image with JetStream enabled.
- Standard Resgate 1.7 image.
- Health checks for both services to ensure readiness.

## Verification Plan

### Automated Tests
- `docker-compose -f deploy/docker-compose/docker-compose.localdev.yaml up -d`
- `go run example/main.go start` (Expect success)

### Manual Verification
- Verify NATS JetStream is active: `nats stream ls` (if nats-cli installed).
- Verify Resgate is reachable at `localhost:8080`.
