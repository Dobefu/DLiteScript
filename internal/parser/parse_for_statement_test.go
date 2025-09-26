package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
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
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 4, EndPos: 5},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 5, EndPos: 6},
			},
			expected: "for { () }",
		},
		{
			name: "loop with condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "0", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
				{Atom: "<", TokenType: token.TokenTypeLessThan, StartPos: 6, EndPos: 7},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 8, EndPos: 9},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 10, EndPos: 11},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 11, EndPos: 12},
			},
			expected: "for (0 < 1) { () }",
		},
		{
			name: "loop with to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "to", TokenType: token.TokenTypeTo, StartPos: 4, EndPos: 6},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 7, EndPos: 8},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 9, EndPos: 10},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 10, EndPos: 11},
			},
			expected: "for from 0 to 1 { () }",
		},
		{
			name: "loop with from and to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "from", TokenType: token.TokenTypeFrom, StartPos: 4, EndPos: 8},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 9, EndPos: 10},
				{Atom: "to", TokenType: token.TokenTypeTo, StartPos: 11, EndPos: 13},
				{Atom: "2", TokenType: token.TokenTypeNumber, StartPos: 14, EndPos: 15},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 16, EndPos: 17},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 17, EndPos: 18},
			},
			expected: "for from 1 to 2 { () }",
		},
		{
			name: "loop with variable declaration and condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "var", TokenType: token.TokenTypeVar, StartPos: 4, EndPos: 7},
				{Atom: "i", TokenType: token.TokenTypeIdentifier, StartPos: 8, EndPos: 9},
				{Atom: "<", TokenType: token.TokenTypeLessThan, StartPos: 10, EndPos: 11},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 12, EndPos: 13},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 14, EndPos: 15},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 15, EndPos: 16},
			},
			expected: "for var i (i < 1) { () }",
		},
		{
			name: "loop with variable declaration and to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "var", TokenType: token.TokenTypeVar, StartPos: 4, EndPos: 7},
				{Atom: "i", TokenType: token.TokenTypeIdentifier, StartPos: 8, EndPos: 9},
				{Atom: "to", TokenType: token.TokenTypeTo, StartPos: 10, EndPos: 12},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 13, EndPos: 14},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 15, EndPos: 16},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 16, EndPos: 17},
			},
			expected: "for var i from 0 to 1 { () }",
		},
		{
			name: "loop with variable declaration and from and to range",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "var", TokenType: token.TokenTypeVar, StartPos: 4, EndPos: 7},
				{Atom: "i", TokenType: token.TokenTypeIdentifier, StartPos: 8, EndPos: 9},
				{Atom: "from", TokenType: token.TokenTypeFrom, StartPos: 10, EndPos: 14},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 15, EndPos: 16},
				{Atom: "to", TokenType: token.TokenTypeTo, StartPos: 17, EndPos: 19},
				{Atom: "2", TokenType: token.TokenTypeNumber, StartPos: 20, EndPos: 21},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 22, EndPos: 23},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 23, EndPos: 24},
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
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 4",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after infinite loop body start",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 4, EndPos: 5},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 5",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after condition loop body start",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "0", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
				{Atom: "<", TokenType: token.TokenTypeLessThan, StartPos: 6, EndPos: 7},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 8, EndPos: 9},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 10, EndPos: 11},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 8",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after to",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "to", TokenType: token.TokenTypeTo, StartPos: 4, EndPos: 6},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 6",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after from",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "from", TokenType: token.TokenTypeFrom, StartPos: 4, EndPos: 8},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 8",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "0", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
				{Atom: "<", TokenType: token.TokenTypeLessThan, StartPos: 6, EndPos: 7},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 6",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token in condition",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "0", TokenType: token.TokenTypeNumber, StartPos: 4, EndPos: 5},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 5",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after variable declaration",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "var", TokenType: token.TokenTypeVar, StartPos: 4, EndPos: 7},
				{Atom: "i", TokenType: token.TokenTypeIdentifier, StartPos: 8, EndPos: 9},
				{Atom: "number", TokenType: token.TokenTypeNumber, StartPos: 10, EndPos: 16},
				{Atom: "to", TokenType: token.TokenTypeTo, StartPos: 17, EndPos: 19},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 16",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "to"),
			),
		},
		{
			name: "missing loop body",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "var", TokenType: token.TokenTypeVar, StartPos: 4, EndPos: 7},
				{Atom: "i", TokenType: token.TokenTypeIdentifier, StartPos: 8, EndPos: 9},
				{Atom: "=", TokenType: token.TokenTypeEqual, StartPos: 10, EndPos: 11},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 12, EndPos: 13},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 10",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseLoop(ast.Position{Offset: 0, Line: 0, Column: 0})

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
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 0, EndPos: 1},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after variable declaration",
			input: []*token.Token{
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 8",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 5",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 9",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseVariableDeclarationLoop(
				ast.Position{Offset: 0, Line: 0, Column: 0},
			)

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
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no next token after variable declaration",
			input: []*token.Token{
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 8",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 8",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 11",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseExplicitRangeLoop(
				ast.Position{Offset: 0, Line: 0, Column: 0},
				"",
			)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseForStatementWithExplicitRangeLoopErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "incomplete from expression in explicit range loop",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "var", TokenType: token.TokenTypeVar, StartPos: 4, EndPos: 7},
				{Atom: "i", TokenType: token.TokenTypeIdentifier, StartPos: 8, EndPos: 9},
				{Atom: "from", TokenType: token.TokenTypeFrom, StartPos: 10, EndPos: 14},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 15, EndPos: 16},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 13",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "incomplete to expression in explicit range loop",
			input: []*token.Token{
				{Atom: "for", TokenType: token.TokenTypeFor, StartPos: 0, EndPos: 3},
				{Atom: "var", TokenType: token.TokenTypeVar, StartPos: 4, EndPos: 7},
				{Atom: "i", TokenType: token.TokenTypeIdentifier, StartPos: 8, EndPos: 9},
				{Atom: "from", TokenType: token.TokenTypeFrom, StartPos: 10, EndPos: 14},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 15, EndPos: 16},
				{Atom: "to", TokenType: token.TokenTypeTo, StartPos: 17, EndPos: 19},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 20, EndPos: 21},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 16",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseImplicitRangeLoopWithVariable(
				ast.Position{Offset: 0, Line: 0, Column: 0},
				"",
			)

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
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 8",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 3",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 3",
				errorutil.StageParse.String(),
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
				"%s: %s line 1 at position 5",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.parseRangeLoopWithoutVariable(
				ast.Position{Offset: 0, Line: 0, Column: 0},
			)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
