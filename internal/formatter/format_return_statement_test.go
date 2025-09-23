package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatReturnStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.ReturnStatement
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "return statement",
			input: &ast.ReturnStatement{
				Values:    []ast.ExprNode{},
				NumValues: 0,
				StartPos:  0,
				EndPos:    0,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "return\n",
		},
		{
			name: "return statement with one value",
			input: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 0},
				},
				NumValues: 1,
				StartPos:  0,
				EndPos:    0,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "return 1\n",
		},
		{
			name: "return statement with two values",
			input: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 0},
					&ast.StringLiteral{Value: "test", StartPos: 0, EndPos: 0},
				},
				NumValues: 2,
				StartPos:  0,
				EndPos:    0,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "return 1, \"test\"\n",
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
