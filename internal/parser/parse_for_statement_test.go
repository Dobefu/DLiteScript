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
		{
			name: "missing loop body",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "i", TokenType: token.TokenTypeIdentifier},
				{Atom: "=", TokenType: token.TokenTypeEqual},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(
				"%s at position 5",
				errorutil.ErrorMsgUnexpectedEOF,
			),
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

func TestParseLoopErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no tokens",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s at position 0",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseLoop(0)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseLoopBodyErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "no tokens",
			input: []*token.Token{
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: fmt.Sprintf(
				"%s at position 0",
				fmt.Sprintf(
					errorutil.ErrorMsgBlockStatementExpected,
					token.Type(token.TokenTypeRBrace),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseLoopBody()

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseVariableDeclarationLoopErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no tokens",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s at position 0",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after variable declaration",
			input: []*token.Token{
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: fmt.Sprintf(
				"%s at position 1",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "missing identifier after variable declaration",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "test", TokenType: token.TokenTypeString},
			},
			expected: fmt.Sprintf(
				"%s at position 2",
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedIdentifier, "test"),
			),
		},
		{
			name: "unexpected EOF after variable name when expecting operator",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "i", TokenType: token.TokenTypeIdentifier},
			},
			expected: fmt.Sprintf(
				"%s at position 2",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected EOF after variable declaration",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "i", TokenType: token.TokenTypeIdentifier},
				{Atom: "test", TokenType: token.TokenTypeString},
			},
			expected: fmt.Sprintf(
				"%s at position 3",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseVariableDeclarationLoop(0)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseExplicitRangeLoopErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no tokens",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s at position 0",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after variable declaration",
			input: []*token.Token{
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: fmt.Sprintf(
				"%s at position 1",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "missing identifier after variable declaration",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "test", TokenType: token.TokenTypeString},
			},
			expected: fmt.Sprintf(
				"%s at position 2",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token instead of to",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "test", TokenType: token.TokenTypeString},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(
				"%s at position 2",
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "1"),
			),
		},
		{
			name: "missing loop body",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "test", TokenType: token.TokenTypeString},
				{Atom: "to", TokenType: token.TokenTypeTo},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(
				"%s at position 4",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseExplicitRangeLoop(0, "")

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseImplicitRangeLoopWithVariableErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no tokens",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s at position 0",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseImplicitRangeLoopWithVariable(0, "")

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseRangeLoopWithoutVariableErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no tokens",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s at position 0",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after from",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "from", TokenType: token.TokenTypeFrom},
			},
			expected: fmt.Sprintf(
				"%s at position 2",
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "from"),
			),
		},
		{
			name: "error peeking next token after from expression",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "2", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(
				"%s at position 2",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token instead of to after from expression",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "bogus", TokenType: token.TokenTypeIdentifier},
			},
			expected: fmt.Sprintf(
				"%s at position 2",
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "bogus"),
			),
		},
		{
			name: "unexpected token instead of to after from expression",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "to", TokenType: token.TokenTypeTo},
			},
			expected: fmt.Sprintf(
				"%s at position 3",
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseRangeLoopWithoutVariable(0)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
