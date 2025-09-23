package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatIndexAssignmentStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.IndexAssignmentStatement
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "index assignment statement",
			input: &ast.IndexAssignmentStatement{
				Array:    &ast.Identifier{Value: "array", StartPos: 0, EndPos: 5},
				Index:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "array[1] = 1\n",
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
