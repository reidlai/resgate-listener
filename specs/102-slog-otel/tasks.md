# Tasks: Adapt slog to OpenTelemetry

**Input**: Design documents from `/specs/102-slog-otel/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 [P] Install `go.opentelemetry.io/contrib/bridges/otelslog` and OTel dependencies via `go get`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure and interface updates required for OTel propagation

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T002 [P] Implement `NATSHeaderCarrier` utility in `internal/otelutils/propagation.go`
- [x] T015 [P] Implement `CompactExporter` in `internal/otelutils/exporter.go`
- [x] T003 Update `ResgateMessageHandler` interface to include `context.Context` in `pkg/resgate_message_handler/resgate_message_handler.go`

**Checkpoint**: Foundation ready - user story implementation can now begin

---

## Phase 3: User Story 1 - Observability Correlation (Priority: P1) 🎯 MVP

**Goal**: Enrich structured logs with TraceID and SpanID extracted from NATS messages.

**Independent Test**: Run example app, send NATS message with headers, verify JSON logs contain `trace_id` at top-level.

### Implementation for User Story 1

- [x] T004 [US1] Update `Listen()` loop to extract OTel context from NATS headers in `pkg/listener/listener.go`
- [x] T013 [US1] Bind `--otel-log-format` flag and `OTEL_LOG_FORMAT` env var in `example/cmd/start.go`
- [x] T017 [US1] Initialize OTel LoggerProvider with CompactExporter in `example/cmd/start.go`
- [x] T018 [US1] Update `slog` initialization logic to use `otelslog.NewHandler` in `example/cmd/start.go`
- [/] T005 [US1] Initialize OTel Trace SDK in `example/cmd/start.go`
- [x] T006 [P] [US1] Convert `fmt.Printf` to `slog.InfoContext` in `pkg/resgate_message_handler/resgate_message_handler.go`
- [x] T007 [P] [US1] Update listener logging to use `slog.InfoContext` in `pkg/listener/listener.go`
- [x] T008 [US1] Ensure `BaseResgateMessageHandler` implements the updated interface in `pkg/resgate_message_handler/resgate_message_handler.go`

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Final verification and documentation

- [x] T009 [P] Update `quickstart.md` with verification steps
- [x] T010 Run manual verification of JSON log output for `trace_id` and `span_id` (Verify SC-002: Readability confirmed)
- [x] T011 Verify performance impact is minimal via benchmark or manual observation (Verify SC-003: <5% overhead)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: Can start immediately
- **Foundational (Phase 2)**: Depends on Phase 1 completion
- **User Story 1 (Phase 3)**: Depends on Phase 2 completion
- **Polish (Final Phase)**: Depends on Phase 3 completion

### Parallel Opportunities

- T001 and T002 can start in parallel as they touch different areas (module info vs internal utils).
- T006 and T007 can be performed in parallel once the interface and listener logic are updated.

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Verify JSON logs have OTel identifiers.
