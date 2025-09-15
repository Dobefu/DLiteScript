package lsp

import (
	"testing"
)

func TestServer(t *testing.T) {
	t.Parallel()

	handler := NewHandler(false)
	server := NewServer(handler)

	exitCode, err := server.Start()

	if err != nil {
		t.Fatalf("error starting server: %s", err.Error())
	}

	if exitCode == 0 {
		t.Fatalf("Expected non-zero exit code, got %d", exitCode)
	}
}
