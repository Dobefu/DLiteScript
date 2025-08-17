package ast

import (
	"testing"
)

func TestStatementList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         []ExprNode
		expectedValue string
		expectedPos   int
	}{
		{
			input: []ExprNode{
				&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			},
			expectedValue: "1",
			expectedPos:   0,
		},
		{
			input: []ExprNode{
				&NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				&NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
			},
			expectedValue: "1\n2",
			expectedPos:   0,
		},
	}

	for _, test := range tests {
		ast := &StatementList{
			Statements: test.input,
			StartPos:   0,
			EndPos:     0,
		}

		if ast.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, ast.Expr())
		}

		if ast.StartPosition() != test.expectedPos {
			t.Errorf(
				"expected pos '%d', got '%d'",
				test.expectedPos,
				ast.StartPosition(),
			)
		}
	}
}
