package lsp

import (
	"encoding/json"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func TestHandleDidClose(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params lsptypes.DidCloseParams
	}{
		{
			name: "close document",
			params: lsptypes.DidCloseParams{
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

			handler.documents["file:///test.dl"] = lsptypes.Document{
				Text:        "",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{0},
			}

			paramsJSON, err := json.Marshal(test.params)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			response, jsonErr := handler.handleDidClose(paramsJSON)

			if jsonErr != nil {
				t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
			}

			if response != nil {
				t.Fatalf("expected nil response, got %s", string(response))
			}

			_, hasDocument := handler.documents["file:///test.dl"]

			if hasDocument {
				t.Errorf("expected document to be removed, but it still exists")
			}
		})
	}
}

func TestHandleDidCloseErr(t *testing.T) {
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

			handler.documents["file:///test.dl"] = lsptypes.Document{
				Text:        "",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{0},
			}

			_, jsonErr := handler.handleDidClose(test.params)

			if jsonErr == nil {
				t.Fatalf("expected error, got nil")
			}

			_, hasDocument := handler.documents["file:///test.dl"]

			if !hasDocument {
				t.Errorf("expected document to still exist, but it has been removed")
			}
		})
	}
}
