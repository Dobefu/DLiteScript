package parser

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

func TestParseBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected *ast.BinaryExpr
	}{
		{
			input: "1 + 1",
			expected: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
	}

	for _, test := range tests {
		to := tokenizer.NewTokenizer(test.input)
		tokens, _ := to.Tokenize()

		parser := NewParser(tokens)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got \"%v\"", err)

			continue
		}

		if result.Expr() != test.expected.Expr() {
			t.Errorf(
				"expected \"%s\", got \"%s\"",
				test.expected.Expr(),
				result.Expr(),
			)
		}
	}
}

func TestParseBinaryExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		operatorToken *token.Token
		leftExpr      ast.ExprNode
		rightToken    *token.Token
		expected      string
	}{
		{
			operatorToken: &token.Token{
				Atom:      "+",
				TokenType: token.TokenTypeOperationAdd,
				StartPos:  0,
				EndPos:    0,
			},
			leftExpr: nil,
			rightToken: &token.Token{
				Atom:      "1",
				TokenType: token.TokenTypeNumber,
				StartPos:  0,
				EndPos:    0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "+"),
		},
		{
			operatorToken: &token.Token{
				Atom:      "/",
				TokenType: token.TokenTypeOperationDiv,
				StartPos:  0,
				EndPos:    0,
			},
			leftExpr: &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			rightToken: &token.Token{
				Atom:      "/",
				TokenType: token.TokenTypeOperationDiv,
				StartPos:  0,
				EndPos:    0,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "/"),
		},
	}

	for _, test := range tests {
		_, err := NewParser([]*token.Token{}).parseBinaryExpr(
			test.operatorToken,
			test.leftExpr,
			test.rightToken,
			0,
		)

		if err == nil {
			t.Fatal("expected error, got none")
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

func BenchmarkParseBinaryExpr(b *testing.B) {
	for b.Loop() {
		p := NewParser([]*token.Token{})

		_, _ = p.parseBinaryExpr(
			&token.Token{
				Atom:      "1",
				TokenType: token.TokenTypeOperationAdd,
				StartPos:  0,
				EndPos:    0,
			},
			nil,
			nil,
			0,
		)
	}
}
