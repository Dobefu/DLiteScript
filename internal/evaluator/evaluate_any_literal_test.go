package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestEvaluateAnyLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected *controlflow.EvaluationResult
	}{
		{
			name:     "any literal",
			input:    &ast.AnyLiteral{Value: "1", StartPos: 0, EndPos: 1},
			expected: controlflow.NewRegularResult(datavalue.Any(1)),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			result, err := ev.Evaluate(test.input)

			if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			if !result.Value.Equals(test.expected.Value) {
				t.Errorf("expected '%v', got '%v'", test.expected.Value, result.Value)
			}
		})
	}
}
