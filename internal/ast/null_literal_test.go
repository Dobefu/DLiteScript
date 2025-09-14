package ast

import (
	"testing"
)

func TestNullLiteral(t *testing.T) {
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
			name: "null literal",
			input: &NullLiteral{
				StartPos: 0,
				EndPos:   4,
			},
			expectedValue:    "null",
			expectedStartPos: 0,
			expectedEndPos:   4,
			expectedNodes:    []string{"null"},
			continueOn:       "",
		},
		{
			name: "null literal at different position",
			input: &NullLiteral{
				StartPos: 5,
				EndPos:   9,
			},
			expectedValue:    "null",
			expectedStartPos: 5,
			expectedEndPos:   9,
			expectedNodes:    []string{"null"},
			continueOn:       "",
		},
		{
			name: "walk early return after null literal",
			input: &NullLiteral{
				StartPos: 0,
				EndPos:   4,
			},
			expectedValue:    "null",
			expectedStartPos: 0,
			expectedEndPos:   4,
			expectedNodes:    []string{"null"},
			continueOn:       "null",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.StartPosition() != test.expectedStartPos {
				t.Errorf(
					"expected pos %d, got %d",
					test.expectedStartPos,
					test.input.StartPosition(),
				)
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Errorf(
					"expected pos %d, got %d",
					test.expectedEndPos,
					test.input.EndPosition(),
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
