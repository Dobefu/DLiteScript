package lsp

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func TestHandleHover(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		params   lsptypes.HoverParams
		expected lsptypes.Hover
	}{
		{
			name: "hover document",
			params: lsptypes.HoverParams{
				TextDocument: lsptypes.TextDocument{
					URI: "file:///test.dl",
				},
				Position: lsptypes.Position{
					Line:      0,
					Character: 0,
				},
			},
			expected: lsptypes.Hover{
				Contents: "\n\n```dlitescript\nfunc printf(format string, ...args any)\n```\n\nPrints a formatted string.\n\n**Parameters:**\n```dlitescript\nformat string\n```\n```dlitescript\n...args any\n```\n",
				Range: &lsptypes.Range{
					Start: lsptypes.Position{
						Line:      0,
						Character: 0,
					},
					End: lsptypes.Position{
						Line:      0,
						Character: 0,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(false)

			handler.documents["file:///test.dl"] = lsptypes.Document{
				Text:        "printf(\"test\")",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{0},
			}

			paramsJSON, err := json.Marshal(test.params)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			response, jsonErr := handler.handleHover(paramsJSON)

			if jsonErr != nil {
				t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
			}

			expectedJSON, err := json.Marshal(test.expected)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if string(response) != string(expectedJSON) {
				t.Fatalf("expected %s, got %s", string(expectedJSON), string(response))
			}
		})
	}
}

func TestHandleHoverErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		params   json.RawMessage
		expected string
	}{
		{
			name:     "unmarshal params",
			params:   json.RawMessage(`{`),
			expected: "unexpected end of JSON input",
		},
		{
			name: "document not found",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///bogus.dl"
				},
				"position": {
					"line": 0,
					"character": 0
				}
			}`),
			expected: "Document not found",
		},
		{
			name: "invalid document",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///invalid.dl"
				},
				"position": {
					"line": 0,
					"character": 0
				}
			}`),
			expected: fmt.Sprintf(
				"%s: %s at position 12",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid position",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///test.dl"
				},
				"position": {
					"line": 0,
					"character": 99
				}
			}`),
			expected: "Could not find AST node at position",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(false)

			handler.documents["file:///test.dl"] = lsptypes.Document{
				Text:        "printf(\"test\")",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{0},
			}

			handler.documents["file:///invalid.dl"] = lsptypes.Document{
				Text:        "printf(\"test",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{0},
			}

			_, jsonErr := handler.handleHover(test.params)

			if jsonErr == nil {
				t.Fatalf("expected error, got nil")
			}

			if jsonErr.Message != test.expected {
				t.Fatalf("expected error message \"%s\", got \"%s\"", test.expected, jsonErr.Message)
			}

			if jsonErr.Code != jsonrpc2.ErrorCodeInvalidParams {
				t.Fatalf("expected error code %d, got %d", jsonrpc2.ErrorCodeInvalidParams, jsonErr.Code)
			}
		})
	}
}
