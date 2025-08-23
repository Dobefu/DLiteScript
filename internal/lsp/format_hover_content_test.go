package lsp

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatHoverContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		node        ast.ExprNode
		isDebugMode bool
		expected    string
	}{
		{
			name: "null literal node",
			node: &ast.NullLiteral{
				StartPos: 0,
				EndPos:   0,
			},
			isDebugMode: false,
			expected:    "",
		},
		{
			name: "null literal node in debug mode",
			node: &ast.NullLiteral{
				StartPos: 0,
				EndPos:   0,
			},
			isDebugMode: true,
			expected:    "**🔴 Debug Mode** Unknown Node: *ast.NullLiteral\n\n```dlitescript\nnull\n```",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			hoverContent := formatHoverContent(test.node, test.isDebugMode)

			if hoverContent != test.expected {
				t.Errorf("expected %s, got %s", test.expected, hoverContent)
			}
		})
	}
}
