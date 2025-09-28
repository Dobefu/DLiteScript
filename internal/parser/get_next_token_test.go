package parser

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestGetNextToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		tokens   []*token.Token
		expected string
	}{
		{
			name: "normal token retrieval",
			tokens: []*token.Token{
				token.NewToken("hello", token.TokenTypeIdentifier, 0, 5),
			},
			expected: "hello",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.tokens)
			token, err := p.GetNextToken()

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if token.Atom != test.expected {
				t.Fatalf(
					"expected token atom '%s', got '%s'",
					test.expected,
					token.Atom,
				)
			}
		})
	}
}

func TestGetNextTokenErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		tokens []*token.Token
		isEOF  bool
	}{
		{
			name:   "EOF error when isEOF is true and no tokens",
			tokens: []*token.Token{},
			isEOF:  true,
		},
		{
			name: "EOF error when isEOF is true and tokenIdx < len(tokens)",
			tokens: []*token.Token{
				token.NewToken("hello", token.TokenTypeIdentifier, 0, 5),
			},
			isEOF: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.tokens)
			p.isEOF = test.isEOF

			_, err := p.GetNextToken()

			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
