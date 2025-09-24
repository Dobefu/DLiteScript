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
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			isDebugMode: false,
			expected:    "",
		},
		{
			name: "null literal node in debug mode",
			node: &ast.NullLiteral{
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			isDebugMode: true,
			expected:    "**🔴 Debug Mode** | **Unknown Node**\n\n---\n\nUnknown Node: *ast.NullLiteral",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			hoverContent := formatHoverContent(test.node, test.isDebugMode)

			if hoverContent != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, hoverContent)
			}
		})
	}
}
