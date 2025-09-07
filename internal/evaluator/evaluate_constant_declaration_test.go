package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateConstantDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *ast.ConstantDeclaration
	}{
		{
			name: "number",
			input: &ast.ConstantDeclaration{
				Name:     "x",
				Value:    &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Type:     datatype.DataTypeNumber.AsString(),
				StartPos: 0,
				EndPos:   1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateConstantDeclaration(test.input)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err)
			}
		})
	}
}

func TestEvaluateConstantDeclarationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.ConstantDeclaration
		expected string
	}{
		{
			name: "type mismatch",
			input: &ast.ConstantDeclaration{
				Name:     "x",
				Value:    &ast.StringLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Type:     datatype.DataTypeNumber.AsString(),
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(
				errorutil.ErrorMsgTypeMismatch,
				datatype.DataTypeNumber.AsString(),
				datatype.DataTypeString.AsString(),
			),
		},
		{
			name: "evaluation error",
			input: &ast.ConstantDeclaration{
				Name: "x",
				Type: datatype.DataTypeNumber.AsString(),
				Value: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "bogus",
					Arguments:    []ast.ExprNode{},
					StartPos:     4,
					EndPos:       18,
				},
				StartPos: 0,
				EndPos:   18,
			},
			expected: fmt.Sprintf(
				errorutil.ErrorMsgUndefinedFunction,
				"bogus",
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateConstantDeclaration(test.input)

			if err == nil {
				t.Fatalf("expected error evaluating %s, got nil", test.input.Expr())
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}
