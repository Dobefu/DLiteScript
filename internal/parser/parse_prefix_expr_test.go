package parser

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParsePrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected *ast.PrefixExpr
	}{
		{
			name: "positive number",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 1, Line: 0, Column: 0},
						End:   ast.Position{Offset: 2, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "negative identifier",
			input: []*token.Token{
				{Atom: "-", TokenType: token.TokenTypeOperationAdd},
				{Atom: "PI", TokenType: token.TokenTypeIdentifier},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.NumberLiteral{
					Value: "PI",
					Range: ast.Range{
						Start: ast.Position{Offset: 1, Line: 0, Column: 0},
						End:   ast.Position{Offset: 2, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "negative function call",
			input: []*token.Token{
				{Atom: "-", TokenType: token.TokenTypeOperationAdd},
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.NumberLiteral{
					Value: "abs(1)",
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "positive string",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "test", TokenType: token.TokenTypeString},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.StringLiteral{
					Value: "test",
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "null literal",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "null", TokenType: token.TokenTypeNull},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.NullLiteral{
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "parenthesized expression",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "function call",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "printf", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "test", TokenType: token.TokenTypeString},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "printf",
					Arguments: []ast.ExprNode{
						&ast.StringLiteral{
							Value: "test",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "namespaced function call",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "math", TokenType: token.TokenTypeIdentifier},
				{Atom: ".", TokenType: token.TokenTypeDot},
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.FunctionCall{
					Namespace:    "math",
					FunctionName: "abs",
					Arguments: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "identifier",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "PI", TokenType: token.TokenTypeIdentifier},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.Identifier{
					Value: "PI",
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "namespaced identifier",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "math", TokenType: token.TokenTypeIdentifier},
				{Atom: ".", TokenType: token.TokenTypeDot},
				{Atom: "PI", TokenType: token.TokenTypeIdentifier},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.Identifier{
					Value: "math.PI",
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		},
		{
			name: "namespaced identifier followed by newline",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "math", TokenType: token.TokenTypeIdentifier},
				{Atom: ".", TokenType: token.TokenTypeDot},
				{Atom: "PI", TokenType: token.TokenTypeIdentifier},
				{Atom: "\n", TokenType: token.TokenTypeNewline},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &ast.Identifier{
					Value: "math.PI",
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
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

			if result.Expr() != test.expected.Expr() {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.Expr(),
					result.Expr(),
				)
			}
		})
	}
}

func TestParsePrefixExprErr(t *testing.T) {
	t.Parallel()

	errNextTokenAfterEOF := errorutil.ErrorMsgUnexpectedEOF

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "two plus signs",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			name: "open parenthesis",
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			name: "closing parenthesis",
			input: []*token.Token{
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, ")"),
		},
		{
			name: "unclosed parenthesis",
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
			name: "two plus signs in parenthesis",
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			name: "unclosed parenthesis in parenthesis",
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "(", TokenType: token.TokenTypeLParen},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgExpectedCloseParen, "("),
		},
		{
			name: "invalid token after dot",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "math", TokenType: token.TokenTypeIdentifier},
				{Atom: ".", TokenType: token.TokenTypeDot},
				{Atom: "123", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "123"),
		},
		{
			name: "EOF after dot",
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "math", TokenType: token.TokenTypeIdentifier},
				{Atom: ".", TokenType: token.TokenTypeDot},
			},
			expected: errNextTokenAfterEOF,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewParser(test.input).Parse()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Errorf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}

func TestParseParenthesizedExprErr(t *testing.T) {
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
			_, err := p.parseParenthesizedExpr(0)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
