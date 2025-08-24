package lsp

import (
	"testing"
)

func TestHandleExit(t *testing.T) {
	t.Parallel()

	handler := NewHandler(false)

	response, jsonErr := handler.handleExit()

	if jsonErr != nil {
		t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
	}

	if response != nil {
		t.Fatalf("expected nil response, got %s", string(response))
	}
}
