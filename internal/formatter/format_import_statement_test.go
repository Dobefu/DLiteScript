package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatImportStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.ImportStatement
		formatter *Formatter
		expected  string
	}{
		{
			name: "imort statement",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "./path/to/file.dl",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 17, Line: 0, Column: 0},
					},
				},
				Namespace: "",
				Alias:     "",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			expected:  "import \"./path/to/file.dl\"\n",
		},
		{
			name: "imort statement with alias",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "./path/to/file.dl",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 17, Line: 0, Column: 0},
					},
				},
				Namespace: "",
				Alias:     "alias",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			expected:  "import \"./path/to/file.dl\" as alias\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatNode(test.input, builder, 0)

			if builder.String() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, builder.String())
			}
		})
	}
}
