package tokenizer

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseSpread(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []*token.Token
	}{
		{
			name:  "spread",
			input: "...",
			expected: []*token.Token{
				token.NewToken("...", token.TokenTypeOperationSpread, 0, 0),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tokenizer := NewTokenizer(test.input)
			token, err := tokenizer.parseSpread(0)

			if err != nil {
				t.Fatalf("Error parsing spread: %v", err)
			}

			if token.Atom != test.expected[0].Atom {
				t.Errorf("Expected %s, got %s", test.expected[0].Atom, token.Atom)
			}
		})
	}
}

func TestParseSpreadErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "empty expression",
			input: "",
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgInvalidUTF8Char,
			),
		},
		{
			name:  "unexpected end of expression",
			input: ".",
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tokenizer := NewTokenizer(test.input)
			tokenizer.isEOF = false
			_, err := tokenizer.parseSpread(0)

			if err == nil {
				t.Fatalf("Expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("Expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
