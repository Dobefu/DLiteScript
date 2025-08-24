package lsp

import (
	"testing"
)

func TestHandleInitialize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "initialize",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(false)

			response, jsonErr := handler.handleInitialize()

			if jsonErr != nil {
				t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
			}

			if response == nil {
				t.Fatalf("expected response, got nil")
			}
		})
	}
}
