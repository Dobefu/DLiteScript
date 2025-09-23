package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatIndexExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.IndexExpr
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "index expression",
			input: &ast.IndexExpr{
				Array:    &ast.Identifier{Value: "array", StartPos: 0, EndPos: 5},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "array[0]\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatIndexExpr(test.input, builder, test.depth)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
