package ast

import (
	"testing"
)

func TestIndexExpr(t *testing.T) {
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
			name: "index expression",
			input: &IndexExpr{
				Array:    &Identifier{Value: "array", StartPos: 0, EndPos: 5},
				Index:    &NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "array[0]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"array[0]", "array", "array", "0", "0"},
			continueOn:       "",
		},
		{
			name: "index expression with nil array",
			input: &IndexExpr{
				Array:    nil,
				Index:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"", "1", "1"},
			continueOn:       "",
		},
		{
			name: "index expression with nil index",
			input: &IndexExpr{
				Array:    &Identifier{Value: "list", StartPos: 0, EndPos: 4},
				Index:    nil,
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"", "list", "list"},
			continueOn:       "",
		},
		{
			name: "walk early return after index expression",
			input: &IndexExpr{
				Array:    &Identifier{Value: "data", StartPos: 0, EndPos: 4},
				Index:    &NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "data[2]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"data[2]"},
			continueOn:       "data[2]",
		},
		{
			name: "walk early return after array",
			input: &IndexExpr{
				Array:    &Identifier{Value: "items", StartPos: 0, EndPos: 5},
				Index:    &NumberLiteral{Value: "3", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "items[3]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"items[3]", "items"},
			continueOn:       "items",
		},
		{
			name: "walk early return after index",
			input: &IndexExpr{
				Array:    &Identifier{Value: "values", StartPos: 0, EndPos: 6},
				Index:    &NumberLiteral{Value: "4", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "values[4]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"values[4]", "values", "values", "4"},
			continueOn:       "4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Fatalf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.StartPosition() != test.expectedStartPos {
				t.Fatalf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.StartPosition(),
				)
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Fatalf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.EndPosition(),
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
