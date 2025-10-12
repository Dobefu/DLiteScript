package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
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
			name: "no return type",
			input: []*token.Token{
				{Atom: "greet", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 5},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 5, EndPos: 6},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 6, EndPos: 7},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 7, EndPos: 8},
				{Atom: "return", TokenType: token.TokenTypeReturn, StartPos: 8, EndPos: 14},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 14, EndPos: 15},
			},
			expected: &ast.FuncDeclarationStatement{
				Name: "greet",
				Args: []ast.FuncParameter{},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values:    []ast.ExprNode{},
							NumValues: 0,
							Range: ast.Range{
								Start: ast.Position{Offset: 8, Line: 0, Column: 0},
								End:   ast.Position{Offset: 14, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 7, Line: 0, Column: 0},
						End:   ast.Position{Offset: 15, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 15, Line: 0, Column: 0},
				},
			},
		},
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
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 33, Line: 0, Column: 0},
										End:   ast.Position{Offset: 34, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 27, Line: 0, Column: 0},
								End:   ast.Position{Offset: 34, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 26, Line: 0, Column: 0},
						End:   ast.Position{Offset: 35, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 35, Line: 0, Column: 0},
				},
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
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 33, Line: 0, Column: 0},
										End:   ast.Position{Offset: 34, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 27, Line: 0, Column: 0},
								End:   ast.Position{Offset: 34, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 26, Line: 0, Column: 0},
						End:   ast.Position{Offset: 35, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number", "string"},
				NumReturnValues: 2,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 35, Line: 0, Column: 0},
				},
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
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 42, Line: 0, Column: 0},
										End:   ast.Position{Offset: 43, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 36, Line: 0, Column: 0},
								End:   ast.Position{Offset: 43, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 35, Line: 0, Column: 0},
						End:   ast.Position{Offset: 44, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number", "string"},
				NumReturnValues: 2,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 44, Line: 0, Column: 0},
				},
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

			if expr.GetRange().Start.Offset != test.expected.GetRange().Start.Offset {
				t.Errorf(
					"expected %d, got %d",
					test.expected.GetRange().Start.Offset,
					expr.GetRange().Start.Offset,
				)
			}

			if expr.GetRange().End.Offset != test.expected.GetRange().End.Offset {
				t.Errorf(
					"expected %d, got %d",
					test.expected.GetRange().End.Offset,
					expr.GetRange().End.Offset,
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
			name:  "missing name",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "name is not an identifier",
			input: []*token.Token{
				{Atom: "123", TokenType: token.TokenTypeNumber, StartPos: 0, EndPos: 3},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 4",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "123"),
			),
		},
		{
			name: "missing closing tag",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 5",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "missing argument type",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 6",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token instead of closing paren",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 3, EndPos: 4},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 5",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "a"),
			),
		},
		{
			name: "unexpected token instead of RParen",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: "a", TokenType: token.TokenTypeIdentifier, StartPos: 4, EndPos: 5},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 6, EndPos: 7},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 7",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "{"),
			),
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
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 28",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
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
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 29",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid body",
			input: []*token.Token{
				{Atom: "add", TokenType: token.TokenTypeIdentifier, StartPos: 0, EndPos: 3},
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 3, EndPos: 4},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 4, EndPos: 5},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber, StartPos: 5, EndPos: 11},
				{Atom: "{", TokenType: token.TokenTypeLBrace, StartPos: 11, EndPos: 12},
				{Atom: "if", TokenType: token.TokenTypeIf, StartPos: 12, EndPos: 14},
				{Atom: "true", TokenType: token.TokenTypeBool, StartPos: 14, EndPos: 18},
				{Atom: "return", TokenType: token.TokenTypeReturn, StartPos: 18, EndPos: 24},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 24, EndPos: 25},
				{Atom: "}", TokenType: token.TokenTypeRBrace, StartPos: 25, EndPos: 26},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 27",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
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

func TestGetArgsErr(t *testing.T) {
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

			parser := NewParser(test.input)
			_, err := parser.getArgs()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseFunctionArgumentErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *token.Token
		expected string
	}{
		{
			name: "name is not an identifier",
			input: &token.Token{
				Atom:      "1",
				TokenType: token.TokenTypeNumber,
				StartPos:  0,
				EndPos:    1,
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "1"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser([]*token.Token{test.input})
			_, err := parser.parseFunctionArgument(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestGetReturnTypesErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "invalid return type in parentheses",
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen, StartPos: 0, EndPos: 1},
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 0, EndPos: 6},
				{Atom: ")", TokenType: token.TokenTypeRParen, StartPos: 6, EndPos: 7},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 3",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "1"),
			),
		},
		{
			name: "invalid return type",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber, StartPos: 0, EndPos: 6},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "1"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			_, err := parser.getReturnTypes()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseReturnTypes(t *testing.T) {
	t.Parallel()

	parser := NewParser([]*token.Token{
		{
			Atom:      "number",
			TokenType: token.TokenTypeTypeNumber,
			StartPos:  0,
			EndPos:    6,
		},
	})

	parser.tokenIdx = 1
	parser.isEOF = false

	result, err := parser.parseReturnTypes(token.TokenTypeLBrace)

	if err != nil {
		t.Fatalf("expected no error, got: %s", err.Error())
	}

	if len(result) != 0 {
		t.Errorf("expected empty result, got: %v", result)
	}
}
