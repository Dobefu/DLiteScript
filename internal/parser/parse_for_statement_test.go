package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseForStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "infinite loop",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for { () }",
		},
		{
			name: "loop with condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "0", TokenType: token.TokenTypeNumber},
				{Atom: "<", TokenType: token.TokenTypeLessThan},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for (0 < 1) { () }",
		},
		{
			name: "loop with to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "to", TokenType: token.TokenTypeTo},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for from 0 to 1 { () }",
		},
		{
			name: "loop with from and to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "from", TokenType: token.TokenTypeFrom},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "to", TokenType: token.TokenTypeTo},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for from 1 to 2 { () }",
		},
		{
			name: "loop with variable declaration and condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "i", TokenType: token.TokenTypeIdentifier},
				{Atom: "<", TokenType: token.TokenTypeLessThan},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for var i (i < 1) { () }",
		},
		{
			name: "loop with variable declaration and to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "i", TokenType: token.TokenTypeIdentifier},
				{Atom: "to", TokenType: token.TokenTypeTo},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for var i from 0 to 1 { () }",
		},
		{
			name: "loop with variable declaration and from and to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "i", TokenType: token.TokenTypeIdentifier},
				{Atom: "from", TokenType: token.TokenTypeFrom},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "to", TokenType: token.TokenTypeTo},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for var i from 1 to 2 { () }",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			expr, err := p.Parse()

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if expr.Expr() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, expr.Expr())
			}
		})
	}
}

func TestParseForStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "no next token after for",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 1",
		},
		{
			name: "no next token after infinite loop body start",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 2",
		},
		{
			name: "no next token after condition loop body start",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "0", TokenType: token.TokenTypeNumber},
				{Atom: "<", TokenType: token.TokenTypeLessThan},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 5",
		},
		{
			name: "no next token after to",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "to", TokenType: token.TokenTypeTo},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 2",
		},
		{
			name: "no next token after from",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "from", TokenType: token.TokenTypeFrom},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 2",
		},
		{
			name: "no next token after condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "0", TokenType: token.TokenTypeNumber},
				{Atom: "<", TokenType: token.TokenTypeLessThan},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 3",
		},
		{
			name: "no next token in condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "0", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 2",
		},
		{
			name: "no next token after variable declaration",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "i", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeNumber},
				{Atom: "to", TokenType: token.TokenTypeTo},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "to") + " at position 5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.Parse()

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, err.Error())
			}
		})
	}
}
