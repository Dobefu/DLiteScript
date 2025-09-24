package ast

import (
	"testing"
)

func TestContinueStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *ContinueStatement
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "continue statement with count 1",
			input: &ContinueStatement{
				Count: 1,
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"continue"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "continue statement with count 2",
			input: &ContinueStatement{
				Count: 2,
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"continue 2"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "continue statement with count 0",
			input: &ContinueStatement{
				Count: 0,
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"continue 0"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after continue statement",
			input: &ContinueStatement{
				Count: 3,
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"continue 3"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "continue 3",
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
