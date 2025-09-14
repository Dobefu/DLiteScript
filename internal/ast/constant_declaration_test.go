package ast

import (
	"testing"
)

func TestConstantDeclaration(t *testing.T) {
	t.Parallel()

	declaration := &ConstantDeclaration{
		Name: "x",
		Type: "int",
		Value: &NumberLiteral{
			Value:    "1",
			StartPos: 0,
			EndPos:   1,
		},
		StartPos: 0,
		EndPos:   1,
	}

	expectedNodes := []string{"const x int = 1", "1", "1"}
	visitedNodes := []string{}

	declaration.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf(
			"Expected %d visited node, got %d",
			len(expectedNodes),
			len(visitedNodes),
		)
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}

	if declaration.StartPosition() != 0 {
		t.Fatalf(
			"Expected start position to be 0, got %d",
			declaration.StartPosition(),
		)
	}

	if declaration.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", declaration.EndPosition())
	}
}
