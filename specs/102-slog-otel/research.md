# Research: slog to OpenTelemetry Integration

## Problem Statement

The `resgate-listener` project needs to integrate `slog` with OpenTelemetry to provide distributed tracing correlation in its structured logs. Currently, logs lack `trace_id` and `span_id` fields, making it difficult to trace requests across the asynchronous NATS messaging boundary.

## Findings

### 1. Log Correlation with slog

To inject tracing identifiers into `slog` output, the most effective approach is using a custom `slog.Handler` that extracts the `SpanContext` from the `context.Context` during the `Handle` call.

- **Option A: `otelslog` Bridge**: The official OTel Go bridge (`go.opentelemetry.io/contrib/bridges/otelslog`) is the standard way to integrate `slog` with OpenTelemetry. It maps `slog.Record` fields to OpenTelemetry log records and handles context-aware logging to ensure trace and span IDs are correctly captured.
- **Option B: Manual Implementation**: One could implement a custom `slog.Handler`, but the official bridge is preferred for compliance with OTel standards.

**Decision**: Use `go.opentelemetry.io/contrib/bridges/otelslog` as the primary handler bridge.

### 2. Context Propagation over NATS

NATS supports headers since v2.2.0. OpenTelemetry provides `TextMapPropagator` to inject/extract context into/from these headers.

- **Mechanism**: Implement the `propagation.TextMapCarrier` interface for `nats.Header`.
- **Implementation**:
  - `Inject`: When publishing (out of scope for now, but good to know).
  - `Extract`: When receiving a message in the `listener`.

**Decision**: Implement a `NATSHeaderCarrier` in a internal utility package to handle context extraction in the listener.

### 3. Interface Changes

To enable OTel, `slog` calls must use `slog.InfoContext(ctx, ...)` instead of `slog.Info(...)`. This requires propagating `context.Context` from the NATS subscription level down to the `ResgateMessageHandler`.

**Decision**: Update the `ResgateMessageHandler` interface's `HandleMessage` method to include a `ctx context.Context` parameter.

## Implementation Details

### Dependencies
- `go.opentelemetry.io/otel`
- `go.opentelemetry.io/otel/trace`
- `go.opentelemetry.io/contrib/bridges/otelslog`

### Field Names
Standard practitioners use `trace_id` and `span_id`. The OTel logging model uses these by default.

### Constraints
- Must handle missing trace context gracefully (FR-001/T002).
- Must not leak OTel implementation into the public API more than necessary (internal utilities preferred).
