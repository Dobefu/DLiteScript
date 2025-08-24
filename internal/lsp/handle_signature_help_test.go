package lsp

import (
	"encoding/json"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

func TestHandleSignatureHelp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params json.RawMessage
	}{
		{
			name: "signature help",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///test.dl"
				},
				"position": {
					"line": 0,
					"character": 0
				}
			}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(false)

			response, jsonErr := handler.handleSignatureHelp(test.params)

			if jsonErr != nil {
				t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
			}

			if response == nil {
				t.Fatalf("expected response, got nil")
			}
		})
	}
}

func TestHandleSignatureHelpErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params json.RawMessage
	}{
		{
			name:   "invalid json",
			params: json.RawMessage(`{`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(false)

			_, jsonErr := handler.handleSignatureHelp(test.params)

			if jsonErr == nil {
				t.Fatalf("expected error, got nil")
			}

			if jsonErr.Code != jsonrpc2.ErrorCodeInvalidParams {
				t.Fatalf("expected error code %d, got %d", jsonrpc2.ErrorCodeInvalidParams, jsonErr.Code)
			}
		})
	}
}
