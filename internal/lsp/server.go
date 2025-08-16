package lsp

import (
	"context"
)

// Server represents the LSP server.
type Server struct {
	stream interface{}
	conn   interface{}
}

// NewServer creates a new LSP server.
func NewServer() *Server {
	return &Server{}
}

// Start starts the LSP server.
func (s *Server) Start(ctx context.Context) error {
	return nil
}
