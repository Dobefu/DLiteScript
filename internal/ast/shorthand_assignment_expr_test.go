package ast

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestShorthandAssignmentExpr(t *testing.T) {
	t.Parallel()

	assignment := &ShorthandAssignmentExpr{
		Left: &Identifier{
			Value:    "x",
			StartPos: 0,
			EndPos:   1,
		},
		Right: &NumberLiteral{
			Value:    "1",
			StartPos: 0,
			EndPos:   1,
		},
		Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
		StartPos: 0,
		EndPos:   1,
	}

	visitedNodes := []string{}
	expectedNodes := []string{"x += 1", "x", "x", "1", "1"}

	assignment.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != 5 {
		t.Fatalf("Expected 5 visited nodes, got %d", len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}

	if assignment.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", assignment.StartPosition())
	}

	if assignment.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", assignment.EndPosition())
	}
}
