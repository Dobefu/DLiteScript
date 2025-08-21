package ast

import (
	"testing"
)

func TestIfStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
	}{
		{
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3}},
					StartPos:   2,
					EndPos:     3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			expectedValue:    "if true { (1) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if true { (1) }",
				"true",
				"(1)",
				"(1)",
				"1",
				"1",
			},
		},
		{
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3}},
					StartPos:   2,
					EndPos:     3,
				},
				StartPos: 0,
				EndPos:   5,
				ElseBlock: &BlockStatement{
					Statements: []ExprNode{&NumberLiteral{Value: "2", StartPos: 4, EndPos: 5}},
					StartPos:   4,
					EndPos:     5,
				},
			},
			expectedValue:    "if true { (1) } else { (2) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if true { (1) } else { (2) }",
				"true",
				"(1)",
				"(1)",
				"1",
				"1",
				"(2)",
				"(2)",
				"2",
				"2",
			},
		},
	}

	for _, test := range tests {
		visitedNodes := []string{}

		test.input.Walk(func(node ExprNode) bool {
			visitedNodes = append(visitedNodes, node.Expr())

			return true
		})

		if len(visitedNodes) != len(test.expectedNodes) {
			t.Fatalf("Expected %d visited nodes, got %d", len(test.expectedNodes), len(visitedNodes))
		}

		for idx, node := range visitedNodes {
			if node != test.expectedNodes[idx] {
				t.Fatalf("Expected \"%s\", got \"%s\"", test.expectedNodes[idx], node)
			}
		}

		if test.input.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
		}

		if test.input.StartPosition() != test.expectedStartPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedStartPos, test.input.StartPosition())
		}

		if test.input.EndPosition() != test.expectedEndPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedEndPos, test.input.EndPosition())
		}
	}
}
