package ast

import (
	"testing"
)

func TestArrayLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *ArrayLiteral
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "empty array literal",
			input: &ArrayLiteral{
				Values:   []ExprNode{},
				StartPos: 0,
				EndPos:   0,
			},
			expectedNodes:    []string{"[]"},
			expectedStartPos: 0,
			expectedEndPos:   0,
			continueOn:       "",
		},
		{
			name: "array literal",
			input: &ArrayLiteral{
				Values: []ExprNode{
					&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					&NumberLiteral{Value: "1", StartPos: 1, EndPos: 2},
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedNodes:    []string{"[1, 1]", "1", "1", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "",
		},
		{
			name: "array literal with nil value",
			input: &ArrayLiteral{
				Values: []ExprNode{
					nil,
				},
				StartPos: 0,
				EndPos:   0,
			},
			expectedNodes:    []string{"[]"},
			expectedStartPos: 0,
			expectedEndPos:   0,
			continueOn:       "",
		},
		{
			name: "walk early return after array node",
			input: &ArrayLiteral{
				Values: []ExprNode{
					&NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
					&NumberLiteral{Value: "24", StartPos: 2, EndPos: 4},
				},
				StartPos: 0,
				EndPos:   4,
			},
			expectedNodes:    []string{"[42, 24]"},
			expectedStartPos: 0,
			expectedEndPos:   4,
			continueOn:       "[42, 24]",
		},
		{
			name: "walk early return after first value",
			input: &ArrayLiteral{
				Values: []ExprNode{
					&NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
					&NumberLiteral{Value: "24", StartPos: 2, EndPos: 4},
				},
				StartPos: 0,
				EndPos:   4,
			},
			expectedNodes:    []string{"[42, 24]", "42"},
			expectedStartPos: 0,
			expectedEndPos:   4,
			continueOn:       "42",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.StartPosition() != test.expectedStartPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedStartPos,
					test.input.StartPosition(),
				)
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedEndPos,
					test.input.EndPosition(),
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
