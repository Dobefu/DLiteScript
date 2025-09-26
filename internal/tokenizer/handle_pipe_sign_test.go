package tokenizer

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestHandlePipeSign(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected *token.Token
	}{
		{
			name:  "pipe sign after pipe sign",
			input: "||",
			expected: token.NewToken(
				"||",
				token.TokenTypeLogicalOr,
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
			token, err := tokenizer.handlePipeSign(1)

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

func TestHandlePipeSignErr(t *testing.T) {
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
		{
			name:  "unexpected character",
			input: "*",
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageTokenize.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedChar, "*"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewTokenizer(test.input).handlePipeSign(1)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
