package ast

import "testing"

func TestAssignmentStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *AssignmentStatement
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "assignment statement",
			input: &AssignmentStatement{
				Left: &Identifier{
					Value:    "x",
					StartPos: 0,
					EndPos:   1,
				},
				Right: &NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"x = 1", "x", "x", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "assignment statement with nil left",
			input: &AssignmentStatement{
				Left: nil,
				Right: &NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after assignment node",
			input: &AssignmentStatement{
				Left: &Identifier{
					Value:    "y",
					StartPos: 0,
					EndPos:   1,
				},
				Right: &NumberLiteral{
					Value:    "42",
					StartPos: 0,
					EndPos:   2,
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedNodes:    []string{"y = 42"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "y = 42",
		},
		{
			name: "walk early return after left identifier",
			input: &AssignmentStatement{
				Left: &Identifier{
					Value:    "y",
					StartPos: 0,
					EndPos:   1,
				},
				Right: &NumberLiteral{
					Value:    "42",
					StartPos: 0,
					EndPos:   2,
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedNodes:    []string{"y = 42", "y"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "y",
		},
		{
			name: "walk early return after right value",
			input: &AssignmentStatement{
				Left: &Identifier{
					Value:    "y",
					StartPos: 0,
					EndPos:   1,
				},
				Right: &NumberLiteral{
					Value:    "42",
					StartPos: 0,
					EndPos:   2,
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedNodes:    []string{"y = 42", "y", "y", "42"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "42",
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
