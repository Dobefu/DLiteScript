package ast

import (
	"testing"
)

func TestNumberLiteral(t *testing.T) {
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
			name:             "number literal 1 at position 0",
			input:            &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			expectedValue:    "1",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"1"},
			continueOn:       "",
		},
		{
			name:             "number literal 1 at position 1",
			input:            &NumberLiteral{Value: "1", StartPos: 1, EndPos: 2},
			expectedValue:    "1",
			expectedStartPos: 1,
			expectedEndPos:   2,
			expectedNodes:    []string{"1"},
			continueOn:       "",
		},
		{
			name:             "number literal with different value",
			input:            &NumberLiteral{Value: "42", StartPos: 5, EndPos: 7},
			expectedValue:    "42",
			expectedStartPos: 5,
			expectedEndPos:   7,
			expectedNodes:    []string{"42"},
			continueOn:       "",
		},
		{
			name:             "number literal with decimal",
			input:            &NumberLiteral{Value: "3.14", StartPos: 0, EndPos: 4},
			expectedValue:    "3.14",
			expectedStartPos: 0,
			expectedEndPos:   4,
			expectedNodes:    []string{"3.14"},
			continueOn:       "",
		},
		{
			name:             "walk early return after number literal",
			input:            &NumberLiteral{Value: "100", StartPos: 0, EndPos: 3},
			expectedValue:    "100",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"100"},
			continueOn:       "100",
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
