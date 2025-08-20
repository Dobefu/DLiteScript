package ast

import "testing"

func TestBoolLiteral(t *testing.T) {
	t.Parallel()

	literal := &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1}

	visitedNodes := []string{}
	literal.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != 1 {
		t.Fatalf("Expected 1 visited node, got %d", len(visitedNodes))
	}

	if visitedNodes[0] != "true" {
		t.Fatalf("Expected \"true\", got \"%s\"", visitedNodes[0])
	}

	if literal.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", literal.StartPosition())
	}

	if literal.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", literal.EndPosition())
	}
}
