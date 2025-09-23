package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatStatementListExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.StatementList
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "statement list",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{
						Value:    "1",
						StartPos: 0,
						EndPos:   0,
					},
				},
				StartPos: 0,
				EndPos:   0,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "1\n",
		},
		{
			name: "empty statement list",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{},
				StartPos:   0,
				EndPos:     0,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "",
		},
		{
			name: "statement list with nil statement",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{nil},
				StartPos:   0,
				EndPos:     0,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "",
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
