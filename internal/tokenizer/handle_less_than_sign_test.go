package tokenizer

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestHandleLessThanSign(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected *token.Token
	}{
		{
			name:  "no equal sign after",
			input: "< ",
			expected: token.NewToken(
				"<",
				token.TokenTypeLessThan,
				0,
				0,
			),
		},
		{
			name:  "equal sign after less than sign",
			input: "<=",
			expected: token.NewToken(
				"<=",
				token.TokenTypeLessThanOrEqual,
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
			token, err := tokenizer.handleLessThanSign(1)

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

func TestHandleLessThanSignErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "unexpected end of expression",
			input:    "",
			expected: "unexpected end of expression at position 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewTokenizer(test.input).handleLessThanSign(1)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
