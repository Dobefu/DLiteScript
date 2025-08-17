package ast

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         ExprNode
		expectedValue string
		expectedPos   int
	}{
		{
			input:         &Identifier{Value: "PI", StartPos: 0, EndPos: 1},
			expectedValue: "PI",
			expectedPos:   0,
		},
		{
			input:         &Identifier{Value: "PI", StartPos: 1, EndPos: 2},
			expectedValue: "PI",
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
