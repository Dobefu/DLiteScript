package lsp

import (
	"encoding/json"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func TestHandleDidOpen(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params lsptypes.DidOpenParams
	}{
		{
			name: "open document",
			params: lsptypes.DidOpenParams{
				TextDocument: lsptypes.TextDocumentItem{
					LanguageID: "dlitescript",
					URI:        "file:///test.dl",
					Version:    1,
					Text:       "",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(false)

			paramsJSON, err := json.Marshal(test.params)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			response, jsonErr := handler.handleDidOpen(paramsJSON)

			if jsonErr != nil {
				t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
			}

			if response != nil {
				t.Fatalf("expected nil response, got %s", string(response))
			}

			_, hasDocument := handler.documents["file:///test.dl"]

			if !hasDocument {
				t.Errorf("expected document to be added, but it doesn't exist")
			}
		})
	}
}

func TestHandleDidOpenErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		params   json.RawMessage
		expected string
	}{
		{
			name:     "unmarshal params",
			params:   json.RawMessage(`{`),
			expected: "code: -32602, message: unexpected end of JSON input",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(false)

			_, jsonErr := handler.handleDidOpen(test.params)

			if jsonErr == nil {
				t.Fatalf("expected error, got nil")
			}

			_, hasDocument := handler.documents["file:///test.dl"]

			if hasDocument {
				t.Errorf("expected document to not exist, but it does")
			}
		})
	}
}
