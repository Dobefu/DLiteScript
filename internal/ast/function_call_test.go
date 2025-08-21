package ast

import (
	"testing"
)

func TestFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
	}{
		{
			input: &FunctionCall{
				FunctionName: "abs",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedValue:    "abs(1)",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"abs(1)", "1", "1"},
		},
		{
			input: &FunctionCall{
				FunctionName: "max",
				Arguments: []ExprNode{
					&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					&NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedValue:    "max(1, 2)",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"max(1, 2)", "1", "1", "2", "2"},
		},
	}

	for _, test := range tests {
		visitedNodes := []string{}

		test.input.Walk(func(node ExprNode) bool {
			visitedNodes = append(visitedNodes, node.Expr())

			return true
		})

		if len(visitedNodes) != len(test.expectedNodes) {
			t.Fatalf("Expected %d visited nodes, got %d", len(test.expectedNodes), len(visitedNodes))
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
