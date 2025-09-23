package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestFormatFuncDeclarationStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.FuncDeclarationStatement
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "func declaration statement",
			input: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					StartPos:   0,
					EndPos:     1,
				},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				StartPos:        0,
				EndPos:          1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "func test() {}\n",
		},
		{
			name: "func declaration statement with body",
			input: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				StartPos:        0,
				EndPos:          1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "func test() {\n  1\n}\n",
		},
		{
			name: "func declaration statement with nil body",
			input: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{nil},
					StartPos:   0,
					EndPos:     1,
				},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				StartPos:        0,
				EndPos:          1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "func test() {\n}\n",
		},
		{
			name: "func declaration statement with args and return values",
			input: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
					{Name: "b", Type: "string"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
								&ast.StringLiteral{Value: "test", StartPos: 1, EndPos: 6},
							},
							NumValues: 1,
							StartPos:  0,
							EndPos:    6,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				ReturnValues:    []string{"number", "string"},
				NumReturnValues: 2,
				StartPos:        0,
				EndPos:          1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "func test(a number, b string) number, string {\n  return 1, \"test\"\n}\n",
		},
		{
			name: "func declaration statement with expression body",
			input: &ast.FuncDeclarationStatement{
				Name: "add",
				Args: []ast.FuncParameter{
					{Name: "x", Type: "number"},
					{Name: "y", Type: "number"},
				},
				Body: &ast.BinaryExpr{
					Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
					Operator: token.Token{Atom: "+", TokenType: token.TokenTypeOperationAdd, StartPos: 1, EndPos: 2},
					Right:    &ast.Identifier{Value: "y", StartPos: 2, EndPos: 3},
					StartPos: 0,
					EndPos:   3,
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				StartPos:        0,
				EndPos:          3,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "func add(x number, y number) number {\n  x + y\n}\n",
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
