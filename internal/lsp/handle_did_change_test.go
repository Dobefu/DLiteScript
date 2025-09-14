package lsp

import (
	"encoding/json"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func TestHandleDidChange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		params   lsptypes.DidChangeParams
		expected string
	}{
		{
			name: "add content",
			params: lsptypes.DidChangeParams{
				TextDocument: lsptypes.TextDocumentItem{
					LanguageID: "dlitescript",
					URI:        "file:///test.dl",
					Version:    1,
					Text:       "",
				},
				ContentChanges: []lsptypes.ContentChange{
					{
						Text: "printf(\"test\")",
						Range: &lsptypes.Range{
							Start: lsptypes.Position{Line: 0, Character: 0},
							End:   lsptypes.Position{Line: 0, Character: 0},
						},
						RangeLength: 0,
					},
				},
			},
			expected: "printf(\"test\")",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			paramsJSON, err := json.Marshal(test.params)

			if err != nil {
				t.Errorf("expected no error, got \"%s\"", err.Error())
			}

			handler := NewHandler(false)
			handler.documents["file:///test.dl"] = lsptypes.Document{
				Text:        "",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{0},
			}

			response, jsonErr := handler.handleDidChange(json.RawMessage(paramsJSON))

			if jsonErr != nil {
				t.Errorf("expected no error, got \"%s\"", jsonErr.Error())
			}

			if response != nil {
				t.Errorf("expected nil response, got %s", string(response))
			}

			updatedDoc := handler.documents["file:///test.dl"]

			if updatedDoc.Text != test.expected {
				t.Errorf("expected document text to be %s, got %s", test.expected, updatedDoc.Text)
			}
		})
	}
}

func TestHandleDidChangeErr(t *testing.T) {
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
		{
			name:     "document not found",
			params:   json.RawMessage(`{"textDocument": {"uri": "file:///bogus.dl"}}`),
			expected: "code: -32602, message: Document not found",
		},
		{
			name: "start index greater than end index",
			params: json.RawMessage(`
			{
				"textDocument": {
					"uri": "file:///test.dl"
				},
				"contentChanges": [
					{
						"text": "printf(\"test\")",
						"range": {
							"start": {
								"line": 0,
								"character": 99
							},
							"end": {
								"line": 0,
								"character": 0
							}
						}
					}
				]
			}`),
			expected: "code: -32602, message: Start index is greater than end index",
		},
		{
			name: "end index out of bounds",
			params: json.RawMessage(`
			{
				"textDocument": {
					"uri": "file:///test.dl"
				},
				"contentChanges": [
					{
						"text": "printf(\"test\")",
						"range": {
							"start": {
								"line": 0,
								"character": 0
							},
							"end": {
								"line": 0,
								"character": 99
							}
						}
					}
				]
			}`),
			expected: "code: -32602, message: End index is out of bounds",
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

			_, jsonErr := handler.handleDidChange(test.params)

			if jsonErr == nil {
				t.Fatalf("expected error, got nil")
			}

			if jsonErr.Error() != test.expected {
				t.Errorf(
					"expected error message to be \"%s\", got \"%s\"",
					test.expected,
					jsonErr.Error(),
				)
			}
		})
	}
}
