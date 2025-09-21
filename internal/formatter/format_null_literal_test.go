package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatNullLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.NullLiteral
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name:      "null literal",
			input:     &ast.NullLiteral{StartPos: 0, EndPos: 0},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "null\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatNullLiteral(test.input, builder, test.depth)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
