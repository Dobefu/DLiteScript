package formatter

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestFormatShorthandAssignmentExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.ShorthandAssignmentExpr
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name: "shorthand assignment expression",
			input: &ast.ShorthandAssignmentExpr{
				Left:  &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right: &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken(
					"+=",
					token.TokenTypeOperationAddAssign,
					0,
					1,
				),
				StartPos: 0,
				EndPos:   1,
			},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "x += 1\n",
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
