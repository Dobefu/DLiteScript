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
				Value:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{"any", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after first node",
			input: &AnyLiteral{
				Value:    &NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
				StartPos: 0,
				EndPos:   2,
			},
			expectedNodes:    []string{"any"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "any",
		},
		{
			name: "walk early return after value node",
			input: &AnyLiteral{
				Value:    &NumberLiteral{Value: "42", StartPos: 0, EndPos: 2},
				StartPos: 0,
				EndPos:   2,
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
