package ast

import (
	"testing"
)

func TestAnyLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *AnyLiteral
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "any literal",
			input: &AnyLiteral{
				Value: &NumberLiteral{
					Value: "1",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"any", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after first node",
			input: &AnyLiteral{
				Value: &NumberLiteral{
					Value: "42",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"any"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "any",
		},
		{
			name: "walk early return after value node",
			input: &AnyLiteral{
				Value: &NumberLiteral{
					Value: "42",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"any", "42"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "42",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != "any" {
				t.Errorf("expected \"any\", got \"%s\"", test.input.Expr())
			}

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
