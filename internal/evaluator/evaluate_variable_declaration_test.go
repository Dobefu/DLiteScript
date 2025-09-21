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

func TestEvaluateVariableDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.VariableDeclaration
		expected datavalue.Value
	}{
		{
			name: "variable declaration with value",
			input: &ast.VariableDeclaration{
				Name:     "x",
				Type:     datatype.DataTypeNumber.AsString(),
				Value:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: datavalue.Number(1),
		},
		{
			name: "variable declaration without value",
			input: &ast.VariableDeclaration{
				Name:     "x",
				Type:     datatype.DataTypeNumber.AsString(),
				Value:    nil,
				StartPos: 0,
				EndPos:   1,
			},
			expected: datavalue.Number(0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.evaluateVariableDeclaration(test.input)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			variable, hasVariable := ev.outerScope[test.input.Name]

			if !hasVariable {
				t.Fatalf("expected variable \"%s\" to be stored in scope", test.input.Name)
			}

			if variable.GetValue().DataType != test.expected.DataType {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.DataType.AsString(),
					variable.GetValue().DataType.AsString(),
				)
			}
		})
	}
}

func TestEvaluateVariableDeclarationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.VariableDeclaration
		expected string
	}{
		{
			name: "evaluation error",
			input: &ast.VariableDeclaration{
				Name: "x",
				Type: "int",
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
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "type mismatch",
			input: &ast.VariableDeclaration{
				Name:     "x",
				Type:     "int",
				Value:    &ast.StringLiteral{Value: "5", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeMismatch, "int", "string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateVariableDeclaration(test.input)

			if err == nil {
				t.Fatalf("expected error evaluating %s, got nil", test.input.Expr())
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, errors.Unwrap(err).Error())
			}
		})
	}
}

func TestGetZeroValueForType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected datavalue.Value
	}{
		{
			name:     "string",
			input:    "string",
			expected: datavalue.String(""),
		},
		{
			name:     "number",
			input:    "number",
			expected: datavalue.Number(0),
		},
		{
			name:     "bool",
			input:    "bool",
			expected: datavalue.Bool(false),
		},
		{
			name:     "any",
			input:    "any",
			expected: datavalue.Any(nil),
		},
		{
			name:     "error",
			input:    "error",
			expected: datavalue.Error(nil),
		},
		{
			name:     "unknown",
			input:    "unknown",
			expected: datavalue.Null(),
		},
		{
			name:     "string array",
			input:    "[]string",
			expected: datavalue.Array(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			if !ev.getZeroValueForType(test.input).Equals(test.expected) {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.ToString(),
					ev.getZeroValueForType(test.input).ToString(),
				)
			}
		})
	}
}
