package ast

import "testing"

func TestBlockStatement(t *testing.T) {
	t.Parallel()

	block := &BlockStatement{
		Statements: []ExprNode{
			&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
		},
		StartPos: 0,
		EndPos:   1,
	}

	visitedNodes := []string{}
	expectedNodes := []string{"(1)", "1", "1"}

	block.Walk(func(node ExprNode) bool {
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

	if block.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", block.StartPosition())
	}

	if block.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", block.EndPosition())
	}
}
