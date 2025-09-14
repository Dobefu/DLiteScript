package ast

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestPrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
	}{
		{
			input: &PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand: &BinaryExpr{
					Left:  &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					Right: &NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					Operator: token.Token{
						Atom:      "+",
						TokenType: token.TokenTypeOperationAdd,
						StartPos:  0,
						EndPos:    0,
					},
					StartPos: 0,
					EndPos:   0,
				},
				StartPos: 0,
				EndPos:   0,
			},
			expectedValue:    "(+ (1 + 1))",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes: []string{
				"(+ (1 + 1))",
				"(1 + 1)",
				"(1 + 1)",
				"1",
				"1",
				"1",
				"1",
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
			t.Fatalf(
				"Expected %d visited nodes, got %d",
				len(test.expectedNodes),
				len(visitedNodes),
			)
		}

		for idx, node := range visitedNodes {
			if node != test.expectedNodes[idx] {
				t.Fatalf("Expected \"%s\", got \"%s\"", test.expectedNodes[idx], node)
			}
		}

		if test.input.Expr() != test.expectedValue {
			t.Errorf(
				"expected '%s', got '%s'",
				test.expectedValue,
				test.input.Expr(),
			)
		}

		if test.input.StartPosition() != test.expectedStartPos {
			t.Errorf(
				"expected pos '%d', got '%d'",
				test.expectedStartPos,
				test.input.StartPosition(),
			)
		}

		if test.input.EndPosition() != test.expectedEndPos {
			t.Errorf(
				"expected pos '%d', got '%d'",
				test.expectedEndPos,
				test.input.EndPosition(),
			)
		}
	}
}
