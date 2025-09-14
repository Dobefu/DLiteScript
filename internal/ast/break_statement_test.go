package ast

import (
	"testing"
)

func TestBreakStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *BreakStatement
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "break statement with count 1",
			input: &BreakStatement{
				Count:    1,
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"break"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "break statement with count 2",
			input: &BreakStatement{
				Count:    2,
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"break 2"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "break statement with count 0",
			input: &BreakStatement{
				Count:    0,
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"break 0"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after break statement",
			input: &BreakStatement{
				Count:    3,
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"break 3"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "break 3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
