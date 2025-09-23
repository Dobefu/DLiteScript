package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestFormatBinaryExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.BinaryExpr
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "binary expression",
			input: &ast.BinaryExpr{
				Left:  &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				StartPos: 0,
				EndPos:   0,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "1 + 1\n",
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
