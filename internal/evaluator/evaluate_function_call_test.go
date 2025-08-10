package evaluator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func evaluateFunctionCallCreateFunctionCall(
	functionName string,
	arguments ...ast.ExprNode,
) ast.ExprNode {
	return &ast.FunctionCall{
		FunctionName: functionName,
		Arguments:    arguments,
		Pos:          0,
	}
}

func TestEvaluateFunctionCallPrint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected string
	}{
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"println",
				&ast.StringLiteral{Value: "test", Pos: 0},
			),
			expected: "test\n",
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"println",
				&ast.StringLiteral{Value: "testing, %g %g %g", Pos: 0},
				&ast.NumberLiteral{Value: "1", Pos: 10},
				&ast.NumberLiteral{Value: "2", Pos: 12},
				&ast.NumberLiteral{Value: "3", Pos: 14},
			),
			expected: "testing, 1 2 3\n",
		},
	}

	for _, test := range tests {
		ev := NewEvaluator()
		_, err := ev.Evaluate(test.input)

		if err != nil {
			t.Errorf("error evaluating '%s': %s", test.input, err.Error())
		}

		if ev.buf.String() != test.expected {
			t.Errorf("expected '%s', got '%s'", test.expected, ev.buf.String())
		}
	}
}

func TestEvaluateFunctionCallPrintErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected string
	}{
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"println",
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "println", 1, 0),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"println",
				&ast.NumberLiteral{Value: "1", Pos: 0},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionArgType, "println", 1, "string", "number"),
		},
	}

	for _, test := range tests {
		ev := NewEvaluator()
		_, err := ev.Evaluate(test.input)

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
	}
}

func TestEvaluateFunctionCallFixedArgsErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected string
	}{
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"bogus",
				&ast.NumberLiteral{Value: "1", Pos: 0},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "1", Pos: 0},
				&ast.NumberLiteral{Value: "1", Pos: 2},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "abs", 1, 2),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "a", Pos: 0},
			),
			expected: "invalid syntax",
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator().Evaluate(test.input)

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
	}
}
