package lsp

import (
	"encoding/json"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		method         string
		params         json.RawMessage
		expectResponse bool
	}{
		{
			name:           "initialize",
			method:         "initialize",
			params:         json.RawMessage(""),
			expectResponse: true,
		},
		{
			name:           "initialized",
			method:         "initialized",
			params:         json.RawMessage(""),
			expectResponse: false,
		},
		{
			name:   "didOpen",
			method: "textDocument/didOpen",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///test.dl",
					"languageId": "dlite",
					"version": 1,
					"text": "print(\"test\")"
				}
			}`),
			expectResponse: false,
		},
		{
			name:   "setTrace",
			method: "$/setTrace",
			params: json.RawMessage(`{
				"value": "messages"
			}`),
			expectResponse: false,
		},
		{
			name:   "didChange",
			method: "textDocument/didChange",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///test.dl",
					"languageId": "dlite",
					"version": 1,
					"text": "print(\"test\")"
				}
			}`),
			expectResponse: false,
		},
		{
			name:   "didClose",
			method: "textDocument/didClose",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///test.dl"
				}
			}`),
			expectResponse: false,
		},
		{
			name:   "hover",
			method: "textDocument/hover",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///test.dl"
				},
				"position": {
					"line": 0,
					"character": 0
				}
			}`),
			expectResponse: true,
		},
		{
			name:   "signatureHelp",
			method: "textDocument/signatureHelp",
			params: json.RawMessage(`{
				"textDocument": {
					"uri": "file:///test.dl"
				}
			}`),
			expectResponse: true,
		},
		{
			name:           "shutdown",
			method:         "shutdown",
			params:         json.RawMessage(""),
			expectResponse: false,
		},
		{
			name:           "exit",
			method:         "exit",
			params:         json.RawMessage(""),
			expectResponse: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(true)

			handler.documents["file:///test.dl"] = lsptypes.Document{
				Text:        "printf(\"test\")",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{13},
			}

			shutdownChan := handler.GetShutdownChan()

			if shutdownChan == nil {
				t.Fatalf("expected shutdown chan, got nil")
			}

			response, jsonErr := handler.Handle(test.method, test.params)

			if jsonErr != nil {
				t.Fatalf("expected no error, got \"%s\"", jsonErr.Error())
			}

			if test.expectResponse && response == nil {
				t.Fatalf("expected response for %s, got nil", test.method)
			} else if !test.expectResponse && response != nil {
				t.Fatalf("expected no response for %s, got %s", test.method, string(response))
			}
		})
	}
}

func TestHandlerErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		method string
		params json.RawMessage
	}{
		{
			name:   "bogus method",
			method: "bogus",
			params: json.RawMessage(""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(true)
			_, jsonErr := handler.Handle(test.method, test.params)

			if jsonErr == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
