# Quickstart: OpenTelemetry Logging (Standard Compliant)

This guide explains how to enable and verify full OpenTelemetry Standard (OTLP) correlation in your logs while maintaining readability.

## Setup

The OpenTelemetry logging integration is fully standard-compliant using the official `otelslog` bridge.

### CLI Flags & Environment Variables

| Flag | Env Var | Default | Description |
|------|---------|---------|-------------|
| `--resgate.nats-url` | `RESGATE_NATS_URL` | `nats://localhost:4222` | NATS Server URL |
| `--log-format` | `LOG_FORMAT` | `json` | Log format: `json` or `text` |

## Verification

### 1. Start the stack
```bash
docker compose -f deploy/docker-compose/docker-compose.localdev.yaml up -d
```

### 2. Run the example server
```bash
go run example/main.go start
```

### 3. Verify OTel Field Schema
The logs now follow the OTel Log Data Model but are formatted as single-line JSON for readability:

**Example Output:**
```json
{
  "time": "2026-04-12T05:00:00.000000Z",
  "level": "INFO",
  "msg": "Starting dummy server...",
  "trace_id": "...",
  "span_id": "...",
  "attributes": {
    "resgate.nats-url": "nats://localhost:4222",
    "server.log-format": "json"
  }
}
```

### 4. Direct Correlation Test
To see active trace IDs, publish a message with a `traceparent` header:

```bash
nats pub "get.example.1" --header "traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01" "hello"
```

Verify that the `trace_id` and `span_id` fields are populated and match the header.

## Developer Usage

Always use `context`-aware logging methods:

```go
slog.InfoContext(ctx, "Operation successful", "detail", "some value")
```
> [!NOTE]
> Attributes provided as key-value pairs in `slog` calls are automatically nested under the `attributes` key in the OTel standard output.
