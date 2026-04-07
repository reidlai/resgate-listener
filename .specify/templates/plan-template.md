# Implementation Plan: [FEATURE]

**Branch**: `[###-feature-name]` | **Date**: [DATE] | **Spec**: [link]
**Input**: Feature specification from `/specs/[###-feature-name]/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/plan-template.md` for the execution workflow.

## Summary

[Extract from feature spec: primary requirement + technical approach from research]

## User Stories

<!--
  ACTION REQUIRED: Copy the User Stories or Developer Stories from the spec and map
  out the high-level technical delivery approach for each.
-->

- **As a** [role], **I want to** [action] **so that** [benefit].
  - **Technical Delivery**: [Briefly describe how this will be implemented via UI/API/components]
  - **Acceptance**: [Testable criteria confirming delivery]

## Technical Context

<!--
  ACTION REQUIRED: Replace the content in this section with the technical details
  for the project. The structure here is presented in advisory capacity to guide
  the iteration process.
-->

**Language/Version**: Golang 
**Primary Dependencies**: RES protocol, go-res
**Storage**: [if applicable, e.g., PostgreSQL, CoreData, files or N/A]  
**Testing**: [e.g., testing, Ginkgo, or NEEDS CLARIFICATION]  
**Target Platform**: distroless Docker image
**Project Type**: go-res consumer web-service
**Performance Goals**: [domain-specific, e.g., 1000 req/s, 10k lines/sec, 60 fps or NEEDS CLARIFICATION]  
**Constraints**: [domain-specific, e.g., <200ms p95, <100MB memory, offline-capable or NEEDS CLARIFICATION]  
**Scale/Scope**: [domain-specific, e.g., 10k users, 1M LOC, 50 screens or NEEDS CLARIFICATION]

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [ ] **Containers**: Uses distroless base image with multi-stage builds.
- [ ] **Containers**: Runs as non-root, minimal capabilities, and includes explicit HEALTHCHECK.
- [ ] **Dependencies**: Uses permissive licenses (MIT, Apache-2.0, BSD) and has zero unaddressed critical CVEs.
- [ ] **DevSecOps**: PR pipeline includes SCA, linting, secrets scanning, SAST, Trivy, and unit tests.
- [ ] **DevSecOps**: CD pipelines per environment incorporate DAST, post-deploy integration, and Cucumber smoke tests.
- [ ] **Engineering**: Adheres to Twelve-Factor and SOLID principles, or includes an approved waiver for deviation.
- [ ] **Project Structure**: Follows the `golang-standards/project-layout` conventions.
- [ ] **Planning**: Ensures all requirements are explicitly delivered via mapped User/Developer Stories within this plan.

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)
<!--
  ACTION REQUIRED: Expand the Go project layout tree below with the concrete 
  packages/commands for this feature.
-->

```text
cmd/
└── [app-name]/         # Main applications
internal/
├── models/             # Private application and library code
├── services/           # (e.g. resgate handlers)
└── [feature]/          
pkg/
└── [library-name]/     # Library code ok to use by external applications
api/                    # OpenAPI/Swagger specs, JSON schema files, etc.
test/                   # Additional external test apps and test data
```

**Structure Decision**: [Document the specific packages, internal vs pkg decisions, and references created above]

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |
