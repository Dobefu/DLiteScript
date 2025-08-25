package parser

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseFunctionDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected ast.ExprNode
	}{
		{
			name: "simple",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 5, EndPos: 11},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 11, EndPos: 12},
				{Atom: "b", TokenType: token.TokenTypeIdentifier, StartPos: 12, EndPos: 13},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 13, EndPos: 19},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 19, EndPos: 20},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 20, EndPos: 26},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 26, EndPos: 27},
				{Atom: "return", TokenType: token.TokenTypeReturn, StartPos: 27, EndPos: 33},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 33, EndPos: 34},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 34, EndPos: 35},
			},
			expected: &ast.FuncDeclarationStatement{
				Name: "add",
				Args: []ast.FuncParameter{
					{
						Name: "a",
						Type: "number",
					},
					{
						Name: "b",
						Type: "number",
					},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{
									Value:    "1",
									StartPos: 33,
									EndPos:   34,
								},
							},
							NumValues: 1,
							StartPos:  27,
							EndPos:    34,
						},
					},
					StartPos: 26,
					EndPos:   35,
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				StartPos:        0,
				EndPos:          35,
			},
		},
		{
			name: "multiple return types",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 5, EndPos: 11},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 11, EndPos: 12},
				{Atom: "b", TokenType: token.TokenTypeIdentifier, StartPos: 12, EndPos: 13},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 13, EndPos: 19},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 19, EndPos: 20},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 20, EndPos: 26},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 26, EndPos: 27},
				{Atom: "string", TokenType: token.TokenTypeTypeString, StartPos: 27, EndPos: 33},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 26, EndPos: 27},
				{Atom: "return", TokenType: token.TokenTypeReturn, StartPos: 27, EndPos: 33},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 33, EndPos: 34},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 34, EndPos: 35},
			},
			expected: &ast.FuncDeclarationStatement{
				Name: "add",
				Args: []ast.FuncParameter{
					{
						Name: "a",
						Type: "number",
					},
					{
						Name: "b",
						Type: "number",
					},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{
									Value:    "1",
									StartPos: 33,
									EndPos:   34,
								},
							},
							NumValues: 1,
							StartPos:  27,
							EndPos:    34,
						},
					},
					StartPos: 26,
					EndPos:   35,
				},
				ReturnValues:    []string{"number", "string"},
				NumReturnValues: 2,
				StartPos:        0,
				EndPos:          35,
			},
		},
		{
			name: "multiple return types in brackets",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 5, EndPos: 11},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 11, EndPos: 12},
				{Atom: "b", TokenType: token.TokenTypeIdentifier, StartPos: 12, EndPos: 13},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 13, EndPos: 19},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 19, EndPos: 20},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 20, EndPos: 21},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 21, EndPos: 27},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 27, EndPos: 28},
				{Atom: "string", TokenType: token.TokenTypeTypeString, StartPos: 28, EndPos: 34},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 34, EndPos: 35},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 35, EndPos: 36},
				{Atom: "return", TokenType: token.TokenTypeReturn, StartPos: 36, EndPos: 42},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 42, EndPos: 43},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 43, EndPos: 44},
			},
			expected: &ast.FuncDeclarationStatement{
				Name: "add",
				Args: []ast.FuncParameter{
					{
						Name: "a",
						Type: "number",
					},
					{
						Name: "b",
						Type: "number",
					},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{
									Value:    "1",
									StartPos: 42,
									EndPos:   43,
								},
							},
							NumValues: 1,
							StartPos:  36,
							EndPos:    43,
						},
					},
					StartPos: 35,
					EndPos:   44,
				},
				ReturnValues:    []string{"number", "string"},
				NumReturnValues: 2,
				StartPos:        0,
				EndPos:          44,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			expr, err := parser.parseFunctionDeclaration()

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if expr.Expr() != test.expected.Expr() {
				t.Errorf("expected %s, got %s", test.expected.Expr(), expr.Expr())
			}

			if expr.StartPosition() != test.expected.StartPosition() {
				t.Errorf(
					"expected %d, got %d",
					test.expected.StartPosition(),
					expr.StartPosition(),
				)
			}

			if expr.EndPosition() != test.expected.EndPosition() {
				t.Errorf(
					"expected %d, got %d",
					test.expected.EndPosition(),
					expr.EndPosition(),
				)
			}
		})
	}
}

func TestParseFunctionDeclarationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:     "missing name",
			input:    []*token.Token{},
			expected: "unexpected end of expression at position 0",
		},
		{
			name: "missing closing tag",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
			},
			expected: "unexpected end of expression at position 2",
		},
		{
			name: "missing argument type",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
			},
			expected: "unexpected end of expression at position 3",
		},
		{
			name: "unexpected token instead of closing paren",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 3, EndPos: 4},
			},
			expected: "unexpected token: 'a' at position 3",
		},
		{
			name: "unexpected token instead of RParen",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 6, EndPos: 7},
			},
			expected: "unexpected token: '{' at position 6",
		},
		{
			name: "missing closing paren",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 5, EndPos: 11},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 11, EndPos: 12},
				{Atom: "b", TokenType: token.TokenTypeIdentifier, StartPos: 12, EndPos: 13},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 13, EndPos: 19},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 19, EndPos: 20},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 20, EndPos: 26},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 26, EndPos: 27},
			},
			expected: "unexpected end of expression at position 10",
		},
		{
			name: "missing closing paren with parens",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 5, EndPos: 11},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 11, EndPos: 12},
				{Atom: "b", TokenType: token.TokenTypeIdentifier, StartPos: 12, EndPos: 13},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 13, EndPos: 19},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 19, EndPos: 20},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 20, EndPos: 21},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 21, EndPos: 27},
				{Atom: ",", TokenType: token.TokenTypeComma, StartPos: 27, EndPos: 28},
			},
			expected: "unexpected end of expression at position 11",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			_, err := parser.parseFunctionDeclaration()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
