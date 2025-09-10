package ast

import "testing"

func TestIndexExpr(t *testing.T) {
	t.Parallel()

	expr := &IndexExpr{
		Array:    &Identifier{Value: "array", StartPos: 0, EndPos: 5},
		Index:    &NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
		StartPos: 0,
		EndPos:   5,
	}

	expectedNodes := []string{"array[0]", "array", "0"}
	expectedValue := "array[0]"
	visitedNodes := []string{}

	expr.Walk(func(node ExprNode) bool {
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

	if expr.Expr() != expectedValue {
		t.Fatalf("Expected \"%s\", got \"%s\"", expectedValue, expr.Expr())
	}

	if expr.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", expr.StartPosition())
	}

	if expr.EndPosition() != 5 {
		t.Fatalf("Expected end position to be 5, got %d", expr.EndPosition())
	}
}
