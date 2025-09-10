package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateIndexAssignmentStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.IndexAssignmentStatement
		expected float64
	}{
		{
			name: "assignment to array variable index",
			input: &ast.IndexAssignmentStatement{
				Array:    &ast.Identifier{Value: "someArray", StartPos: 0, EndPos: 9},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   9,
			},
			expected: 1,
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

			result, err := ev.evaluateIndexAssignmentStatement(test.input)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			array, err := result.Value.AsArray()

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			num, err := array[0].AsNumber()

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if num != test.expected {
				t.Errorf("expected %f, got: %f", test.expected, num)
			}
		})
	}
}

func TestEvaluateIndexAssignmentStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.IndexAssignmentStatement
		expected string
	}{
		{
			name: "evaluation error",
			input: &ast.IndexAssignmentStatement{
				Array: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "bogus",
					Arguments:    []ast.ExprNode{},
					StartPos:     0,
					EndPos:       1,
				},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "not an array",
			input: &ast.IndexAssignmentStatement{
				Array:    &ast.NullLiteral{StartPos: 0, EndPos: 1},
				Index:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "array", "null"),
		},
		{
			name: "evaluation error for index",
			input: &ast.IndexAssignmentStatement{
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
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "index is not a number",
			input: &ast.IndexAssignmentStatement{
				Array: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				Index:    &ast.StringLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "\"0\""),
		},
		{
			name: "evaluation error for right",
			input: &ast.IndexAssignmentStatement{
				Array: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				Index: &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right: &ast.FunctionCall{
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
			name: "index is out of bounds",
			input: &ast.IndexAssignmentStatement{
				Array: &ast.ArrayLiteral{
					Values: []ast.ExprNode{
						&ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				Index:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgArrayIndexOutOfBounds, "1"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.evaluateIndexAssignmentStatement(test.input)

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
