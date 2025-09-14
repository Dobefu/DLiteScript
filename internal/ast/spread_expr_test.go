package ast

import "testing"

func TestSpreadExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *SpreadExpr
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "spread expression",
			input: &SpreadExpr{
				Expression: &NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expectedValue:    "...1",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"...1", "1", "1"},
			continueOn:       "",
		},
		{
			name: "walk early return after spread node",
			input: &SpreadExpr{
				Expression: &NumberLiteral{
					Value:    "42",
					StartPos: 0,
					EndPos:   4,
				},
				StartPos: 0,
				EndPos:   4,
			},
			expectedValue:    "...42",
			expectedStartPos: 0,
			expectedEndPos:   4,
			expectedNodes:    []string{"...42"},
			continueOn:       "...42",
		},
		{
			name: "walk early return after expression",
			input: &SpreadExpr{
				Expression: &NumberLiteral{
					Value:    "42",
					StartPos: 0,
					EndPos:   4,
				},
				StartPos: 0,
				EndPos:   4,
			},
			expectedValue:    "...42",
			expectedStartPos: 0,
			expectedEndPos:   4,
			expectedNodes:    []string{"...42", "42"},
			continueOn:       "42",
		},
		{
			name: "spread expression with nil",
			input: &SpreadExpr{
				Expression: nil,
				StartPos:   0,
				EndPos:     3,
			},
			expectedValue:    "...",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"..."},
			continueOn:       "",
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
