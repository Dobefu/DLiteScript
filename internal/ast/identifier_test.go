package ast

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
	}{
		{
			input:            &Identifier{Value: "PI", StartPos: 0, EndPos: 1},
			expectedValue:    "PI",
			expectedStartPos: 0,
			expectedEndPos:   1,
		},
		{
			input:            &Identifier{Value: "PI", StartPos: 1, EndPos: 2},
			expectedValue:    "PI",
			expectedStartPos: 1,
			expectedEndPos:   2,
		},
	}

	for _, test := range tests {
		if test.input.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
		}

		if test.input.StartPosition() != test.expectedStartPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedStartPos, test.input.StartPosition())
		}

		if test.input.EndPosition() != test.expectedEndPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedEndPos, test.input.EndPosition())
		}
	}
}
