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
		input    []*token.Token
		expected *ast.PrefixExpr
	}{
		{
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				Operand:  &ast.NumberLiteral{Value: "1", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: []*token.Token{
				{Atom: "-", TokenType: token.TokenTypeOperationAdd},
				{Atom: "PI", TokenType: token.TokenTypeIdentifier},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{Atom: "-", TokenType: token.TokenTypeOperationSub},
				Operand:  &ast.NumberLiteral{Value: "PI", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: []*token.Token{
				{Atom: "-", TokenType: token.TokenTypeOperationAdd},
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{Atom: "-", TokenType: token.TokenTypeOperationSub},
				Operand:  &ast.NumberLiteral{Value: "abs(1)", StartPos: 4, EndPos: 5},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "test", TokenType: token.TokenTypeString},
			},
			expected: &ast.PrefixExpr{
				Operator: token.Token{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				Operand:  &ast.StringLiteral{Value: "test", StartPos: 4, EndPos: 5},
				StartPos: 0,
				EndPos:   0,
			},
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got \"%s\"", err.Error())

			continue
		}

		if result.Expr() != test.expected.Expr() {
			t.Errorf("expected \"%s\", got \"%s\"", test.expected.Expr(), result.Expr())
		}
	}
}

func TestParsePrefixExprErr(t *testing.T) {
	t.Parallel()

	errNextTokenAfterEOF := errorutil.ErrorMsgUnexpectedEOF

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input: []*token.Token{
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
			},
			expected: errNextTokenAfterEOF,
		},
		{
			input: []*token.Token{
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, ")"),
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgParenNotClosedAtEOF,
		},
		{
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
		_, err := NewParser(test.input).Parse()

		if err == nil {
			t.Fatalf("expected error for \"%v\", got none", test.input)
		}

		if errors.Unwrap(err).Error() != test.expected {
			t.Errorf(
				"expected error \"%s\", got \"%s\"",
				test.expected,
				errors.Unwrap(err).Error(),
			)
		}
	}
}

func BenchmarkParsePrefixExpr(b *testing.B) {
	for b.Loop() {
		_, _ = NewParser([]*token.Token{
			{Atom: "-", TokenType: token.TokenTypeOperationSub},
			{Atom: "1", TokenType: token.TokenTypeNumber},
		}).Parse()
	}
}
