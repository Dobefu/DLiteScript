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
			name: "identifier PI at position 0",
			input: &Identifier{
				Value: "PI",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "PI",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"PI"},
			continueOn:       "",
		},
		{
			name: "identifier PI at position 1",
			input: &Identifier{
				Value: "PI",
				Range: Range{
					Start: Position{Offset: 1, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
			},
			expectedValue:    "PI",
			expectedStartPos: 1,
			expectedEndPos:   2,
			expectedNodes:    []string{"PI"},
			continueOn:       "",
		},
		{
			name: "identifier with different value",
			input: &Identifier{
				Value: "count",
				Range: Range{
					Start: Position{Offset: 5, Line: 0, Column: 0},
					End:   Position{Offset: 10, Line: 0, Column: 0},
				},
			},
			expectedValue:    "count",
			expectedStartPos: 5,
			expectedEndPos:   10,
			expectedNodes:    []string{"count"},
			continueOn:       "",
		},
		{
			name: "walk early return after identifier",
			input: &Identifier{
				Value: "x",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
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
				t.Errorf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.GetRange().Start.Offset,
				)
			}

			if test.input.GetRange().End.Offset != test.expectedEndPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.GetRange().End.Offset,
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
