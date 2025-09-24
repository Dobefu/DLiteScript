package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestEvaluateNullLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected datavalue.Value
	}{
		{
			name: "null literal",
			input: &ast.NullLiteral{
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Null(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := NewEvaluator(io.Discard).Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating null literal: \"%s\"", err.Error())
			}

			if result.Value.DataType.AsString() != test.expected.DataType.AsString() {
				t.Fatalf(
					"expected \"%v\", got \"%v\" at position %d",
					test.expected.DataType.AsString(),
					result.Value.DataType.AsString(),
					test.input.GetRange().Start.Offset,
				)
			}
		})
	}
}
