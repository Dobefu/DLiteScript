package ast

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
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
			name:             "identifier PI at position 0",
			input:            &Identifier{Value: "PI", StartPos: 0, EndPos: 1},
			expectedValue:    "PI",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"PI"},
			continueOn:       "",
		},
		{
			name:             "identifier PI at position 1",
			input:            &Identifier{Value: "PI", StartPos: 1, EndPos: 2},
			expectedValue:    "PI",
			expectedStartPos: 1,
			expectedEndPos:   2,
			expectedNodes:    []string{"PI"},
			continueOn:       "",
		},
		{
			name:             "identifier with different value",
			input:            &Identifier{Value: "count", StartPos: 5, EndPos: 10},
			expectedValue:    "count",
			expectedStartPos: 5,
			expectedEndPos:   10,
			expectedNodes:    []string{"count"},
			continueOn:       "",
		},
		{
			name:             "walk early return after identifier",
			input:            &Identifier{Value: "x", StartPos: 0, EndPos: 1},
			expectedValue:    "x",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"x"},
			continueOn:       "x",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
			}

			if test.input.StartPosition() != test.expectedStartPos {
				t.Errorf("expected pos '%d', got '%d'", test.expectedStartPos, test.input.StartPosition())
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Errorf("expected pos '%d', got '%d'", test.expectedEndPos, test.input.EndPosition())
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
