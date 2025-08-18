package lsp

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)

	go func() {
		errChan <- server.Start()
	}()

	select {
	case <-sigChan:
		shutdownChan := s.Handler.GetShutdownChan()

		if shutdownChan != nil {
			close(shutdownChan)
		}

		return nil

	case err := <-errChan:
		return err
	}
}
