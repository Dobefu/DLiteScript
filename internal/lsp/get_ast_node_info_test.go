package lsp

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestGetAstNodeLabel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		node     ast.ExprNode
		expected string
	}{
		{
			name: "function call",
			node: &ast.FunctionCall{
				Namespace:    "math",
				FunctionName: "sqrt",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "4",
						Range: ast.Range{
							Start: ast.Position{Offset: 6, Line: 0, Column: 0},
							End:   ast.Position{Offset: 10, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 10, Line: 0, Column: 0},
				},
			},
			expected: "Function Call",
		},
		{
			name: "bogus namespace",
			node: &ast.FunctionCall{
				Namespace:    "bogus",
				FunctionName: "sqrt",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "4",
						Range: ast.Range{
							Start: ast.Position{Offset: 6, Line: 0, Column: 0},
							End:   ast.Position{Offset: 10, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 10, Line: 0, Column: 0},
				},
			},
			expected: "Function Call",
		},
		{
			name: "bogus function name",
			node: &ast.FunctionCall{
				Namespace:    "math",
				FunctionName: "bogus",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "4",
						Range: ast.Range{
							Start: ast.Position{Offset: 6, Line: 0, Column: 0},
							End:   ast.Position{Offset: 10, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 10, Line: 0, Column: 0},
				},
			},
			expected: "Function Call",
		},
		{
			name: "identifier",
			node: &ast.Identifier{
				Value: "x",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "Identifier",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			label := getAstNodeInfo(test.node, false)

			if label.Label != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, label)
			}
		})
	}
}
