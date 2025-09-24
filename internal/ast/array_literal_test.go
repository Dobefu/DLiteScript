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
				Values: []ExprNode{},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
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
					&NumberLiteral{
						Value: "1",
						Range: Range{
							Start: Position{Offset: 0, Line: 0, Column: 0},
							End:   Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&NumberLiteral{
						Value: "1",
						Range: Range{
							Start: Position{Offset: 1, Line: 0, Column: 0},
							End:   Position{Offset: 2, Line: 0, Column: 0},
						},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
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
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
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
					&NumberLiteral{
						Value: "42",
						Range: Range{
							Start: Position{Offset: 0, Line: 0, Column: 0},
							End:   Position{Offset: 2, Line: 0, Column: 0},
						},
					},
					&NumberLiteral{
						Value: "24",
						Range: Range{
							Start: Position{Offset: 2, Line: 0, Column: 0},
							End:   Position{Offset: 4, Line: 0, Column: 0},
						},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 4, Line: 0, Column: 0},
				},
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
					&NumberLiteral{
						Value: "42",
						Range: Range{
							Start: Position{Offset: 0, Line: 0, Column: 0},
							End:   Position{Offset: 2, Line: 0, Column: 0},
						},
					},
					&NumberLiteral{
						Value: "24",
						Range: Range{
							Start: Position{Offset: 2, Line: 0, Column: 0},
							End:   Position{Offset: 4, Line: 0, Column: 0},
						},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 4, Line: 0, Column: 0},
				},
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

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedStartPos,
					test.input.GetRange().Start.Offset,
				)
			}

			if test.input.GetRange().End.Offset != test.expectedEndPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedEndPos,
					test.input.GetRange().End.Offset,
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
