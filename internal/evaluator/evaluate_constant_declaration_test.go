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
				Name: "x",
				Value: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Type: datatype.DataTypeNumber.AsString(),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
				Name: "x",
				Value: &ast.StringLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Type: datatype.DataTypeNumber.AsString(),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 18, Line: 0, Column: 0},
				},
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
