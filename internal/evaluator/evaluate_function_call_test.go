package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func evaluateFunctionCallCreateFunctionCall(
	functionName string,
	arguments ...ast.ExprNode,
) ast.ExprNode {
	if len(arguments) == 0 {
		return &ast.FunctionCall{
			FunctionName: functionName,
			Arguments:    arguments,
			StartPos:     0,
			EndPos:       0,
		}
	}

	return &ast.FunctionCall{
		FunctionName: functionName,
		Arguments:    arguments,
		StartPos:     arguments[0].StartPosition(),
		EndPos:       arguments[len(arguments)-1].EndPosition(),
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
				"printf",
				&ast.StringLiteral{Value: "test", StartPos: 0, EndPos: 1},
			),
			expected: "test",
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"printf",
				&ast.StringLiteral{Value: "testing, %g %g %g\n", StartPos: 0, EndPos: 1},
				&ast.NumberLiteral{Value: "1", StartPos: 10, EndPos: 11},
				&ast.NumberLiteral{Value: "2", StartPos: 12, EndPos: 13},
				&ast.NumberLiteral{Value: "3", StartPos: 14, EndPos: 15},
			),
			expected: "testing, 1 2 3\n",
		},
	}

	for _, test := range tests {
		ev := NewEvaluator(io.Discard)
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
				"printf",
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "printf", 1, 0),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"printf",
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionArgType, "printf", 1, "string", "number"),
		},
	}

	for _, test := range tests {
		ev := NewEvaluator(io.Discard)
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
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				&ast.NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "abs", 1, 2),
		},
		{
			input: evaluateFunctionCallCreateFunctionCall(
				"abs",
				&ast.NumberLiteral{Value: "a", StartPos: 0, EndPos: 1},
			),
			expected: "invalid syntax",
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator(io.Discard).Evaluate(test.input)

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
