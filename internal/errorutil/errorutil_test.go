package errorutil

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestNewError(t *testing.T) {
	t.Parallel()

	const expectedErrorMsg = "expected error to be '%s', got '%s'"

	tests := []struct {
		name     string
		input    ErrorMsg
		expected string
	}{
		{
			name:  "paren not closed at eof",
			input: ErrorMsgParenNotClosedAtEOF,
			expected: fmt.Sprintf(
				"%s: %s",
				StageTokenize.String(),
				ErrorMsgParenNotClosedAtEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := NewError(StageTokenize, test.input)

			if err.Error() != test.expected {
				t.Errorf(expectedErrorMsg, test.expected, err.Error())
			}

			if errors.Unwrap(err).Error() != ErrorMsgParenNotClosedAtEOF {
				t.Errorf(
					expectedErrorMsg,
					ErrorMsgParenNotClosedAtEOF,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}

func TestNewErrorAt(t *testing.T) {
	t.Parallel()

	const expectedErrorMsg = "expected error to be '%s', got '%s'"

	tests := []struct {
		input    ErrorMsg
		pos      ast.Range
		expected string
	}{
		{
			input: ErrorMsgParenNotClosedAtEOF,
			pos: ast.Range{
				Start: ast.Position{Offset: 0, Line: 0, Column: 0},
				End:   ast.Position{Offset: 0, Line: 0, Column: 0},
			},
			expected: fmt.Sprintf(
				"%s: %s",
				StageTokenize.String(),
				ErrorMsgParenNotClosedAtEOF,
			),
		},
	}

	for _, test := range tests {
		err := NewErrorAt(StageTokenize, test.input, test.pos)

		if err == nil {
			t.Errorf("expected error, got nil")
		}

		expected := fmt.Sprintf(
			"%s %s",
			test.expected,
			test.pos.String(),
		)

		if err.Error() != expected {
			t.Errorf(expectedErrorMsg, expected, err.Error())
		}

		expected = ErrorMsgParenNotClosedAtEOF

		if errors.Unwrap(err).Error() != expected {
			t.Errorf(expectedErrorMsg, expected, errors.Unwrap(err).Error())
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
