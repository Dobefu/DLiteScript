package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatAnyLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.AnyLiteral
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name:      "any literal",
			input:     &ast.AnyLiteral{Value: nil, StartPos: 0, EndPos: 1},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "any\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatNode(test.input, builder, test.depth)

			if builder.String() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, builder.String())
			}
		})
	}
}
