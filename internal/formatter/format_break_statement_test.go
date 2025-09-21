package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatBreakStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.BreakStatement
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name:      "break statement",
			input:     &ast.BreakStatement{Count: 1, StartPos: 0, EndPos: 1},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "break\n",
		},
		{
			name:      "break statement with count 2",
			input:     &ast.BreakStatement{Count: 2, StartPos: 0, EndPos: 1},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "break 2\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatBreakStatement(test.input, builder, test.depth)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
