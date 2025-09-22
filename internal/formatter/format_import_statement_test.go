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
					Value:    "./path/to/file.dl",
					StartPos: 0,
					EndPos:   17,
				},
				Namespace: "",
				Alias:     "",
				StartPos:  0,
				EndPos:    1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			expected:  "import \"./path/to/file.dl\"\n",
		},
		{
			name: "imort statement with alias",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value:    "./path/to/file.dl",
					StartPos: 0,
					EndPos:   17,
				},
				Namespace: "",
				Alias:     "alias",
				StartPos:  0,
				EndPos:    1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			expected:  "import \"./path/to/file.dl\" as alias\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatImportStatement(test.input, builder, 0)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
