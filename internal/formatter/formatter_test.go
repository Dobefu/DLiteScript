package formatter

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestFormatter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "formatter",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.BinaryExpr{
						Left: &ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
						Right: &ast.NumberLiteral{
							Value: "2",
							Range: ast.Range{
								Start: ast.Position{Offset: 1, Line: 0, Column: 0},
								End:   ast.Position{Offset: 2, Line: 0, Column: 0},
							},
						},
						Operator: token.Token{
							Atom:      "+",
							TokenType: token.TokenTypeOperationAdd,
							StartPos:  0,
							EndPos:    1,
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "{\n  1 + 2\n}\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			formatter := New()
			result := formatter.Format(test.input)

			if result != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}
