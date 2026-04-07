# Feature Specification: Resgate Listener Struct

## Background
The application acts as a consumer for RES protocol messages via NATS. We need a modular listener that can map incoming NATS topics to specific handler modules.

## Requirements
1. **Location**: Create a `resgate listener` struct in the `pkg/` folder.
2. **Construction**: The listener should be constructed with a mapping of NATS topics to their corresponding module structs (using dependency injection).
3. **Listen Method**: The listener must provide a `Listen()` method.
4. **Behavior**: The `Listen()` method should consume all incoming payments broadcast over the topics defined in the listener struct's mapping, dispatching them to the injected module struct handlers.

## Acceptance Criteria
- A developer can instantiate the listener by passing a valid initialized NATS connection.
- Topic mappings MUST support NATS wildcards (`*` and `>`) to correctly route broad streams.
- Topics SHOULD adhere to standard RES specifications (`get.<resource>`, `event.<resource>.<action>`, `call.<resource>.<method>`) as namespace conventions natively without requiring manual Regex rejection at runtime.
- A developer can bind specific topic strings to their module handlers via dependency injection.
- When a message arrives on a bound topic, the exact configured handler is invoked.
- If an unmapped topic receives a message (e.g. from wildcard subscriptions), it does not crash the listener.
- The listener MUST recover safely from any handler panics/errors without dropping the core NATS subscription.

## User Stories
- **As a** module developer, **I want to** route specific NATS topics to my handlers **so that** I don't have to rewrite boilerplate subscription logic for every module.

## Clarifications

### Session 2026-04-06
- Q: How should we capture integration requirements given the lack of traditional User Stories? → A: Add formal Acceptance Criteria bullet points under the requirements.
- Q: Does the topic mapping need to support NATS wildcards? → A: Wildcard support (* and >) is required logic.
- Q: What happens if a handler panics/fails? → A: Safely recover and continue processing.
- Q: How should we format User Stories to comply with Constitution v3.3.0? → A: Feature the Module Developer persona avoiding subscription boilerplate.
- Q: How should we manage NATS subscription teardown if `subs` is exposed as an array of strings? → A: Maintain an internal map of `*nats.Subscription` to handle safe termination organically while exposing strings for tracking.
- Q: How strictly should the listener enforce RES protocol domains? → A: Treat RES protocol (`get`, `call`, `event`) loosely as domain conventions, trusting mapping rather than enforcing runtime panic validation.
