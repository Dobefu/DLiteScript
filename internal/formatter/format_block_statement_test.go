package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatBlockStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.BlockStatement
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "empty block statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{},
				StartPos:   0,
				EndPos:     1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "{}\n",
		},
		{
			name: "block statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "{\n  1\n}\n",
		},
		{
			name: "block statement with nil statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{nil},
				StartPos:   0,
				EndPos:     1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "{\n}\n",
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
