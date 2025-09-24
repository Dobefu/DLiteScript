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
					&NumberLiteral{
						Value: "1",
						Range: Range{
							Start: Position{Offset: 0, Line: 0, Column: 0},
							End:   Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
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
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
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
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
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
					&NumberLiteral{
						Value: "1",
						Range: Range{
							Start: Position{Offset: 0, Line: 0, Column: 0},
							End:   Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&NumberLiteral{
						Value: "2",
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
			expectedNodes:    []string{"(1 2)", "1", "1", "2", "2"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "",
		},
		{
			name: "walk early return after block node",
			input: &BlockStatement{
				Statements: []ExprNode{
					&NumberLiteral{
						Value: "42",
						Range: Range{
							Start: Position{Offset: 0, Line: 0, Column: 0},
							End:   Position{Offset: 2, Line: 0, Column: 0},
						},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
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
			expectedNodes:    []string{"(42 24)", "42"},
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
