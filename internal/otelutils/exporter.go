package otelutils

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"

	otellog "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/sdk/log"
)

// CompactExporter is an OpenTelemetry Log Exporter that writes log records
// as single-line (compact) JSON to an underlying io.Writer.
type CompactExporter struct {
	mu     sync.Mutex
	writer io.Writer
}

// NewCompactExporter creates a new CompactExporter.
// If writer is nil, os.Stdout is used.
func NewCompactExporter(w io.Writer) *CompactExporter {
	if w == nil {
		w = os.Stdout
	}
	return &CompactExporter{writer: w}
}

// Export serializes log records to JSON and writes them to the output.
func (e *CompactExporter) Export(ctx context.Context, records []log.Record) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	enc := json.NewEncoder(e.writer)
	for _, r := range records {
		m := map[string]interface{}{
			"time":  r.Timestamp().Format(time.RFC3339Nano),
			"level": r.SeverityText(),
			"msg":   r.Body().AsString(),
		}

		if tid := r.TraceID(); tid.IsValid() {
			m["trace_id"] = tid.String()
		}
		if sid := r.SpanID(); sid.IsValid() {
			m["span_id"] = sid.String()
		}

		// Handle attributes
		attrs := make(map[string]interface{})
		r.WalkAttributes(func(kv otellog.KeyValue) bool {
			attrs[kv.Key] = valueToAny(kv.Value)
			return true
		})
		if len(attrs) > 0 {
			m["attributes"] = attrs
		}

		if err := enc.Encode(m); err != nil {
			return err
		}
	}
	return nil
}

func valueToAny(v otellog.Value) interface{} {
	switch v.Kind() {
	case otellog.KindBool:
		return v.AsBool()
	case otellog.KindInt64:
		return v.AsInt64()
	case otellog.KindFloat64:
		return v.AsFloat64()
	case otellog.KindString:
		return v.AsString()
	case otellog.KindBytes:
		return v.AsBytes()
	case otellog.KindSlice:
		slice := v.AsSlice()
		res := make([]interface{}, len(slice))
		for i, val := range slice {
			res[i] = valueToAny(val)
		}
		return res
	case otellog.KindMap:
		kvs := v.AsMap()
		res := make(map[string]interface{}, len(kvs))
		for _, kv := range kvs {
			res[kv.Key] = valueToAny(kv.Value)
		}
		return res
	default:
		return nil
	}
}

// ForceFlush does nothing in this implementation.
func (e *CompactExporter) ForceFlush(ctx context.Context) error {
	return nil
}

// Shutdown closes the exporter.
func (e *CompactExporter) Shutdown(ctx context.Context) error {
	return nil
}
