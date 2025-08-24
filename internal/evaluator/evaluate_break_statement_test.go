package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
)

func TestEvaluateBreakStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.BreakStatement
		expected *controlflow.EvaluationResult
	}{
		{
			name: "zero count",
			input: &ast.BreakStatement{
				Count:    0,
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewBreakResult(0),
		},
		{
			name: "one count",
			input: &ast.BreakStatement{
				Count:    1,
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewBreakResult(1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := NewEvaluator(io.Discard).evaluateBreakStatement(test.input)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err)
			}

			if !result.IsBreakResult() {
				t.Fatalf("expected break result, got \"%v\"", result.Control.Type)
			}

			if result.Control.Count != test.expected.Control.Count {
				t.Fatalf("expected %d, got %d", test.expected.Control.Count, result.Control.Count)
			}
		})
	}
}
