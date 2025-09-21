package parser

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:     "empty",
			input:    []*token.Token{},
			expected: "",
		},
		{
			name: "underscore",
			input: []*token.Token{
				{Atom: "_", TokenType: token.TokenTypeNumber},
			},
			expected: "_",
		},
		{
			name: "number",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "1",
		},
		{
			name: "number with newlines",
			input: []*token.Token{
				{Atom: "\n", TokenType: token.TokenTypeNewline},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "\n", TokenType: token.TokenTypeNewline},
			},
			expected: "1",
		},
		{
			name: "variable declaration",
			input: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "var x number = 1",
		},
		{
			name: "constant declaration",
			input: []*token.Token{
				{Atom: "const", TokenType: token.TokenTypeConst},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "const x number = 1",
		},
		{
			name: "if statement",
			input: []*token.Token{
				{Atom: "if", TokenType: token.TokenTypeIf},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "if 1 { (1) }",
		},
		{
			name: "for loop with break",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "break", TokenType: token.TokenTypeBreak},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for { (break) }",
		},
		{
			name: "for loop with continue",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "continue", TokenType: token.TokenTypeContinue},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "for { (continue) }",
		},
		{
			name: "block statement",
			input: []*token.Token{
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "(1)",
		},
		{
			name: "function declaration",
			input: []*token.Token{
				{Atom: "func", TokenType: token.TokenTypeFunc},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: "func x() number",
		},
		{
			name: "import statement",
			input: []*token.Token{
				{Atom: "import", TokenType: token.TokenTypeImport},
				{Atom: "test", TokenType: token.TokenTypeString},
			},
			expected: "import \"test\"",
		},
		{
			name: "multiple statements",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "\n", TokenType: token.TokenTypeNewline},
				{Atom: "2", TokenType: token.TokenTypeNumber},
			},
			expected: "1\n2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			result, err := parser.Parse()

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if len(test.input) == 0 {
				if result != nil {
					t.Fatalf("expected nil result, got \"%s\"", result.Expr())
				}

				return
			}

			if result.Expr() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, result.Expr())
			}
		})
	}
}

func TestParseErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:     "empty",
			input:    []*token.Token{{}},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			name: "unexpected EOF",
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			name: "unexpected operation",
			input: []*token.Token{
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "/"),
		},
		{
			name: "unexpected token",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "2", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "2"),
		},
		{
			name: "block statement unexpected EOF",
			input: []*token.Token{
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)

			// Test the EOF error.
			if len(test.input) == 1 && test.input[0].Atom == "_" {
				p.isEOF = true
			}

			_, err := p.Parse()

			if err == nil {
				t.Fatalf("expected error for \"%v\", got none", test.input)
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Fatalf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}

func TestParseStatementErr(t *testing.T) {
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
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseStatement()

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for b.Loop() {
		p := NewParser(
			[]*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				{Atom: "3", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "4", TokenType: token.TokenTypeNumber},
			},
		)
		_, _ = p.Parse()
	}
}
