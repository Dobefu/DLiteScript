package ast

import (
	"testing"
)

func TestStatementList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input            []ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
	}{
		{
			input: []ExprNode{
				&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			},
			expectedValue:    "1",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"1", "1", "1"},
		},
		{
			input: []ExprNode{
				&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				&NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
			},
			expectedValue:    "1\n2",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"1\n2", "1", "1", "2", "2"},
		},
	}

	for _, test := range tests {
		visitedNodes := []string{}

		ast := &StatementList{
			Statements: test.input,
			StartPos:   test.expectedStartPos,
			EndPos:     test.expectedEndPos,
		}

		ast.Walk(func(node ExprNode) bool {
			visitedNodes = append(visitedNodes, node.Expr())

			return true
		})

		if len(visitedNodes) != len(test.expectedNodes) {
			t.Fatalf(
				"Expected %d visited nodes, got %d (%v)",
				len(test.expectedNodes),
				len(visitedNodes),
				visitedNodes,
			)
		}

		for idx, node := range visitedNodes {
			if node != test.expectedNodes[idx] {
				t.Fatalf("Expected \"%s\", got \"%s\"", test.expectedNodes[idx], node)
			}
		}

		if ast.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, ast.Expr())
		}

		if ast.StartPosition() != test.expectedStartPos {
			t.Errorf(
				"expected pos '%d', got '%d'",
				test.expectedStartPos,
				ast.StartPosition(),
			)
		}

		if ast.EndPosition() != test.expectedEndPos {
			t.Errorf(
				"expected pos '%d', got '%d'",
				test.expectedEndPos,
				ast.EndPosition(),
			)
		}
	}
}
