package ast

import (
	"testing"
)

func TestBoolLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *BoolLiteral
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "true literal",
			input: &BoolLiteral{
				Value:    "true",
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"true"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "false literal",
			input: &BoolLiteral{
				Value:    "false",
				StartPos: 0,
				EndPos:   5,
			},
			expectedNodes:    []string{"false"},
			expectedStartPos: 0,
			expectedEndPos:   5,
			continueOn:       "",
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
