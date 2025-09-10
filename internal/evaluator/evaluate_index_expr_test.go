package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateIndexExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.IndexExpr
		expected float64
	}{
		{
			name: "index expression",
			input: &ast.IndexExpr{
				Array:    &ast.Identifier{Value: "someArray", StartPos: 0, EndPos: 1},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			ev.outerScope["someArray"] = &Variable{
				Value: datavalue.Array(datavalue.Number(0)),
				Type:  "array",
			}

			result, err := ev.evaluateIndexExpr(test.input)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if result.Value.DataType() != datatype.DataTypeNumber {
				t.Fatalf("expected number result, got: %v", result.Value.DataType())
			}

			number, err := result.Value.AsNumber()

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if number != test.expected {
				t.Errorf("expected %f, got %f", test.expected, number)
			}
		})
	}
}

func TestEvaluateIndexExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.IndexExpr
		expected string
	}{
		{
			name: "evaluation error",
			input: &ast.IndexExpr{
				Array: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "bogus",
					Arguments:    []ast.ExprNode{},
					StartPos:     0,
					EndPos:       1,
				},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "not an array",
			input: &ast.IndexExpr{
				Array:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "array", "number"),
		},
		{
			name: "index evaluation error",
			input: &ast.IndexExpr{
				Array: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				Index: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "bogus",
					Arguments:    []ast.ExprNode{},
					StartPos:     0,
					EndPos:       1,
				},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "index is not a number",
			input: &ast.IndexExpr{
				Array: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				Index:    &ast.StringLiteral{Value: "0", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "string"),
		},
		{
			name: "index is out of bounds",
			input: &ast.IndexExpr{
				Array: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				Index:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgArrayIndexOutOfBounds, "1"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateIndexExpr(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Errorf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}
