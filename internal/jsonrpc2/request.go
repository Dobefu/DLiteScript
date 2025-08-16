package jsonrpc2

import "encoding/json"

// Request represents a JSON-RPC request.
// For more information, see the [Specification].
//
// [Specification]: https://www.jsonrpc.org/specification#request_object
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      RequestID       `json:"id"`
}

// NewRequest creates a new JSON-RPC request.
func NewRequest(method string, params json.RawMessage) *Request {
	return &Request{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      *NewRequestID(nil),
	}
}
