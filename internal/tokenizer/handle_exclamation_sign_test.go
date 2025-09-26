package tokenizer

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestHandleExclamationSign(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected *token.Token
	}{
		{
			name:  "no equal sign after",
			input: "! ",
			expected: token.NewToken(
				"!",
				token.TokenTypeNot,
				0,
				0,
			),
		},
		{
			name:  "equal sign after exclamation sign",
			input: "!=",
			expected: token.NewToken(
				"!=",
				token.TokenTypeNotEqual,
				0,
				0,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tokenizer := NewTokenizer(test.input)
			_, _ = tokenizer.GetNext()
			token, err := tokenizer.handleExclamationSign(1)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if token.Atom != test.expected.Atom {
				t.Fatalf("expected %s, got %s", test.expected.Atom, token.Atom)
			}

			if token.TokenType != test.expected.TokenType {
				t.Fatalf(
					"expected %T, got %T",
					test.expected.TokenType,
					token.TokenType,
				)
			}
		})
	}
}

func TestHandleExclamationSignErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "unexpected end of expression",
			input: "",
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewTokenizer(test.input).handleExclamationSign(1)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
