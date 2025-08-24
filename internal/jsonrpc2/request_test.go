package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		method   string
		params   json.RawMessage
		expected string
	}{
		{
			name:     "test",
			method:   "test",
			params:   json.RawMessage("1"),
			expected: "1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			req := NewRequest(test.method, test.params)

			if req.Method != test.method {
				t.Errorf("expected \"%v\", got \"%v\"", test.method, req.Method)
			}

			if !bytes.Equal(req.Params, test.params) {
				t.Errorf("expected \"%v\", got \"%v\"", test.params, req.Params)
			}
		})
	}
}
