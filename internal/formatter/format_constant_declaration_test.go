package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatConstantDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.ConstantDeclaration
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "constant declaration",
			input: &ast.ConstantDeclaration{
				Name: "x",
				Type: "int",
				Value: &ast.NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
				StartPos: 0,
				EndPos:   1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "const x int = 1\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatConstantDeclaration(test.input, builder, test.depth)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
