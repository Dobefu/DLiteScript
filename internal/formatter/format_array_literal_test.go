package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatArrayLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.ArrayLiteral
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "empty array literal",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "[]\n",
		},
		{
			name: "array literal with nil element",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{nil},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "[]\n",
		},
		{
			name: "array literal with one element",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "[1]\n",
		},
		{
			name: "array literal with two elements",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.StringLiteral{
						Value: "test",
						Range: ast.Range{
							Start: ast.Position{Offset: 1, Line: 0, Column: 0},
							End:   ast.Position{Offset: 6, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 6, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "[1, \"test\"]\n",
		},
		{
			name: "array literal with long content that should wrap",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{
					&ast.StringLiteral{
						Value: "This is a very long string that will definitely exceed the 80 character limit when combined with other elements",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 100, Line: 0, Column: 0},
						},
					},
					&ast.StringLiteral{
						Value: "Another long string",
						Range: ast.Range{
							Start: ast.Position{Offset: 101, Line: 0, Column: 0},
							End:   ast.Position{Offset: 120, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "42",
						Range: ast.Range{
							Start: ast.Position{Offset: 121, Line: 0, Column: 0},
							End:   ast.Position{Offset: 123, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 123, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected: `[
  "This is a very long string that will definitely exceed the 80 character limit when combined with other elements",
  "Another long string",
  42
]
`,
		},
		{
			name: "array literal with short content that should not wrap",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "2",
						Range: ast.Range{
							Start: ast.Position{Offset: 1, Line: 0, Column: 0},
							End:   ast.Position{Offset: 2, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "3",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 3, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "[1, 2, 3]\n",
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
