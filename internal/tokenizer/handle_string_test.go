package tokenizer

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestHandleString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected *token.Token
	}{
		{
			name:  "simple string",
			input: `"test"`,
			expected: token.NewToken(
				"test",
				token.TokenTypeString,
				0,
				0,
			),
		},
		{
			name:  "string with escape",
			input: `"te\"st"`,
			expected: token.NewToken(
				"te\"st",
				token.TokenTypeString,
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
			token, err := tokenizer.handleString(1)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if token.Atom != test.expected.Atom {
				t.Fatalf("expected %s, got %s", test.expected.Atom, token.Atom)
			}
		})
	}
}

func TestHandleStringErr(t *testing.T) {
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
				"%s: %s at position 0",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewTokenizer(test.input).handleString(1)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, err.Error())
			}
		})
	}
}
