package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateArrayLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected *controlflow.EvaluationResult
	}{
		{
			name: "array literal",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(
				datavalue.Array(datavalue.Number(1)),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			result, err := ev.Evaluate(test.input)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if !result.Value.Equals(test.expected.Value) {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected.Value.ToString(),
					result.Value.ToString(),
				)
			}
		})
	}
}

func TestEvaluateArrayLiteralErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "evaluation error",
			input: &ast.ArrayLiteral{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					&ast.FunctionCall{
						Namespace:    "",
						FunctionName: "bogus",
						Arguments:    []ast.ExprNode{},
						StartPos:     0,
						EndPos:       1,
					},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.Evaluate(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
