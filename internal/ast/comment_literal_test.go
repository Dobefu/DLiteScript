package ast

import (
	"testing"
)

func TestCommentLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *CommentLiteral
		expectedValue    string
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "comment literal",
			input: &CommentLiteral{
				Value:    "Comment",
				StartPos: 0,
				EndPos:   1,
			},
			expectedValue:    "Comment",
			expectedNodes:    []string{},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Errorf("expected '%s', got '%s'", test.expectedNodes, test.input.Expr())
			}

			if test.input.StartPosition() != test.expectedStartPos {
				t.Errorf("expected pos '%d', got '%d'", test.expectedStartPos, test.input.StartPosition())
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Errorf("expected pos '%d', got '%d'", test.expectedEndPos, test.input.EndPosition())
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
