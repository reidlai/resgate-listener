# Research: Resgate Listener

## NATS Subscription teardown patterns
- Decision: Internal mapping wrapper.
- Rationale: While standard visibility requires strings (`[]string`), Native NATS logic relies on memory pointers to the actual Subscription object `*nats.Subscription` to call `Unsubscribe()`. Keeping them in an unexported mapped structure solves both constraints transparently without Memory leaks or panic faults.
- Alternatives: Not explicitly calling `Close()` and letting the system reap them upon shutdown was rejected as it makes testing and modular graceful reloads impossible.

## RES Protocol Format 
- Decision: Treated as Domain conventions rather than strictly enforced validations.
- Rationale: Rejecting mismatched topic arrays dynamically checking `get.<topic>` strings against regex creates massive runtime penalty loops and complex error bounds. By treating it as a Domain convention, developers hold the responsibility for RES compliance while gaining highly performant raw memory mapping.
- Alternatives: Running strict string splitting and error returns on `Listen()`. Rejected.
