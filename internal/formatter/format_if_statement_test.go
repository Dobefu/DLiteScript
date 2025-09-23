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
				Condition: &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "if true {\n  1\n}\n",
		},
		{
			name: "if statement with empty body",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					StartPos:   2,
					EndPos:     3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "if true {}\n",
		},
		{
			name: "if statement with else block",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   5,
				ElseBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{Value: "2", StartPos: 4, EndPos: 5},
					},
					StartPos: 4,
					EndPos:   5,
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "if true {\n  1\n} else {\n  2\n}\n",
		},
		{
			name: "if statement with empty else block",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   5,
				ElseBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					StartPos:   4,
					EndPos:     5,
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "if true {\n  1\n} else {}\n",
		},
		{
			name: "if statement with empty then block and else block",
			input: &ast.IfStatement{
				Condition: &ast.BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					StartPos:   2,
					EndPos:     3,
				},
				StartPos: 0,
				EndPos:   5,
				ElseBlock: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					StartPos:   4,
					EndPos:     5,
				},
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
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
