# Implementation Plan: Resgate Listener Struct

**Branch**: `100-create_resgate_listener` | **Date**: 2026-04-06 | **Spec**: [spec.md](file:///Users/reidlai/GitLocal/resgate-listener/specs/100-create_resgate_listener/spec.md)
**Input**: Feature specification from `/specs/100-create_resgate_listener/spec.md`

## Summary

Build a dynamic, decoupled `ResgateListener` library component in the `pkg/listener` structure. It serves as an infrastructural layer consuming RES-compliant NATS messages and routing them securely to dependency-injected module handlers while insulating the system from panics.

## User Stories

- **As a** module developer, **I want to** route specific NATS topics to my handlers **so that** I don't have to rewrite boilerplate subscription logic for every module.
  - **Technical Delivery**: Implement an exported `ResgateListener` tracking logic in `pkg/listener` constructed via dependency injection (`map[string]MessageHandler`). Topics passed in are registered as `subs []string` for visibility, while NATS Subscriptions are tracked in an unexported map to manage reliable teardowns.
  - **Acceptance**: A mapped string matching an active subscription triggers the handler; errors or panics within a handler trigger `defer recover()` natively.

## Technical Context

**Language/Version**: Golang 1.25+
**Primary Dependencies**: `nats-io/nats.go`, `jirenius/go-res`
**Storage**: N/A (In-memory router layer)
**Testing**: testing, nats-server/v2/test embedded servers
**Target Platform**: distroless Docker image (for eventual consumers)
**Project Type**: library package (`pkg/`) within a go-res consumer web-service
**Performance Goals**: N/A natively, dictated by inner handler execution logic.
**Constraints**: Zero-downtime routing. Panics inside mapped handlers MUST not break sibling routes.
**Scale/Scope**: Intended to be bound dynamically upon service boot.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] **Containers**: Uses distroless base image with multi-stage builds. (Consumer dependency check)
- [x] **Containers**: Runs as non-root, minimal capabilities, and includes explicit HEALTHCHECK.
- [x] **Dependencies**: Uses permissive licenses (MIT, Apache-2.0, BSD) and has zero unaddressed critical CVEs.
- [x] **DevSecOps**: PR pipeline includes SCA, linting, secrets scanning, SAST, Trivy, and unit tests.
- [x] **DevSecOps**: CD pipelines per environment incorporate DAST, post-deploy integration, and Cucumber smoke tests.
- [x] **Engineering**: Adheres to Twelve-Factor and SOLID principles, or includes an approved waiver for deviation.
- [x] **Project Structure**: Follows the `golang-standards/project-layout` conventions.
- [x] **Planning**: Ensures all requirements are explicitly delivered via mapped User/Developer Stories within this plan.

## Project Structure

### Documentation (this feature)

```text
specs/100-create_resgate_listener/
├── plan.md              # This file
├── research.md          # Architecture and constraints research
├── data-model.md        # The Listener Struct / MessageHandler details
├── quickstart.md        # Docs on instantiation
├── contracts/           # N/A
└── tasks.md             # To be populated heavily via /speckit-tasks
```

### Source Code

```text
pkg/
└── listener/            # Library domain for Resgate connection handling
    ├── listener.go      # Core ResgateListener & Context
    └── listener_test.go # TDD test specs including embedded NATS
```

**Structure Decision**: The implementation fundamentally lives strictly encapsulated inside `pkg/` alongside its test matrix allowing domain separation from the main executable blocks.
