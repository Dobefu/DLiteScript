package ast

import (
	"testing"
)

func TestNewlineLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *NewlineLiteral
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "newline literal",
			input: &NewlineLiteral{
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "\n",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"\n"},
			continueOn:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
		})

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
	}
}
