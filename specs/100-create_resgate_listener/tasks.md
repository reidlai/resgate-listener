# Tasks: Resgate Listener Struct

**Input**: Design documents from `/specs/100-create_resgate_listener/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md

**Tests**: Test tasks have been mapped because TDD and high coverage is implicitly necessary for a core event router.
**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and package structure setup for the new library.

- [ ] T001 Initialize the go.mod file and grab dependencies (`nats-io/nats.go`, `jirenius/go-res`) in repo root.
- [ ] T002 [P] Create directory structure `pkg/listener` based on the golang-standards project layout.

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T003 Setup standard structured logging abstraction to track subscription mappings actively securely.

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Module Topic Router (Priority: P1) 🎯 MVP

**Goal**: As a module developer, I want to route specific NATS topics to my handlers so that I don't have to rewrite boilerplate subscription logic for every module.

**Independent Test**: Provide an initialized map containing topics adhering to the RES domain convention. Push NATS messages across exact and wildcard bounds and verify matched execution. Test dynamic panic and structural Close() methods explicitly leveraging the mock `nats-server` layer natively without leaks.

### Tests for User Story 1 ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T004 [P] [US1] Write test verifying successful routing of messages bound via RES domain strings and exact/wildcards in `pkg/listener/listener_test.go`
- [ ] T005 [P] [US1] Write test ensuring panicking handlers do NOT terminate the core NATS namespace pipeline stream in `pkg/listener/listener_test.go`

### Implementation for User Story 1

- [ ] T006 [US1] Define `MessageHandler` interface and `ResgateListener` containing strict memory segregation (`subs []string` vs `active map[string]*nats.Subscription`) in `pkg/listener/listener.go`
- [ ] T007 [US1] Implement `NewResgateListener()` constructor mapping dynamic DI handlers against allocated map capacities in `pkg/listener/listener.go`
- [ ] T008 [US1] Implement `Listen()` mapping execution block looping natively into raw `nc.Subscribe()`, pushing native subscriptions to the `active` internal storage map, and pushing tracking strings to the public `subs` array in `pkg/listener/listener.go`
- [ ] T009 [US1] Wrap `Listen()` invocation closure logic in a `defer recover()` abstraction to guard against panic crashes in `pkg/listener/listener.go`
- [ ] T010 [US1] Implement `Close()` routing iteration strictly against `active map[string]*nats.Subscription` safely bypassing string extraction mismatches in `pkg/listener/listener.go`

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [ ] T011 [P] Ensure SAST/DAST/Trivy rules are compliant with constitution in GitHub Actions or CI layer
- [ ] T012 Run quickstart.md validation locally to verify generic consumer API logic instantiation

---

## Dependencies & Execution Order

### Phase Dependencies
- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: Depend on Foundational phase completion
- **Polish (Final Phase)**: Depends on all desired user stories being complete

### User Story Dependencies
- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories

### Within Each User Story
- Tests MUST be written and FAIL before implementation
- Internal architectural segregation (T006) strictly precedes Construction & Methods securely.
- Methods scale incrementally: Constructor -> Listen -> Close

### Parallel Opportunities
- All Setup tasks marked [P] can run in parallel
- All tests for a user story marked [P] can run in parallel (T004, T005)

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently securely against NATS server instance.
5. Scale Polish integrations logically into the pipeline.
