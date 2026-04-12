# resgate-listener Development Guidelines

Auto-generated from all feature plans. Last updated: 2026-04-12

## Active Technologies
- Golang 1.25+ + `nats-io/nats.go`, `jirenius/go-res` (100-create_resgate_listener)
- N/A (In-memory router layer) (100-create_resgate_listener)
- Golang 1.22+ + NATS 2.10, Resgate 1.7 (101-docker-localdev-setup)
- N/A (In-memory JetStream) (101-docker-localdev-setup)
- Golang 1.22+ + `github.com/uptrace/otel-slog`, `go.opentelemetry.io/otel`, `go.opentelemetry.io/otel/trace` (102-slog-otel)

- Golang + `github.com/nats-io/nats.go`, `github.com/jirenius/go-res` (If interacting natively with res protocol models) (100-create_resgate_listener)

## Project Structure

```text
backend/
frontend/
tests/
```

## Commands

# Add commands for Golang

## Code Style

Golang: Follow standard conventions

## Recent Changes
- 102-slog-otel: Added Golang 1.22+ + `github.com/uptrace/otel-slog`, `go.opentelemetry.io/otel`, `go.opentelemetry.io/otel/trace`
- 102-slog-otel: Added Golang 1.22+ + `github.com/uptrace/otel-slog`, `go.opentelemetry.io/otel`, `go.opentelemetry.io/otel/trace`
- 101-docker-localdev-setup: Added Golang 1.22+ + NATS 2.10, Resgate 1.7


<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
