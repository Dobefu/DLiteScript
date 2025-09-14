package lsp

import (
	"testing"
)

func TestServer(t *testing.T) {
	t.Parallel()

	handler := NewHandler(false)
	server := NewServer(handler)

	err := server.Start()

	if err != nil {
		t.Errorf("error starting server: %v", err)
	}
}
