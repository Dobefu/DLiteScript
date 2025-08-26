package tokenizer

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseSpread(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected []*token.Token
	}{
		{
			input: "...",
			expected: []*token.Token{
				token.NewToken("...", token.TokenTypeOperationSpread, 0, 0),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
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
		input    string
		expected string
	}{
		{
			input:    "..",
			expected: "unexpected end of expression at position 2",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			tokenizer := NewTokenizer(test.input)
			_, err := tokenizer.Tokenize()

			if err == nil {
				t.Fatalf("Expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("Expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
