package jsonrpc2

import (
	"encoding/json"
	"fmt"
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

// MarshalJSON marshals the RequestID to JSON.
func (r *RequestID) MarshalJSON() ([]byte, error) {
	if len(r.value) == 0 {
		return []byte("null"), nil
	}

	if !json.Valid(r.value) {
		return nil, fmt.Errorf("invalid JSON in RequestID: %s", string(r.value))
	}

	return r.value, nil
}

// UnmarshalJSON unmarshals JSON into the RequestID.
func (r *RequestID) UnmarshalJSON(data []byte) error {
	r.value = json.RawMessage(data)

	return nil
}

// IsNull returns true if the RequestID is null.
func (r *RequestID) IsNull() bool {
	return string(r.value) == "null"
}

// String returns the string representation of the RequestID.
func (r *RequestID) String() string {
	return string(r.value)
}
