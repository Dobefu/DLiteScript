package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatIfStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.IfStatement
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "if statement",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 2, Line: 0, Column: 0},
								End:   ast.Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
				ElseBlock: nil,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "if true {\n  1\n}\n",
		},
		{
			name: "if statement with empty body",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
				ElseBlock: nil,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "if true {}\n",
		},
		{
			name: "if statement with else block",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 2, Line: 0, Column: 0},
								End:   ast.Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
				ElseBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "2",
							Range: ast.Range{
								Start: ast.Position{Offset: 4, Line: 0, Column: 0},
								End:   ast.Position{Offset: 5, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "if true {\n  1\n} else {\n  2\n}\n",
		},
		{
			name: "if statement with empty else block",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 2, Line: 0, Column: 0},
								End:   ast.Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
				ElseBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "if true {\n  1\n} else {}\n",
		},
		{
			name: "if statement with empty then block and else block",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
				ElseBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 4, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "if true {} else {}\n",
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
