package jsonrpc2

import (
	"encoding/json"
	"testing"
)

func TestRequestID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *RequestID
		expected string
	}{
		{
			input:    &RequestID{value: json.RawMessage("1")},
			expected: "1",
		},
	}

	for _, test := range tests {
		t.Run(test.input.String(), func(t *testing.T) {
			t.Parallel()

			if test.input.String() != test.expected {
				t.Errorf("expected \"%v\", got \"%v\"", test.expected, test.input.String())
			}

			if test.input.IsNull() {
				t.Errorf("expected \"%v\", got \"%v\"", false, test.input.IsNull())
			}

			jsonBytes, err := test.input.MarshalJSON()

			if err != nil {
				t.Errorf("expected no error, got \"%v\"", err)
			}

			err = test.input.UnmarshalJSON(jsonBytes)

			if err != nil {
				t.Errorf("expected no error, got \"%v\"", err)
			}
		})
	}
}
