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
				&NumberLiteral{Value: "1", Pos: 0},
			},
			expectedValue: "1",
			expectedPos:   0,
		},
		{
			input: []ExprNode{
				&NumberLiteral{Value: "1", Pos: 0},
				&NumberLiteral{Value: "2", Pos: 2},
			},
			expectedValue: "1\n2",
			expectedPos:   0,
		},
	}

	for _, test := range tests {
		ast := &StatementList{
			Statements: test.input,
			Pos:        0,
		}

		if ast.Expr() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, ast.Expr())
		}

		if ast.Position() != test.expectedPos {
			t.Errorf("expected pos '%d', got '%d'", test.expectedPos, ast.Position())
		}
	}
}
