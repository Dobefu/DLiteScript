package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatVariableDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.VariableDeclaration
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "variable declaration",
			input: &ast.VariableDeclaration{
				Name: "x",
				Type: "int",
				Value: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "var x int = 1\n",
		},
		{
			name: "variable declaration with array literal",
			input: &ast.VariableDeclaration{
				Name: "x",
				Type: "[]string",
				Value: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.StringLiteral{
							Value: "test1",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 5, Line: 0, Column: 0},
							},
						},
						&ast.StringLiteral{
							Value: "test2",
							Range: ast.Range{
								Start: ast.Position{Offset: 7, Line: 0, Column: 0},
								End:   ast.Position{Offset: 12, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 12, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 12, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 0},
			depth:     0,
			expected:  "var x []string = [\n  \"test1\",\n  \"test2\",\n]\n",
		},
		{
			name: "variable declaration without value",
			input: &ast.VariableDeclaration{
				Name:  "y",
				Type:  "string",
				Value: nil,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "var y string\n",
		},
		{
			name: "variable declaration with array literal containing nil elements",
			input: &ast.VariableDeclaration{
				Name: "z",
				Type: "[]string",
				Value: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.StringLiteral{
							Value: "test1",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 5, Line: 0, Column: 0},
							},
						},
						nil,
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 12, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 12, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 0},
			depth:     0,
			expected:  "var z []string = [\n  \"test1\",\n]\n",
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
