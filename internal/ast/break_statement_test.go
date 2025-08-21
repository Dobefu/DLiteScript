package ast

import (
	"testing"
)

func TestBreakStatement(t *testing.T) {
	t.Parallel()

	statement := &BreakStatement{Count: 1, StartPos: 0, EndPos: 1}
	expectedNodes := []string{"break"}
	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != 1 {
		t.Fatalf("Expected 1 visited node, got %d", len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}

	if statement.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", statement.StartPosition())
	}

	if statement.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", statement.EndPosition())
	}
}

func TestBreakStatementWithCount(t *testing.T) {
	t.Parallel()

	statement := &BreakStatement{Count: 2, StartPos: 0, EndPos: 1}
	expectedNodes := []string{"break 2"}
	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf("Expected %d visited node, got %d", len(expectedNodes), len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}

	if statement.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", statement.StartPosition())
	}

	if statement.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", statement.EndPosition())
	}
}
