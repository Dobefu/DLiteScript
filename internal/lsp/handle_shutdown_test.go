package lsp

import (
	"testing"
)

func TestHandleShutdown(t *testing.T) {
	t.Parallel()

	handler := NewHandler(false)

	response, jsonErr := handler.handleShutdown()

	if jsonErr != nil {
		t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
	}

	if response != nil {
		t.Fatalf("expected nil response, got %s", string(response))
	}
}
