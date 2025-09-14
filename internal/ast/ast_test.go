package ast

import "testing"

func WalkUntil(
	t *testing.T,
	input ExprNode,
	expectedNodes []string,
	continueOn string,
) {
	t.Helper()

	if continueOn == "" {
		visitedNodes := []string{}

		input.Walk(func(node ExprNode) bool {
			if node == nil {
				return false
			}

			visitedNodes = append(visitedNodes, node.Expr())

			return true
		})

		if len(visitedNodes) != len(expectedNodes) {
			t.Fatalf(
				"Expected %d visited nodes, got %d",
				len(expectedNodes),
				len(visitedNodes),
			)
		}

		for idx, node := range visitedNodes {
			if node != expectedNodes[idx] {
				t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
			}
		}

		return
	}

	visitedNodes := []string{}

	input.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return node.Expr() != continueOn
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf(
			"expected %d visited nodes, got %d",
			len(expectedNodes),
			len(visitedNodes),
		)
	}

	for i, expected := range expectedNodes {
		if visitedNodes[i] != expected {
			t.Errorf(
				"expected node %d to be \"%s\", got \"%s\"",
				i,
				expected,
				visitedNodes[i],
			)
		}
	}
}
