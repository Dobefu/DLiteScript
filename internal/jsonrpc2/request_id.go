package jsonrpc2

import (
	"encoding/json"
)

// RequestID represents a JSON-RPC request ID.
type RequestID struct {
	value json.RawMessage
}

// NewRequestID creates a new JSON-RPC request ID.
func NewRequestID(value json.RawMessage) *RequestID {
	return &RequestID{
		value: value,
	}
}
