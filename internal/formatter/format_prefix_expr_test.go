package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestFormatPrefixExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.PrefixExpr
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "prefix expression",
			input: &ast.PrefixExpr{
				Operator: token.Token{
					Atom:      "!",
					TokenType: token.TokenTypeNot,
					StartPos:  0,
					EndPos:    0,
				},
				Operand:  &ast.BoolLiteral{Value: "true", StartPos: 1, EndPos: 2},
				StartPos: 0,
				EndPos:   2,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "!true\n",
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
