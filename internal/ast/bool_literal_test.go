package ast

import (
	"testing"
)

func TestBoolLiteral(t *testing.T) {
	t.Parallel()

	literal := &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1}
	expectedNodes := []string{"true"}

	visitedNodes := []string{}
	literal.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf(
			"Expected %d visited node, got %d (%v)",
			len(expectedNodes),
			len(visitedNodes),
			visitedNodes,
		)
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}

	if literal.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", literal.StartPosition())
	}

	if literal.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", literal.EndPosition())
	}
}
