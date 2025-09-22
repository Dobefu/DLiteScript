package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatBoolLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.BoolLiteral
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name:      "true literal",
			input:     &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 0},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "true\n",
		},
		{
			name:      "false literal",
			input:     &ast.BoolLiteral{Value: "false", StartPos: 0, EndPos: 0},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "false\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatBoolLiteral(test.input, builder, test.depth)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
