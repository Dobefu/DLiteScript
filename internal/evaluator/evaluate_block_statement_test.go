package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestEvaluateBlockStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.BlockStatement
		expected datavalue.Value
	}{
		{
			name: "single statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			expected: datavalue.Number(5),
		},
		{
			name: "break statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.BreakStatement{Count: 1, StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   3,
			},
			expected: datavalue.Null(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rawResult, err := NewEvaluator(io.Discard).evaluateBlockStatement(test.input)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err.Error())
			}

			if rawResult.Value.DataType != test.expected.DataType {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.DataType.AsString(),
					rawResult.Value.DataType.AsString(),
				)
			}
		})
	}
}
