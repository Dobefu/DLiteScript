package jsonrpc2

import "encoding/json"

// Response represents a JSON-RPC response.
// For more information, see the [specification].
//
// [specification]: https://www.jsonrpc.org/specification#response_object
type Response struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   json.RawMessage `json:"error,omitempty"`
	ID      RequestID       `json:"id"`
}

// NewResponse creates a new JSON-RPC response.
func NewResponse(result json.RawMessage, id RequestID) *Response {
	return &Response{
		JSONRPC: "2.0",
		Result:  result,
		Error:   nil,
		ID:      id,
	}
}

// NewErrorResponse creates a new JSON-RPC error response.
func NewErrorResponse(errData []byte, id RequestID) *Response {
	return &Response{
		JSONRPC: "2.0",
		Result:  nil,
		Error:   errData,
		ID:      id,
	}
}
