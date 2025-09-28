package linter

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestLinter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "empty",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			linter := New(io.Discard)
			linter.Lint(test.input)

			if linter.HasIssues() {
				t.Fatalf(
					"expected no issues, got %d issues",
					len(linter.reporter.GetIssues()),
				)
			}

			linter.PrintIssues("test.dl")
		})
	}
}
