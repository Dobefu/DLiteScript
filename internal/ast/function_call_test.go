package ast

import (
	"testing"
)

func TestFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         ExprNode
		expectedValue string
		expectedPos   int
	}{
		{
			input: &FunctionCall{
				FunctionName: "abs",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", Pos: 0},
				},
				Pos: 0,
			},
			expectedValue: "abs(1)",
			expectedPos:   0,
		},
		{
			input: &FunctionCall{
				FunctionName: "max",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", Pos: 0},
					&NumberLiteral{Value: "2", Pos: 2},
				},
				Pos: 0,
			},
			expectedValue: "max(1, 2)",
			expectedPos:   0,
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
		}

		if test.input.Position() != test.expectedPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedPos, test.input.Position())
		}
	}
}
