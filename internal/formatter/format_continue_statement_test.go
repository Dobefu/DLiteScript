package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatContinueStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.ContinueStatement
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name:      "continue statement",
			input:     &ast.ContinueStatement{Count: 1, StartPos: 0, EndPos: 1},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "continue\n",
		},
		{
			name:      "continue statement with count 2",
			input:     &ast.ContinueStatement{Count: 2, StartPos: 0, EndPos: 1},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "continue 2\n",
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
