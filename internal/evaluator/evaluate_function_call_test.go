package evaluator

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/function"
)

func evaluateFunctionCallCreateFunctionCall(
	functionName string,
	arguments ...ast.ExprNode,
) ast.ExprNode {
	namespace := ""
	fullFunctionName := functionName

	if strings.Contains(functionName, ".") {
		parts := strings.Split(functionName, ".")

		if len(parts) == 2 {
			namespace = parts[0]
			fullFunctionName = parts[1]
		}
	}

	if len(arguments) == 0 {
		return &ast.FunctionCall{
			Namespace:    namespace,
			FunctionName: fullFunctionName,
			Arguments:    arguments,
			StartPos:     0,
			EndPos:       0,
		}
	}

	return &ast.FunctionCall{
		Namespace:    namespace,
		FunctionName: fullFunctionName,
		Arguments:    arguments,
		StartPos:     arguments[0].StartPosition(),
		EndPos:       arguments[len(arguments)-1].EndPosition(),
	}
}

func TestEvaluateFunctionCallPrint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "single argument",
			input: evaluateFunctionCallCreateFunctionCall(
				"printf",
				&ast.StringLiteral{Value: "test", StartPos: 0, EndPos: 1},
			),
			expected: "test",
		},
		{
			name: "multiple arguments",
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
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.Evaluate(test.input)

			if err != nil {
				t.Errorf("error evaluating '%s': %s", test.input, err.Error())
			}

			if ev.buf.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, ev.buf.String())
			}
		})
	}
}

func TestEvaluateFunctionCallPrintErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "no arguments",
			input: evaluateFunctionCallCreateFunctionCall(
				"printf",
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "printf", 1, 0),
		},
		{
			name: "single argument",
			input: evaluateFunctionCallCreateFunctionCall(
				"printf",
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionArgType, "printf", 1, "string", "number"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
		})
	}
}

func TestEvaluateFunctionCallFixedArgsErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "undefined function",
			input: evaluateFunctionCallCreateFunctionCall(
				"bogus",
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "too many arguments",
			input: evaluateFunctionCallCreateFunctionCall(
				"math.abs",
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				&ast.NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
			),
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "math.abs", 1, 2),
		},
		{
			name: "invalid argument",
			input: evaluateFunctionCallCreateFunctionCall(
				"math.abs",
				&ast.NumberLiteral{Value: "a", StartPos: 0, EndPos: 1},
			),
			expected: "invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
		})
	}
}

func TestEvaluateUserFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		functionCall        *ast.FunctionCall
		functionDeclaration *ast.FuncDeclarationStatement
		expected            string
	}{
		{
			name: "single argument",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			functionDeclaration: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
							},
							NumValues: 1,
							StartPos:  0,
							EndPos:    1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				StartPos:        0,
				EndPos:          1,
			},
			expected: "1",
		},
		{
			name: "tuple return value",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					&ast.NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
				},
				StartPos: 0,
				EndPos:   3,
			},
			functionDeclaration: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
					{Name: "b", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
								&ast.NumberLiteral{Value: "2", StartPos: 2, EndPos: 3},
							},
							NumValues: 2,
							StartPos:  0,
							EndPos:    3,
						},
					},
					StartPos: 0,
					EndPos:   3,
				},
				ReturnValues:    []string{"number", "number"},
				NumReturnValues: 2,
				StartPos:        0,
				EndPos:          1,
			},
			expected: "(1, 2)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			result, err := ev.evaluateUserFunctionCall(
				test.functionCall,
				test.functionDeclaration,
			)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if result.Value.ToString() != test.expected {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					result.Value.ToString(),
				)
			}
		})
	}
}

func TestEvaluateUserFunctionCallErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		functionCall        *ast.FunctionCall
		functionDeclaration *ast.FuncDeclarationStatement
		expected            string
	}{
		{
			name: "not enough arguments",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			functionDeclaration: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
					{Name: "b", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
							},
							NumValues: 1,
							StartPos:  0,
							EndPos:    1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				StartPos:        0,
				EndPos:          1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "test", 2, 1),
		},
		{
			name: "invalid argument",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "a", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			functionDeclaration: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
							},
							NumValues: 1,
							StartPos:  0,
							EndPos:    1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				StartPos:        0,
				EndPos:          1,
			},
			expected: "invalid syntax",
		},
		{
			name: "invalid function body",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				},
				StartPos: 0,
				EndPos:   1,
			},
			functionDeclaration: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{Value: "a", StartPos: 0, EndPos: 1},
					},
					StartPos: 0,
					EndPos:   1,
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				StartPos:        0,
				EndPos:          1,
			},
			expected: "invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.evaluateUserFunctionCall(
				test.functionCall,
				test.functionDeclaration,
			)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}

func TestValidateVariadicArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        []ast.ExprNode
		functionInfo function.Info
		expected     string
	}{
		{
			name: "single argument",
			input: []ast.ExprNode{
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			},
			functionInfo: function.Info{
				FunctionType: function.FunctionTypeVariadic,
				ArgKinds:     []datatype.DataType{datatype.DataTypeNumber},
				Handler:      nil,
			},
			expected: "1",
		},
		{
			name: "no arguments",
			input: []ast.ExprNode{
				&ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			},
			functionInfo: function.Info{
				FunctionType: function.FunctionTypeVariadic,
				ArgKinds:     []datatype.DataType{},
				Handler:      nil,
			},
			expected: "1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			argValues := make([]datavalue.Value, len(test.input))

			for i, arg := range test.input {
				value, err := ev.Evaluate(arg)

				if err != nil {
					t.Fatalf("failed to evaluate argument %d: %s", i, err.Error())
				}

				argValues[i] = value.Value
			}

			result, err := ev.validateVariadicArgs(
				argValues,
				test.functionInfo,
				"test",
				&ast.FunctionCall{
					FunctionName: "test",
					Namespace:    "",
					Arguments:    test.input,
					StartPos:     0,
					EndPos:       0,
				},
			)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if result[0].ToString() != test.expected {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					result[0].ToString(),
				)
			}
		})
	}
}

func TestValidateVariadicArgsErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        []ast.ExprNode
		functionInfo function.Info
		expected     string
	}{
		{
			name: "invalid argument type",
			input: []ast.ExprNode{
				&ast.StringLiteral{Value: "1", StartPos: 0, EndPos: 1},
			},
			functionInfo: function.Info{
				FunctionType: function.FunctionTypeVariadic,
				ArgKinds:     []datatype.DataType{datatype.DataTypeNumber},
				Handler:      nil,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionArgType, "test", 1, "number", "string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			argValues := make([]datavalue.Value, len(test.input))

			for i, arg := range test.input {
				value, err := ev.Evaluate(arg)

				if err != nil {
					t.Fatalf("failed to evaluate argument %d: %s", i, err.Error())
				}

				argValues[i] = value.Value
			}

			_, err := ev.validateVariadicArgs(
				argValues,
				test.functionInfo,
				"test",
				&ast.FunctionCall{
					FunctionName: "test",
					Namespace:    "",
					Arguments:    test.input,
					StartPos:     0,
					EndPos:       0,
				},
			)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}
