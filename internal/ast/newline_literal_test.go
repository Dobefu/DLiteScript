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
			name:             "newline literal",
			input:            &NewlineLiteral{StartPos: 0, EndPos: 1},
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
			t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
		}

		if test.input.StartPosition() != test.expectedStartPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedStartPos, test.input.StartPosition())
		}

		if test.input.EndPosition() != test.expectedEndPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedEndPos, test.input.EndPosition())
		}

		WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
	}
}
