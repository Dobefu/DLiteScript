package parser

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseNullLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *token.Token
		expected ast.ExprNode
	}{
		{
			input: token.NewToken("null", token.TokenTypeNull, 0, 0),
			expected: &ast.NullLiteral{
				StartPos: 0,
				EndPos:   0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input.Atom, func(t *testing.T) {
			t.Parallel()

			parser := NewParser([]*token.Token{test.input})
			expr, err := parser.parseNullLiteral()

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if expr.Expr() != test.expected.Expr() {
				t.Fatalf("expected %s, got %s", test.expected.Expr(), expr.Expr())
			}
		})
	}
}
