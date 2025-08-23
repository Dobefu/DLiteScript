package jsonrpc2

import (
	"strings"
	"testing"
)

func TestStream(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid request",
			input:    "Content-Length: 45\r\n\r\n{\"jsonrpc\": \"2.0\", \"method\": \"test\", \"id\": 1}",
			expected: `{"jsonrpc": "2.0", "method": "test", "id": 1}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			stream := NewStream(strings.NewReader(test.input), nil)

			msg, err := stream.ReadMessage()

			if err != nil {
				t.Fatalf("expected no error, got \"%v\"", err)
			}

			if string(msg) != test.expected {
				t.Errorf("expected \"%v\", got \"%v\"", test.expected, string(msg))
			}
		})
	}
}

func TestStreamErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "bogus content length",
			input:    "Content-Length: bogus\r\n\r\n{\"jsonrpc\": \"2.0\", \"method\": \"test\", \"id\": 1}",
			expected: "could not parse content length",
		},
		{
			name:     "negative content length",
			input:    "Content-Length: -1\r\n\r\n{\"jsonrpc\": \"2.0\", \"method\": \"test\", \"id\": 1}",
			expected: "invalid content length",
		},
		{
			name:     "missing content length",
			input:    "\r\n\r\n{\"jsonrpc\": \"2.0\", \"method\": \"test\", \"id\": 1}",
			expected: "no Content-Length header found",
		},
		{
			name:     "missing content",
			input:    "Content-Length: 1\r\n\r\n",
			expected: "could not read the message body",
		},
		{
			name:     "read error",
			input:    "Content-Length: 3",
			expected: "could not read the message: EOF",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			stream := NewStream(strings.NewReader(test.input), nil)

			_, err := stream.ReadMessage()

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if !strings.Contains(err.Error(), test.expected) {
				t.Errorf(
					"expected error to contain \"%v\", got \"%v\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
