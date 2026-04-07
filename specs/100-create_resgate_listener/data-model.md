# Data Model: Resgate Listener

This is a library abstraction over NATS and `go-res`, holding only in-memory states instead of persistent data entities.

## Entities

### `ResgateListener` (In-Memory Struct)
- **Role**: Primary routing manager for NATS bindings.
- **Fields**:
  - `nc` (`*nats.Conn`): Embedded raw stream connection pointing back to nats-server.
  - `handlers` (`map[string]MessageHandler`): The bound functions executing dynamically.
  - `subs` (`[]string`): Public-facing tracking array for active listener identifiers.
  - `active` (`map[string]*nats.Subscription`): Unexported memory mapping maintaining direct un-subscribe links for tear down.

### `MessageHandler` (Interface Contract)
- **Role**: The consumer standard format defining what an injected Module must implement to safely be called dynamically by the listener router.
- **Fields**:
  - `HandleMessage(msg *nats.Msg)`: Processing loop.

## Integration Constraints
- **Format Schema**: Mapped Topic strings SHOULD align cleanly inside the RES bounds:
  - `get.<resource>`
  - `call.<resource>.<method>`
  - `event.<resource>.<action>`
