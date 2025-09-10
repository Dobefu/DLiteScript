package jsonrpc2

import (
	"encoding/json"
	"testing"
)

func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		code     ErrorCode
		message  string
		data     *json.RawMessage
		expected string
	}{
		{
			name:     "test",
			code:     -32000,
			message:  "test",
			data:     &json.RawMessage{},
			expected: "code: -32000, message: test, data: ",
		},
		{
			name:     "no data",
			code:     -32000,
			message:  "test",
			data:     nil,
			expected: "code: -32000, message: test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			testErr := NewError(
				test.code,
				test.message,
				test.data,
			)

			if testErr.Error() != test.expected {
				t.Errorf("expected \"%v\", got \"%v\"", test.expected, testErr.Error())
			}
		})
	}
}
