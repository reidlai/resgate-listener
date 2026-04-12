package otelutils

import "github.com/nats-io/nats.go"

// NATSHeaderCarrier adapts nats.Header to satisfy the propagation.TextMapCarrier interface.
type NATSHeaderCarrier nats.Header

// Get returns the value associated with the passed key.
func (c NATSHeaderCarrier) Get(key string) string {
	return nats.Header(c).Get(key)
}

// Set stores the key-value pair.
func (c NATSHeaderCarrier) Set(key string, value string) {
	nats.Header(c).Set(key, value)
}

// Keys lists the keys stored in this carrier.
func (c NATSHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}
