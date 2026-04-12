# Data Model: OTel Log Attributes

This document defines the structured logging format and attributes used for OpenTelemetry correlation.

## LogRecord Attributes

| Field | Type | Description |
|-------|------|-------------|
| `time` | String (ISO8601) | Timestamp of the log event. |
| `level` | String | Log level (INFO, ERROR, WARN, DEBUG). |
| `msg` | String | The log message. |
| `trace_id` | String (Hex) | The 128-bit OpenTelemetry Trace ID. **(Top-level)** |
| `span_id` | String (Hex) | The 64-bit OpenTelemetry Span ID. **(Top-level)** |
| `target` | String | The Go package or component name. |

## Context Propagation

The `TraceID` and `SpanID` are extracted from the `context.Context` object passed to `slog` methods.

### Propagation Logic
1. **Extraction**: `otel.GetTextMapPropagator().Extract(ctx, carrier)`
2. **Correlation**: `slog.InfoContext(ctx, ...)`
