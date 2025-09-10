package jsonrpc2

import (
	"encoding/json"
	"testing"
)

func TestRequestID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *RequestID
		expected string
	}{
		{
			name:     "test",
			input:    &RequestID{value: json.RawMessage("1")},
			expected: "1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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

func TestRequestIDErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *RequestID
		expected string
	}{
		{
			name:     "invalid JSON",
			input:    &RequestID{value: json.RawMessage("{")},
			expected: "invalid JSON in RequestID: '{'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.input.MarshalJSON()

			if err == nil {
				t.Errorf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
