package parser

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseBoolLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *token.Token
		expected ast.ExprNode
	}{
		{
			name:  "true",
			input: token.NewToken("true", token.TokenTypeBool, 0, 0),
			expected: &ast.BoolLiteral{
				Value:    "true",
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			name:  "false",
			input: token.NewToken("false", token.TokenTypeBool, 0, 0),
			expected: &ast.BoolLiteral{
				Value:    "false",
				StartPos: 0,
				EndPos:   0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser([]*token.Token{test.input})
			expr, err := parser.parseBoolLiteral(test.input)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if expr.Expr() != test.expected.Expr() {
				t.Fatalf("expected %s, got %s", test.expected.Expr(), expr.Expr())
			}
		})
	}
}
