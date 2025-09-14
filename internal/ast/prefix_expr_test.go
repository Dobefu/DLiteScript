package ast

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestPrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "prefix expression with addition",
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
			continueOn: "",
		},
		{
			name: "prefix expression with negation",
			input: &PrefixExpr{
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &NumberLiteral{Value: "5", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   2,
			},
			expectedValue:    "(- 5)",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes: []string{
				"(- 5)",
				"5",
				"5",
			},
			continueOn: "",
		},
		{
			name: "prefix expression with nil operand",
			input: &PrefixExpr{
				Operator: token.Token{
					Atom:      "!",
					TokenType: token.TokenTypeNot,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  nil,
				StartPos: 0,
				EndPos:   1,
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes: []string{
				"",
			},
			continueOn: "",
		},
		{
			name: "walk early return after prefix expression",
			input: &PrefixExpr{
				Operator: token.Token{
					Atom:      "~",
					TokenType: token.TokenTypeNot,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &NumberLiteral{Value: "3", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   2,
			},
			expectedValue:    "(~ 3)",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes: []string{
				"(~ 3)",
			},
			continueOn: "(~ 3)",
		},
		{
			name: "walk early return after operand",
			input: &PrefixExpr{
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &NumberLiteral{Value: "7", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   2,
			},
			expectedValue:    "(+ 7)",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes: []string{
				"(+ 7)",
				"7",
			},
			continueOn: "7",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
