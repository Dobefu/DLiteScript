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
				Operand:  &ast.NumberLiteral{Value: "1", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
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
				Operand:  &ast.NumberLiteral{Value: "PI", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
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
				Operand:  &ast.NumberLiteral{Value: "abs(1)", StartPos: 4, EndPos: 5},
				StartPos: 0,
				EndPos:   0,
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
				Operand:  &ast.StringLiteral{Value: "test", StartPos: 4, EndPos: 5},
				StartPos: 0,
				EndPos:   0,
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
