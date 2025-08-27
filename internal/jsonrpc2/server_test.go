package jsonrpc2

import (
	"encoding/json"
	"io"
	"testing"
)

type mockHandler struct {
	shutdownChan chan struct{}
	shouldError  bool
	response     json.RawMessage
}

func (m *mockHandler) Handle(
	_ string,
	_ json.RawMessage,
) (json.RawMessage, *Error) {
	if m.shouldError {
		return nil, &Error{
			Code:    ErrorCodeInvalidRequest,
			Message: "some error",
			Data:    nil,
		}
	}

	return m.response, nil
}

func (m *mockHandler) GetShutdownChan() chan struct{} {
	return m.shutdownChan
}

func TestServer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		handler  Handler
		message  []byte
		expected error
	}{
		{
			name: "basic",
			handler: &mockHandler{
				shutdownChan: make(chan struct{}),
				shouldError:  false,
				response:     json.RawMessage(`"success"`),
			},
			message:  []byte(`{"jsonrpc": "2.0", "method": "test", "id": 1}`),
			expected: nil,
		},
		{
			name: "request without response",
			handler: &mockHandler{
				shutdownChan: make(chan struct{}),
				shouldError:  false,
				response:     nil,
			},
			message:  []byte(`{"jsonrpc": "2.0", "method": "test", "id": 1}`),
			expected: nil,
		},
		{
			name: "notification",
			handler: &mockHandler{
				shutdownChan: make(chan struct{}),
				shouldError:  false,
				response:     nil,
			},
			message:  []byte(`{"jsonrpc": "2.0", "method": "test"}`),
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			server, err := NewServer(test.handler, &io.SectionReader{}, io.Discard)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			err = server.Start()

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			err = server.handleMessage(test.message)

			if test.expected == nil && err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
		})
	}
}

func TestServerErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		handler  Handler
		message  []byte
		expected string
	}{
		{
			name:     "nil handler",
			handler:  nil,
			message:  nil,
			expected: "nil handler provided",
		},
		{
			name: "invalid json",
			handler: &mockHandler{
				shutdownChan: make(chan struct{}),
				shouldError:  false,
				response:     nil,
			},
			message:  []byte(`{`),
			expected: "could not unmarshal message: unexpected end of JSON input",
		},
		{
			name: "handler error",
			handler: &mockHandler{
				shutdownChan: make(chan struct{}),
				shouldError:  true,
				response:     nil,
			},
			message:  []byte(`{"jsonrpc": "2.0", "method": "test", "id": 1}`),
			expected: "could not handle message: some error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			server, err := NewServer(test.handler, &io.SectionReader{}, io.Discard)

			if test.handler == nil {
				if err == nil {
					t.Fatalf("expected error, got none")
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			err = server.handleMessage(test.message)

			if err.Error() != test.expected {
				t.Errorf(
					"expected error to be \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
