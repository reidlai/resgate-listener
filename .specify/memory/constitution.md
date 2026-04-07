<!--
Sync Impact Report:
- Version change: 3.2.0 → 3.3.0
- Configured principles and domains: Added `Planning & Specifications` rule mandating user stories in implementation plans.
- Templates requiring updates (✅ updated):
  - `.specify/templates/plan-template.md`: ✅ Added `User Stories` mapping section to ensure requirements context is preserved during planning.
- Follow-up TODOs: None at this time.
-->
# Resgate Listener Constitution

## Purpose

This file defines immutable, non-negotiable engineering guardrails for this repository.
All specs, plans, and implementations MUST comply. Changes require the change-control
process below. This constitution is intentionally concise and enforceable.

## Project Identity

This is a go-res consumer project based on Golang and the RES protocol.

## Immutable Rules

### Branching, Protection, and Promotions

- No direct pushes to protected env branches (main, dev, sit, uat, staging, demo, prod).
  All changes flow through PRs from issue branches named
  `{issueNumber}-{issue_title_snake_case}`.
- Git commits MUST be pushed only to issue-generated branch `#####-XXXXX`, where #####
  is the GitHub issue number found in GitHub without zero padding and XXXXX is the issue
  title in snake_case with variable length not limited to 5 characters.
- Required flow:
  1. local work → push → issue branch
  2. PR issue → main (CI must pass)
  3. PR main → dev (CD to dev-test; run unit, mocked integration, smoke)
  4. On dev pass: create a release tag
  5. PR dev (release tag) → sit (CD to sit; run integration tests)
  6. On sit pass: tag sit at head
  7. With approval: PR sit → uat (CD; run UAT)
  8. On uat pass: tag uat at head
  9. With approval: PR uat → staging (CD; pre-prod checks)
  10. On staging pass: tag staging at head
  11. With approval: PR staging → prod (CD; prod deploy)
- main is PR-only aggregation; never deployed.


### Environments & Configuration

- Config is via environment variables (12-Factor); no secrets in repo. Use dotenv or
  platform env management for local only.
- Dev/prod parity: keep envs as similar as possible (12-Factor).

### Containers & Base Images

- MUST use distroless base docker image with multi-stage builds.
- Image MUST run as non-root, with minimal capabilities, and explicit HEALTHCHECK.

### Secrets & Sensitive Data

- Never commit secrets, keys, tokens, or certificates. All secrets come from the runtime
  secret store or env variables.
- CI MUST run secret detection; any finding blocks merge until remediated.

### Dependencies & Licensing

- Only permissive licenses (MIT, Apache-2.0, BSD). Critical CVEs (and container
  HIGH/CRITICAL, see rule 7) are zero-tolerance: remove, patch, or provide an approved,
  time-bound waiver before merge.

### DevSecOps Gates (Blocking)

**CI (on PR to main)**:

- SCA passes (no CRITICAL; HIGH requires approved waiver).
- Formatting/linting passes.
- Secrets scanning passes.
- Unit tests pass.
- SAST passes.

**Container build (pre-merge to env branch)**:

- Build with distroless base.
- Container scan with Trivy; CRITICAL/HIGH findings block unless approved waiver is
  attached to the PR and time-boxed.

**CD per env**:

- dev-test: DAST (e.g., ZAP) MUST run and attach a report; post-deploy integration tests with mock/stub, post-deploy smoke tests (via Cucumber) with mock/stub pass
  CRITICAL runtime vulns must page and trigger rollback policy (see runbook).
- SIT: DAST (e.g., ZAP) MUST run and attach a report; post-deploy integration tests, post-deploy smoke tests (via Cucumber) pass.
- UAT: DAST (e.g., ZAP) MUST run and attach a report; post-deploy Integration tests, post-deploy smoke tests (via Cucumber) pass.
- STAGING: DAST (e.g., ZAP) MUST run and attach a report; post-deploy integration tests, post-deploy smoke tests (via Cucumber) pass.
- PROD: DAST (e.g., ZAP) MUST run and attach a report

### Folder Structure

- Project MUST adhere to the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) conventions for repository structure.

### Planning & Specifications

- Feature plans (`plan.md`) MUST include fully-mapped User Stories (or Developer Stories for internal APIs) that detail how all product or technical requirements are delivered.

### Engineering Principles

- Adopt Twelve-Factor and SOLID principles as defaults for maintainability and scalability. Deviation requires an approved waiver.

## Decision Rights & Change Control

**File ownership**: `constitution.md` is protected. Changes require a PR approved by:
Security Lead + Product Owner (PO) + one Maintainer.

**Recording**: Each change MUST include a rationale and enforcement impact in the PR
description and be labeled `constitution-change`.

**Effective date**: Changes take effect only after CI policy jobs pass and merge.

**Version**: 3.3.0 | **Ratified**: 2026-04-05 | **Last Amended**: 2026-04-06
