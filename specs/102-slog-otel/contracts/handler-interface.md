# Contract: ResgateMessageHandler Interface

This contract defines the public interface for NATS message handlers in the `resgate-listener` library.

## Interface Definition

```go
type ResgateMessageHandler interface {
    // HandleMessage processes an incoming NATS message.
    // The ctx parameter provides the OpenTelemetry trace context extracted from NATS headers.
    HandleMessage(ctx context.Context, topic string, msg *nats.Msg)

    // Resource handlers (Access, Get, Add, etc.)
    AccessResourceHandler() (interface{}, error)
    GetResourcesHandler() (interface{}, error)
    GetResourceHandler() (interface{}, error)
    AddResourceHandler() (interface{}, error)
    ChangeResourceHandler() (interface{}, error)
    RemoveResourceHandler() (interface{}, error)
}
```

## Context Propagation Contract
1. **Source**: The `ResgateListener` extracts trace information from the `nats.Msg.Header`.
2. **Behavior**: If a `traceparent` header is present, `HandleMessage` receives a context containing that span context. If not, a background context (or a new root span) is provided.
3. **Usage**: Implementers MUST use `ctx` for all logging and downstream service calls to maintain trace continuity.
