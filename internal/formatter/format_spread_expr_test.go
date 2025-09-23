package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatSpreadExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.SpreadExpr
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "spread expression",
			input: &ast.SpreadExpr{
				Expression: &ast.NumberLiteral{Value: "1.1", StartPos: 0, EndPos: 3},
				StartPos:   0,
				EndPos:     3,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "...1.1\n",
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
