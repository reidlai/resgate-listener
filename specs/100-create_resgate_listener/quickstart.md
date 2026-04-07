# Quickstart: Resgate Listener

## Usage

Here's how to instantiate and use the `ResgateListener` in the main application:

```go
package main

import (
    "context"
    "log"
    
    "github.com/reidlai/resgate-listener/pkg/listener"
    "github.com/my-org/resgate-listener/internal/payment"
)

func main() {
    // 1. Establish NATS Connection
    nc := setupNATSConnection()
    
    // 2. Initialize Module Handlers
    paymentHandler := payment.NewPaymentModule()
    
    // 3. Define Topic Mapping
    topicMap := map[string]listener.MessageHandler{
        "payments.incoming": paymentHandler,
    }
    
    // 4. Instantiate Listener
    lsnr, err := listener.NewResgateListener(nc, topicMap)
    if err != nil {
        log.Fatalf("failed to create listener: %v", err)
    }
    
    // 5. Start Listening
    if err := lsnr.Listen(); err != nil {
        log.Fatalf("failed to start listener: %v", err)
    }
}
```
