package ast

import (
	"testing"
)

func TestNullLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
	}{
		{
			input: &NullLiteral{
				StartPos: 0,
				EndPos:   4,
			},
			expectedValue:    "null",
			expectedStartPos: 0,
			expectedEndPos:   4,
			expectedNodes:    []string{"null"},
		},
	}

	for _, test := range tests {
		visitedNodes := []string{}

		test.input.Walk(func(node ExprNode) bool {
			visitedNodes = append(visitedNodes, node.Expr())

			return true
		})

		if len(visitedNodes) != len(test.expectedNodes) {
			t.Fatalf(
				"Expected %d visited nodes, got %d (%v)",
				len(test.expectedNodes),
				len(visitedNodes),
				visitedNodes,
			)
		}

		for idx, node := range visitedNodes {
			if node != test.expectedNodes[idx] {
				t.Fatalf("Expected \"%s\", got \"%s\"", test.expectedNodes[idx], node)
			}
		}

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
