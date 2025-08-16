package lsp

import (
	"context"
	"os"

	"go.lsp.dev/jsonrpc2"
)

// Server represents the LSP server.
type Server struct {
	stream jsonrpc2.Stream
	conn   jsonrpc2.Conn
}

// NewServer creates a new LSP server.
func NewServer() *Server {
	stream := jsonrpc2.NewStream(os.Stdin)

	return &Server{
		stream: stream,
		conn:   jsonrpc2.NewConn(stream),
	}
}

// Start starts the LSP server.
func (s *Server) Start(_ context.Context) error {
	return nil
}
