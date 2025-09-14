package ast

import (
	"testing"
)

func TestBlockStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *BlockStatement
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "block statement with single statement",
			input: &BlockStatement{
				Statements: []ExprNode{
					&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"(1)", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "block statement with nil statement",
			input: &BlockStatement{
				Statements: []ExprNode{nil},
				StartPos:   0,
				EndPos:     0,
			},
			expectedNodes:    []string{"()"},
			expectedStartPos: 0,
			expectedEndPos:   0,
			continueOn:       "",
		},
		{
			name: "empty block statement",
			input: &BlockStatement{
				Statements: []ExprNode{},
				StartPos:   0,
				EndPos:     0,
			},
			expectedNodes:    []string{"()"},
			expectedStartPos: 0,
			expectedEndPos:   0,
			continueOn:       "",
		},
		{
			name: "block statement with multiple statements",
			input: &BlockStatement{
				Statements: []ExprNode{
					&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					&NumberLiteral{Value: "2", StartPos: 1, EndPos: 2},
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedNodes:    []string{"(1 2)", "1", "1", "2", "2"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "",
		},
		{
			name: "walk early return after block node",
			input: &BlockStatement{
				Statements: []ExprNode{
					&NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedNodes:    []string{"(42)"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "(42)",
		},
		{
			name: "walk early return after first statement",
			input: &BlockStatement{
				Statements: []ExprNode{
					&NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
					&NumberLiteral{Value: "24", StartPos: 2, EndPos: 4},
				},
				StartPos: 0,
				EndPos:   4,
			},
			expectedNodes:    []string{"(42 24)", "42"},
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
