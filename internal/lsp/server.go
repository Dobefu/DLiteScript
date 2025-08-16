package lsp

import (
	"fmt"
	"os"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

// Server represents the LSP server.
type Server struct {
	Handler jsonrpc2.Handler
}

// NewServer creates a new LSP server.
func NewServer(handler jsonrpc2.Handler) *Server {
	return &Server{
		Handler: handler,
	}
}

// Start starts the LSP server.
func (s *Server) Start() error {
	server, err := jsonrpc2.NewServer(s.Handler, os.Stdin, os.Stdout)

	if err != nil {
		return fmt.Errorf("could not create JSON-RPC server: %w", err)
	}

	go func() {
		<-s.Handler.GetShutdownChan()

		os.Exit(0)
	}()

	return server.Start()
}
