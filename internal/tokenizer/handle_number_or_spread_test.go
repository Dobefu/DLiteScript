package tokenizer

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestHandleNumberOrSpread(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []rune
		expected *token.Token
	}{
		{
			name:  "one dot",
			input: []rune(". "),
			expected: token.NewToken(
				".",
				token.TokenTypeDot,
				0,
				0,
			),
		},
		{
			name:  "dot with number",
			input: []rune(".1"),
			expected: token.NewToken(
				".1",
				token.TokenTypeNumber,
				0,
				0,
			),
		},
		{
			name:  "spread",
			input: []rune("..."),
			expected: token.NewToken(
				"...",
				token.TokenTypeOperationSpread,
				0,
				0,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tokenizer := NewTokenizer(string(test.input))
			_, _ = tokenizer.GetNext()
			token, err := tokenizer.handleNumberOrSpread(test.input[0], 0)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if token.Atom != test.expected.Atom {
				t.Fatalf("expected %s, got %s", test.expected.Atom, token.Atom)
			}
		})
	}
}

func TestHandleNumberOrSpreadErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []rune
		expected string
	}{
		{
			name:  "unexpected end of expression",
			input: []rune("."),
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tokenizer := NewTokenizer(string(test.input))
			tokenizer.isEOF = false
			_, _ = tokenizer.GetNext()

			var err error

			if len(test.input) > 0 {
				_, err = tokenizer.handleNumberOrSpread(test.input[0], 0)
			} else {
				_, err = tokenizer.handleNumberOrSpread(0, 0)
			}

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, err.Error())
			}
		})
	}
}
