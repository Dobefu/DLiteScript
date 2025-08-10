package evaluator

import (
	"errors"
	"fmt"
	"math"
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

func TestEvaluateFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    ast.ExprNode
		expected float64
	}{
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "5", Pos: 0},
			),
			expected: 5,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"sin",
				&ast.NumberLiteral{Value: fmt.Sprintf("%f", math.Pi/2), Pos: 0},
			),
			expected: 1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"cos",
				&ast.NumberLiteral{Value: fmt.Sprintf("%f", math.Pi), Pos: 0},
			),
			expected: -1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"tan",
				&ast.NumberLiteral{Value: fmt.Sprintf("%f", math.Pi/4), Pos: 0},
			),
			expected: 1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"sqrt",
				&ast.NumberLiteral{Value: "16", Pos: 0},
			),
			expected: 4,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"round",
				&ast.NumberLiteral{Value: "3.14", Pos: 0},
			),
			expected: 3,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"floor",
				&ast.NumberLiteral{Value: "6.9", Pos: 0},
			),
			expected: 6,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"ceil",
				&ast.NumberLiteral{Value: "3.14", Pos: 0},
			),
			expected: 4,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"min",
				&ast.NumberLiteral{Value: "1", Pos: 0},
				&ast.NumberLiteral{Value: "2", Pos: 2},
			),
			expected: 1,
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"max",
				&ast.NumberLiteral{Value: "1", Pos: 0},
				&ast.NumberLiteral{Value: "2", Pos: 2},
			),
			expected: 2,
		},
	}

	for _, test := range tests {
		rawResult, err := NewEvaluator().Evaluate(test.input)

		if err != nil {
			t.Errorf("error evaluating '%s': %s", test.input, err.Error())
		}

		resultNum, err := rawResult.AsNumber()

		if err != nil {
			t.Fatalf("expected number, got type error: %s", err.Error())
		}

		result := math.Round(resultNum*1000) / 1000

		if result != test.expected {
			t.Errorf("expected '%f', got '%f'", test.expected, result)
		}
	}
}

func TestEvaluateFunctionCallErr(t *testing.T) {
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
