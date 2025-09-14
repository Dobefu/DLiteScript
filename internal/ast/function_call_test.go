package ast

import (
	"testing"
)

func TestFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "function call with single argument",
			input: &FunctionCall{
				Namespace:    "math",
				FunctionName: "abs",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedValue:    "math.abs(1)",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"math.abs(1)", "1", "1"},
			continueOn:       "",
		},
		{
			name: "function call with nil argument",
			input: &FunctionCall{
				Namespace:    "math",
				FunctionName: "abs",
				Arguments:    []ExprNode{nil},
				StartPos:     0,
				EndPos:       1,
			},
			expectedValue:    "math.abs()",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"math.abs()"},
			continueOn:       "",
		},
		{
			name: "function call with multiple arguments",
			input: &FunctionCall{
				Namespace:    "math",
				FunctionName: "max",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					&NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedValue:    "math.max(1, 2)",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"math.max(1, 2)", "1", "1", "2", "2"},
			continueOn:       "",
		},
		{
			name: "function call with no arguments",
			input: &FunctionCall{
				Namespace:    "util",
				FunctionName: "random",
				Arguments:    []ExprNode{},
				StartPos:     0,
				EndPos:       1,
			},
			expectedValue:    "random()",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"random()"},
			continueOn:       "",
		},
		{
			name: "walk early return after function call",
			input: &FunctionCall{
				Namespace:    "math",
				FunctionName: "sqrt",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedValue:    "math.sqrt(42)",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"math.sqrt(42)"},
			continueOn:       "math.sqrt(42)",
		},
		{
			name: "walk early return after first argument",
			input: &FunctionCall{
				Namespace:    "math",
				FunctionName: "pow",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
					&NumberLiteral{Value: "3", StartPos: 1, EndPos: 2},
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedValue:    "math.pow(2, 3)",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"math.pow(2, 3)", "2"},
			continueOn:       "2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Errorf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.StartPosition() != test.expectedStartPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.StartPosition(),
				)
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.EndPosition(),
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
