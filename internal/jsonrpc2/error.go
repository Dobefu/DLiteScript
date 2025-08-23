package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

// ErrorCode represents a JSON-RPC error code.
// For more information, see the [specification].
//
// [specification]: https://www.jsonrpc.org/specification#error_object
type ErrorCode int

const (
	// ErrorCodeParseError is the error code for parsing errors.
	ErrorCodeParseError = -32700
	// ErrorCodeInvalidRequest is the error code for invalid request errors.
	ErrorCodeInvalidRequest = -32600
	// ErrorCodeMethodNotFound is the error code for method not found errors.
	ErrorCodeMethodNotFound = -32601
	// ErrorCodeInvalidParams is the error code for invalid params errors.
	ErrorCodeInvalidParams = -32602
	// ErrorCodeInternalError is the error code for internal errors.
	ErrorCodeInternalError = -32603
	// ErrorCodeServerError is the error code for server errors.
	ErrorCodeServerError = -32000
)

// Error represents a JSON-RPC error.
// For more information, see the [specification].
//
// [specification]: https://www.jsonrpc.org/specification#error_object
type Error struct {
	Code    ErrorCode        `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data,omitempty"`
}

// NewError creates a new JSON-RPC error.
func NewError(code ErrorCode, message string, data *json.RawMessage) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// Error returns the error message.
func (e *Error) Error() string {
	if e.Data == nil {
		return fmt.Sprintf(
			"code: %d, message: %s",
			e.Code,
			e.Message,
		)
	}

	return fmt.Sprintf(
		"code: %d, message: %s, data: %s",
		e.Code,
		e.Message,
		string(*e.Data),
	)
}
