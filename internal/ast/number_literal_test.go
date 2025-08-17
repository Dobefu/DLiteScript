package ast

import (
	"testing"
)

func TestNumberLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         ExprNode
		expectedValue string
		expectedPos   int
	}{
		{
			input:         &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			expectedValue: "1",
			expectedPos:   0,
		},
		{
			input:         &NumberLiteral{Value: "1", StartPos: 1, EndPos: 2},
			expectedValue: "1",
			expectedPos:   1,
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
		}

		if test.input.StartPosition() != test.expectedPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedPos, test.input.StartPosition())
		}
	}
}
