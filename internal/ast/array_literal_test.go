package ast

import (
	"testing"
)

func TestArrayLiteral(t *testing.T) {
	t.Parallel()

	array := &ArrayLiteral{
		Values: []ExprNode{
			&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			&NumberLiteral{Value: "1", StartPos: 1, EndPos: 2},
		},
		StartPos: 0,
		EndPos:   2,
	}

	expectedNodes := []string{"[1, 1]", "1", "1"}
	expectedValue := "[1, 1]"
	visitedNodes := []string{}

	array.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf("Expected %d visited nodes, got %d", len(expectedNodes), len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}

	if array.Expr() != expectedValue {
		t.Fatalf("Expected \"%s\", got \"%s\"", expectedValue, array.Expr())
	}

	if array.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", array.StartPosition())
	}

	if array.EndPosition() != 2 {
		t.Fatalf("Expected end position to be 2, got %d", array.EndPosition())
	}
}
