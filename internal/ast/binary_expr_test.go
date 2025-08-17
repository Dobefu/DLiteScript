package ast

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         ExprNode
		expectedValue string
		expectedPos   int
	}{
		{
			input: &BinaryExpr{
				Left:  &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right: &NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
				},
				StartPos: 0,
				EndPos:   0,
			},
			expectedValue: "(1 + 1)",
			expectedPos:   0,
		},
		{
			input: &BinaryExpr{
				Left:  &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right: &NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
				},
				StartPos: 0,
				EndPos:   0,
			},
			expectedValue: "(1 * 2)",
			expectedPos:   0,
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
		}

		if test.input.StartPosition() != test.expectedPos {
			t.Errorf(
				"expected pos '%d', got '%d'",
				test.expectedPos,
				test.input.StartPosition(),
			)
		}
	}
}
