package ast

import (
	"testing"
)

func TestStatementList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            []ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "single statement",
			input: []ExprNode{
				&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			},
			expectedValue:    "1",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"1", "1", "1"},
			continueOn:       "",
		},
		{
			name: "multiple statements",
			input: []ExprNode{
				&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				&NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
			},
			expectedValue:    "1\n2",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"1\n2", "1", "1", "2", "2"},
			continueOn:       "",
		},
		{
			name: "walk early return after statement list",
			input: []ExprNode{
				&NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
				&NumberLiteral{Value: "24", StartPos: 3, EndPos: 5},
			},
			expectedValue:    "42\n24",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"42\n24"},
			continueOn:       "42\n24",
		},
		{
			name: "walk early return after first statement",
			input: []ExprNode{
				&NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
				&NumberLiteral{Value: "24", StartPos: 3, EndPos: 5},
			},
			expectedValue:    "42\n24",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"42\n24", "42"},
			continueOn:       "42",
		},
		{
			name:             "empty statement list",
			input:            []ExprNode{},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes:    []string{""},
			continueOn:       "",
		},
		{
			name: "statement list with nil statement",
			input: []ExprNode{
				nil,
				&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			},
			expectedValue:    "1",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"1"},
			continueOn:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ast := &StatementList{
				Statements: test.input,
				StartPos:   test.expectedStartPos,
				EndPos:     test.expectedEndPos,
			}

			if ast.Expr() != test.expectedValue {
				t.Fatalf("expected '%s', got '%s'", test.expectedValue, ast.Expr())
			}

			if ast.StartPosition() != test.expectedStartPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedStartPos,
					ast.StartPosition(),
				)
			}

			if ast.EndPosition() != test.expectedEndPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedEndPos,
					ast.EndPosition(),
				)
			}

			WalkUntil(t, ast, test.expectedNodes, test.continueOn)
		})
	}
}
