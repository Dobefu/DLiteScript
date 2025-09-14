package tokenizer

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestHandleAsteriskSign(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected *token.Token
	}{
		{
			name:  "no extra asterisk",
			input: "* ",
			expected: token.NewToken(
				"*",
				token.TokenTypeOperationMul,
				0,
				0,
			),
		},
		{
			name:  "extra asterisk",
			input: "**",
			expected: token.NewToken(
				"**",
				token.TokenTypeOperationPow,
				0,
				0,
			),
		},
		{
			name:  "equals sign after asterisk",
			input: "*=",
			expected: token.NewToken(
				"*=",
				token.TokenTypeOperationMulAssign,
				0,
				0,
			),
		},
		{
			name:  "extra asterisk and equals sign",
			input: "**=",
			expected: token.NewToken(
				"**=",
				token.TokenTypeOperationPowAssign,
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
			token, err := tokenizer.handleAsteriskSign(1)

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
