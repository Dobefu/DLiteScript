package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestIdentifierRegistry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input ast.ExprNode
	}{
		{
			name: "PI",
			input: &ast.Identifier{
				Value:    "PI",
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			name: "TAU",
			input: &ast.Identifier{
				Value:    "TAU",
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			name: "E",
			input: &ast.Identifier{
				Value:    "E",
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			name: "PHI",
			input: &ast.Identifier{
				Value:    "PHI",
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			name: "LN2",
			input: &ast.Identifier{
				Value:    "LN2",
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			name: "LN10",
			input: &ast.Identifier{
				Value:    "LN10",
				StartPos: 0,
				EndPos:   0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(io.Discard)

			_, err := evaluator.Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating '%s': %s", test.input.Expr(), err.Error())
			}
		})
	}
}
