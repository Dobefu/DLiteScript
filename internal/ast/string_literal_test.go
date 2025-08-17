package ast

import (
	"testing"
)

func TestStringLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         ExprNode
		expectedValue string
		expectedPos   int
	}{
		{
			input:         &StringLiteral{Value: "test", StartPos: 0, EndPos: 1},
			expectedValue: `"test"`,
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
