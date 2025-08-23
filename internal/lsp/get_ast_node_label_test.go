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
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{Value: "test", StartPos: 6, EndPos: 10},
				},
				StartPos: 0,
				EndPos:   10,
			},
			expected: "Function Call",
		},
		{
			name:     "identifier",
			node:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
			expected: "Identifier",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			label := getAstNodeLabel(test.node, false)

			if label != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, label)
			}
		})
	}
}
