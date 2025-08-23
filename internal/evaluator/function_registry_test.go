package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFunctionRegistry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input *ast.FunctionCall
	}{
		{
			input: &ast.FunctionCall{
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{Value: "test", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "-5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "sin",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "0.5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "cos",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "0.5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "tan",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "0.5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "sqrt",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "25", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "round",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1.5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "floor",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1.5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "ceil",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1.5", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "min",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "10", StartPos: 0, EndPos: 1},
					&ast.NumberLiteral{Value: "20", StartPos: 2, EndPos: 3},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
		{
			input: &ast.FunctionCall{
				FunctionName: "max",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "10", StartPos: 0, EndPos: 1},
					&ast.NumberLiteral{Value: "20", StartPos: 2, EndPos: 3},
				},
				StartPos: 0,
				EndPos:   0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input.FunctionName, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(io.Discard)

			_, err := evaluator.Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating '%s': %s", test.input.FunctionName, err.Error())
			}
		})
	}
}
