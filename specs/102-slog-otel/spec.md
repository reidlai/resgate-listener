# Feature Specification: Adapt slog to OpenTelemetry

**Feature Branch**: `102-slog-otel`  
**Created**: 2026-04-12  
**Status**: Draft  
**Input**: User description: "need to adapt slog to opentelemetry https://uptrace.dev/guides/opentelemetry-slog"

## Clarifications

### Session 2026-04-12
- Q: Should OTel attributes be top-level or nested? → A: Top-level (directly in root JSON).
- Q: Global logger or Injected? → A: Global (Set as `slog.SetDefault()`).
- Q: Support multiple log formats? → A: Yes, via `--otel-log-format` flag (`text` vs `json`).

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Observability correlation (Priority: P1)

As a developer or site reliability engineer, I want our structured logs to automatically include OpenTelemetry trace and span identifiers so that I can correlate log messages with distributed traces when debugging issues.

**Why this priority**: Essential for modern observability and troubleshooting in microservice or event-driven environments (like NATS). It enables the "three pillars" (logs, metrics, and traces) to work together.

**Independent Test**: Can be tested by initiating a traced operation (e.g., a NATS message handler with a trace context), generating a log message using `slog`, and verifying that the output JSON contains `trace_id` and `span_id` fields that match the active trace.

**Acceptance Scenarios**:

1. **Given** an active OpenTelemetry trace, **When** a log is emitted using the `slog` logger with the active context, **Then** the log output MUST contain the correct `trace_id` and `span_id`.
2. **Given** no active trace is present in the context, **When** a log is emitted, **Then** the log output MUST NOT contain empty or invalid trace/span ID fields.
3. **Given** the system is configured for JSON logging, **When** logs are generated, **Then** they MUST remain valid JSON and include the OTel fields (`trace_id`, `span_id`) as top-level attributes.

---

### Edge Cases

- **Missing Context**: If a log is called without a context (e.g., `slog.Info(...)` instead of `slog.InfoContext(ctx, ...)`), how handles the system the missing trace data? (Default: Log without trace IDs, as slog's standard package-level functions don't take context).
- **Disabled Tracing**: If the trace exporter is disabled or not initialized, logs should still work gracefully without OTel attributes.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST implement/use a `slog.Handler` that extracts OpenTelemetry `TraceID` and `SpanID` from the `context.Context`.
- **FR-002**: The logger initialization MUST wrap the existing structured logger (JSON or text) with the OTel-aware handler and MUST set it as the global default using `slog.SetDefault()`.
- **FR-003**: Log field names for OTel identifiers MUST follow standard conventions (e.g., `trace_id`, `span_id`) and MUST be placed as top-level attributes in the JSON output.
- **FR-004**: The system MUST prioritize `slog.Handler` wrappers mentioned in the [uptrace/otel-slog](https://github.com/uptrace/otel-slog) guide.
- **FR-005**: All core components (`listener`, `handler`) MUST be updated to use `context`-aware logging methods (`slog.InfoContext`, etc.) to ensure trace context is propagated.
- **FR-006**: The system MUST support a `--otel-log-format` CLI flag (and `OTEL_LOG_FORMAT` env var) that toggles between `json` (default) and `text` output.
- **FR-007**: All log records MUST be emitted according to the official OpenTelemetry Log Data Model (OTLP) to ensure full compatibility with OTel collectors and standard processing pipelines.

### Key Entities *(include if feature involves data)*

- **Log Entry**: The structured representation of an event, now enriched with OTel metadata.
- **Trace Context**: The carrier of distributed tracing information (TraceID, SpanID) passed through the application.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 100% of logs generated within an active NATS message handler (if trace was propagated) must contain a valid `trace_id`.
- **SC-002**: Zero impact on log readability for systems not using OTel (if JSON field names are descriptive).
- **SC-003**: Log serialization overhead increases by less than 5% compared to the current structured logging implementation.

## Assumptions

- **Go Version**: Using Go 1.21+ (built-in `slog` support).
- **Log Format**: The user prefers JSON output for observability, but the system should support OTel attributes in text format if that is the current global setting.
- **Library Choice**: We will use the `uptrace/otel-slog` package as suggested by the user's provided link.
- **Global Logger**: The change will affect the global `slog` default logger or the logger instance injected into the main application components.
