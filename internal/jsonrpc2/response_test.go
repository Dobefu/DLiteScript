package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		result   json.RawMessage
		id       RequestID
		expected string
	}{
		{
			name:     "test",
			result:   json.RawMessage("test"),
			id:       *NewRequestID(json.RawMessage("1")),
			expected: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			resp := NewResponse(test.result, test.id)

			if !bytes.Equal(resp.Result, test.result) {
				t.Errorf("expected \"%v\", got \"%v\"", test.result, resp.Result)
			}
		})
	}
}

func TestErrorResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		errData  []byte
		id       RequestID
		expected string
	}{
		{
			name:     "test",
			errData:  json.RawMessage("test"),
			id:       *NewRequestID(json.RawMessage("1")),
			expected: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			resp := NewErrorResponse(test.errData, test.id)

			if !bytes.Equal(resp.Error, test.errData) {
				t.Errorf("expected \"%v\", got \"%v\"", test.errData, resp.Error)
			}
		})
	}
}
