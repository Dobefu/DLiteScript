package jsonrpc2

import "io"

// Stream represents the JSON-RPC stream.
type Stream struct {
	conn io.Closer
}

// NewServer creates a new JSON-RPC server.
func NewStream() *Stream {
	return &Stream{}
}
