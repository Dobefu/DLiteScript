package ast

import (
	"testing"
)

func TestIndexAssignmentStatement(t *testing.T) {
	t.Parallel()

	statement := &IndexAssignmentStatement{
		Array:    &Identifier{Value: "array", StartPos: 0, EndPos: 5},
		Index:    &NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
		Right:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
		StartPos: 0,
		EndPos:   5,
	}

	expectedNodes := []string{"array[0] = 1", "array", "array", "0", "0", "1", "1"}
	expectedValue := "array[0] = 1"
	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
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

	if statement.Expr() != expectedValue {
		t.Fatalf("Expected \"%s\", got \"%s\"", expectedValue, statement.Expr())
	}

	if statement.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", statement.StartPosition())
	}

	if statement.EndPosition() != 5 {
		t.Fatalf("Expected end position to be 5, got %d", statement.EndPosition())
	}
}
