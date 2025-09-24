package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.FunctionCall
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "function call",
			input: &ast.FunctionCall{
				Namespace:    "math",
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
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
			expected:  "math.abs(1)\n",
		},
		{
			name: "function call with multiple arguments",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{
						Value: "test %s",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 8, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 9, Line: 0, Column: 0},
							End:   ast.Position{Offset: 10, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 10, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "printf(\"test %s\", 1)\n",
		},
		{
			name: "function call with nil argument",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments:    []ast.ExprNode{nil, nil},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 8, Line: 0, Column: 0},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "printf()\n",
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
