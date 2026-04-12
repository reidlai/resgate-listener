# Implementation Plan: Adapt slog to OpenTelemetry

**Branch**: `102-slog-otel` | **Date**: 2026-04-12 | **Spec**: [spec.md](file:///Users/reidlai/GitLocal/resgate-listener/specs/102-slog-otel/spec.md)
**Input**: Feature specification from `/specs/102-slog-otel/spec.md`

## Summary

Integrate OpenTelemetry tracing correlation into `slog` logs by implementing a custom handler wrapper and updating the message processing pipeline to propagate `context.Context`. This ensure that logs generated during NATS message handling are automatically linked to their respective distributed traces.

## User Stories

- **As a** developer, **I want to** see `trace_id` and `span_id` in my logs **so that** I can debug asynchronous NATS message flows.
  - **Technical Delivery**: 
    1. Implement a custom `sdklog.Exporter` (`CompactExporter`) in `internal/otelutils` for single-line JSON.
    2. Use the official `otelslog.NewHandler` to bridge `slog` to the OTel Log SDK.
    3. Ensure logs follow the OpenTelemetry Log Data Model (OTLP).
  - **Acceptance**: Log output is single-line JSON but adheres to the OTel Log Record schema.

## Technical Context

**Language/Version**: Golang 1.22+
**Primary Dependencies**: `go.opentelemetry.io/contrib/bridges/otelslog`, `go.opentelemetry.io/otel`, `go.opentelemetry.io/otel/trace`
**Storage**: N/A
**Testing**: `testing`, `nats-server`
**Target Platform**: distroless Docker image
**Project Type**: core library (pkg/) and application logic (example/)
**Performance Goals**: <5% log serialization overhead.
**Constraints**: Context-aware logging required throughout the call stack.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] **Containers**: N/A for library, but `example/` uses distroless if containerized.
- [x] **Dependencies**: All OTel libraries are Apache-2.0.
- [x] **DevSecOps**: N/A for this internal feature.
- [x] **Engineering**: Adheres to SOLID and 12-Factor principles.
- [x] **Project Structure**: Updates `pkg/` and `internal/` according to golang-standards.
- [x] **Planning**: All requirements mapped, including clarified top-level field positioning.

## Project Structure

### Documentation (this feature)

```text
specs/102-slog-otel/
├── spec.md              # Requirements
├── plan.md              # This file
├── research.md          # Technology choices and NATS propagation patterns
├── data-model.md        # Log attribute definitions
└── quickstart.md        # How to enable OTel logging
```

### Source Code (repository root)

```text
internal/
└── otelutils/
    ├── propagation.go  # NATS TextMapCarrier implementation
    └── exporter.go     # [NEW] Custom OTel CompactExporter
pkg/
├── listener/
│   └── listener.go     # [MODIFY] Context extraction and propagation
└── resgate_message_handler/
    └── resgate_message_handler.go # [MODIFY] HandleMessage interface update
example/
└── cmd/
    └── start.go        # [MODIFY] Global slog and OTel SDK initialization
```

**Structure Decision**: Created `internal/otelutils` to keep the OTel propagation boilerplate separated from the core listener logic, following `golang-standards/project-layout`.

## Proposed Changes

### [Component: core-library]

#### [MODIFY] [resgate_message_handler.go](file:///Users/reidlai/GitLocal/resgate-listener/pkg/resgate_message_handler/resgate_message_handler.go)
- Add `context` import.
- Update `ResgateMessageHandler` interface: `HandleMessage(ctx context.Context, topic string, msg *nats.Msg)`.

#### [MODIFY] [listener.go](file:///Users/reidlai/GitLocal/resgate-listener/pkg/listener/listener.go)
- Import `go.opentelemetry.io/otel` and `internal/otelutils`.
- In `Listen()` loop:
  - Extract context from `msg.Header` using `otelutils.NATSHeaderCarrier`.
  - Pass extracted `ctx` to `handler.HandleMessage`.

### [Component: internal-utils]

#### [NEW] [propagation.go](file:///Users/reidlai/GitLocal/resgate-listener/internal/otelutils/propagation.go)
- Implement `NATSHeaderCarrier` satisfying `propagation.TextMapCarrier`.

### [Component: application]

#### [MODIFY] [start.go](file:///Users/reidlai/GitLocal/resgate-listener/example/cmd/start.go)
- Initialize OTel Trace and Log SDKs.
- Use `otelutils.CompactExporter` for log output.
- Initialize `slog` using `otelslog.NewHandler`.

## Verification Plan

### Automated Tests
- `go test ./pkg/listener/...`: Verify context is extracted from NATS headers.
- `go test ./pkg/resgate_message_handler/...`: Verify handlers can access trace IDs.

### Manual Verification
1. Start NATS and the example server.
2. Publish a message with `traceparent` header.
3. Observe logs for top-level `trace_id` and `span_id` presence.
