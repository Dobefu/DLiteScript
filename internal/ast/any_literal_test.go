package ast

import (
	"testing"
)

func TestAnyLiteral(t *testing.T) {
	t.Parallel()

	anyLiteral := &AnyLiteral{
		Value:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
		StartPos: 0,
		EndPos:   1,
	}

	if anyLiteral.Expr() != "any" {
		t.Errorf("expected 'any', got '%s'", anyLiteral.Expr())
	}

	if anyLiteral.StartPosition() != 0 {
		t.Errorf("expected 0, got '%d'", anyLiteral.StartPosition())
	}

	if anyLiteral.EndPosition() != 1 {
		t.Errorf("expected 1, got '%d'", anyLiteral.EndPosition())
	}

	visitedNodes := []string{}
	expectedNodes := []string{"any", "1", "1"}

	anyLiteral.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf(
			"expected %d visited node, got %d",
			len(expectedNodes),
			len(visitedNodes),
		)
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("expected '%s', got '%s'", expectedNodes[idx], node)
		}
	}

	if visitedNodes[0] != "any" {
		t.Errorf("expected '%s', got '%s'", expectedNodes[0], visitedNodes[0])
	}
}
