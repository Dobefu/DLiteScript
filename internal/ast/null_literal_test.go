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
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 4, Line: 0, Column: 0},
				},
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
				Range: Range{
					Start: Position{Offset: 5, Line: 0, Column: 0},
					End:   Position{Offset: 9, Line: 0, Column: 0},
				},
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
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 4, Line: 0, Column: 0},
				},
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

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Errorf(
					"expected pos %d, got %d",
					test.expectedStartPos,
					test.input.GetRange().Start.Offset,
				)
			}

			if test.input.GetRange().End.Offset != test.expectedEndPos {
				t.Errorf(
					"expected pos %d, got %d",
					test.expectedEndPos,
					test.input.GetRange().End.Offset,
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
