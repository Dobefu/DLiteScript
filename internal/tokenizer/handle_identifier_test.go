package tokenizer

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestHandleIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []rune
		expected *token.Token
	}{
		{
			name:  "simple identifier",
			input: []rune("test"),
			expected: token.NewToken(
				"test",
				token.TokenTypeIdentifier,
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
			token, err := tokenizer.handleIdentifier(test.input[0], 0)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if token.Atom != test.expected.Atom {
				t.Fatalf("expected %s, got %s", test.expected.Atom, token.Atom)
			}
		})
	}
}

func TestHandleIdentifierErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []rune
		expected string
	}{
		{
			name:  "unexpected end of expression",
			input: []rune(""),
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgInvalidUTF8Char,
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
				_, err = tokenizer.handleIdentifier(test.input[0], 0)
			} else {
				_, err = tokenizer.handleIdentifier(0, 0)
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
