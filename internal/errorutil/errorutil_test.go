package errorutil

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Parallel()

	const expectedErrorMsg = "expected error to be '%s', got '%s'"

	tests := []struct {
		input    ErrorMsg
		expected string
	}{
		{
			input:    ErrorMsgParenNotClosedAtEOF,
			expected: ErrorMsgParenNotClosedAtEOF,
		},
	}

	for _, test := range tests {
		err := NewError(test.input)

		if err.Error() != test.expected {
			t.Errorf(expectedErrorMsg, test.expected, err.Error())
		}

		if errors.Unwrap(err).Error() != test.expected {
			t.Errorf(expectedErrorMsg, test.expected, errors.Unwrap(err).Error())
		}
	}
}

func TestNewErrorAt(t *testing.T) {
	t.Parallel()

	const expectedErrorMsg = "expected error to be '%s', got '%s'"

	tests := []struct {
		input    ErrorMsg
		pos      int
		expected string
	}{
		{
			input:    ErrorMsgParenNotClosedAtEOF,
			pos:      0,
			expected: ErrorMsgParenNotClosedAtEOF,
		},
	}

	for _, test := range tests {
		err := NewErrorAt(test.input, test.pos)

		if err.Error() != fmt.Sprintf("%s at position %d", test.expected, test.pos) {
			t.Errorf(expectedErrorMsg, test.expected, err.Error())
		}

		if errors.Unwrap(err).Error() != test.expected {
			t.Errorf(expectedErrorMsg, test.expected, errors.Unwrap(err).Error())
		}

		if err.Position() != test.pos {
			t.Errorf(
				"expected position to be %d, got %d",
				test.pos,
				err.Position(),
			)
		}
	}
}
