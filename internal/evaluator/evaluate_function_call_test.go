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
	"github.com/Dobefu/DLiteScript/internal/function"
	"github.com/Dobefu/DLiteScript/internal/stdlib"
)

func TestEvaluateFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.FunctionCall
		expected string
	}{
		{
			name: "single argument",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{
						Value: "test",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: "test",
		},
		{
			name: "multiple arguments",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{
						Value: "testing, %g %g %g\n",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 10, Line: 0, Column: 0},
							End:   ast.Position{Offset: 11, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "2",
						Range: ast.Range{
							Start: ast.Position{Offset: 12, Line: 0, Column: 0},
							End:   ast.Position{Offset: 13, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "3",
						Range: ast.Range{
							Start: ast.Position{Offset: 14, Line: 0, Column: 0},
							End:   ast.Position{Offset: 15, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: "testing, 1 2 3\n",
		},
		{
			name: "spread array arguments",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{
						Value: "testing, %g %g %g\n",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					&ast.SpreadExpr{
						Expression: &ast.ArrayLiteral{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 10, Line: 0, Column: 0},
										End:   ast.Position{Offset: 11, Line: 0, Column: 0},
									},
								},
								&ast.NumberLiteral{
									Value: "2",
									Range: ast.Range{
										Start: ast.Position{Offset: 12, Line: 0, Column: 0},
										End:   ast.Position{Offset: 13, Line: 0, Column: 0},
									},
								},
								&ast.NumberLiteral{
									Value: "3",
									Range: ast.Range{
										Start: ast.Position{Offset: 14, Line: 0, Column: 0},
										End:   ast.Position{Offset: 15, Line: 0, Column: 0},
									},
								},
							},
							Range: ast.Range{
								Start: ast.Position{Offset: 10, Line: 0, Column: 0},
								End:   ast.Position{Offset: 15, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 10, Line: 0, Column: 0},
							End:   ast.Position{Offset: 15, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: "testing, 1 2 3\n",
		},
		{
			name: "user function",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "userFunction",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 10, Line: 0, Column: 0},
							End:   ast.Position{Offset: 11, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			ev.userFunctions[test.input.FunctionName] = &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			}

			_, err := ev.evaluateFunctionCall(test.input)

			if err != nil {
				t.Errorf("error evaluating \"%s\": %s", test.input.Expr(), err.Error())
			}

			if ev.buf.String() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, ev.buf.String())
			}
		})
	}
}

func TestEvaluateNamespaceFunctionCall(t *testing.T) {
	t.Parallel()

	input := &ast.FunctionCall{
		Namespace:    "testNamespace",
		FunctionName: "testFunc",
		Arguments: []ast.ExprNode{
			&ast.NumberLiteral{
				Value: "42",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 2, Line: 0, Column: 0},
				},
			},
		},
		Range: ast.Range{
			Start: ast.Position{Offset: 0, Line: 0, Column: 0},
			End:   ast.Position{Offset: 2, Line: 0, Column: 0},
		},
	}

	ev := NewEvaluator(io.Discard)
	ev.namespaceFunctions = map[string]map[string]*ast.FuncDeclarationStatement{
		"testNamespace": {
			"testFunc": {
				Name: "testFunc",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.Identifier{
									Value: "a",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
		},
	}

	result, err := ev.evaluateFunctionCall(input)

	if err != nil {
		t.Fatalf("error evaluating namespace function call: %s", err.Error())
	}

	expected := "42"
	if result.Value.ToString() != expected {
		t.Errorf("expected \"%s\", got \"%s\"", expected, result.Value.ToString())
	}
}

func TestEvaluateFunctionCallErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.FunctionCall
		expected string
	}{
		{
			name: "no arguments",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments:    []ast.ExprNode{},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "printf", 1, 0),
		},
		{
			name: "single argument",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				errorutil.ErrorMsgFunctionArgType,
				"printf",
				1,
				"string",
				"number",
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.evaluateFunctionCall(test.input)

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
		input    *ast.FunctionCall
		expected string
	}{
		{
			name: "undefined namespace",
			input: &ast.FunctionCall{
				Namespace:    "bogus",
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedNamespace, "bogus"),
		},
		{
			name: "undefined function",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "bogus",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
		{
			name: "too many arguments",
			input: &ast.FunctionCall{
				Namespace:    "math",
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 3, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "math.abs", 1, 2),
		},
		{
			name: "invalid argument",
			input: &ast.FunctionCall{
				Namespace:    "math",
				FunctionName: "abs",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "a",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "invalid syntax",
		},
		{
			name: "function handler error",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "functionHandlerError",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{
						Value: "test",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.StringLiteral{
						Value: "extra",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 3, Line: 0, Column: 0},
				},
			},
			expected: "'functionHandlerError()' expects exactly 1 argument(s), but got 2",
		},
		{
			name: "undefined function in existing namespace",
			input: &ast.FunctionCall{
				Namespace:    "math",
				FunctionName: "bogus",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "bogus"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			if test.name == "function handler error" {
				ev.userFunctions[test.input.FunctionName] = &ast.FuncDeclarationStatement{ //nolint:exhaustruct
					Name: test.input.FunctionName,
					Args: []ast.FuncParameter{
						{Name: "arg", Type: "string"},
					},
					Body: &ast.BlockStatement{
						Statements: []ast.ExprNode{
							&ast.ReturnStatement{
								Values: []ast.ExprNode{
									&ast.StringLiteral{
										Value: "test",
										Range: ast.Range{
											Start: ast.Position{Offset: 0, Line: 0, Column: 0},
											End:   ast.Position{Offset: 4, Line: 0, Column: 0},
										},
									},
								},
								NumValues: 1,
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 4, Line: 0, Column: 0},
								},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 4, Line: 0, Column: 0},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 4, Line: 0, Column: 0},
					},
				}
			}

			_, err := ev.evaluateFunctionCall(test.input)

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
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "1",
		},
		{
			name: "tuple return value",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "2",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 3, Line: 0, Column: 0},
				},
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
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
								&ast.NumberLiteral{
									Value: "2",
									Range: ast.Range{
										Start: ast.Position{Offset: 2, Line: 0, Column: 0},
										End:   ast.Position{Offset: 3, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 2,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number", "number"},
				NumReturnValues: 2,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "test", 2, 1),
		},
		{
			name: "invalid argument",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "a",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "invalid syntax",
		},
		{
			name: "invalid function body",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			functionDeclaration: &ast.FuncDeclarationStatement{
				Name: "test",
				Args: []ast.FuncParameter{
					{Name: "a", Type: "number"},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "a",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "invalid syntax",
		},
		{
			name: "multiple return values but single value returned",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number", "number"},
				NumReturnValues: 2,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionReturnCount, "test", 2, 1),
		},
		{
			name: "multiple return values but not enough returned",
			functionCall: &ast.FunctionCall{
				FunctionName: "test",
				Namespace:    "",
				Arguments: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
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
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 1, Line: 0, Column: 0},
									},
								},
								&ast.NumberLiteral{
									Value: "2",
									Range: ast.Range{
										Start: ast.Position{Offset: 2, Line: 0, Column: 0},
										End:   ast.Position{Offset: 3, Line: 0, Column: 0},
									},
								},
								&ast.NumberLiteral{
									Value: "3",
									Range: ast.Range{
										Start: ast.Position{Offset: 4, Line: 0, Column: 0},
										End:   ast.Position{Offset: 5, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 3,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 5, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				ReturnValues:    []string{"number", "number"},
				NumReturnValues: 2,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionReturnCount, "test", 2, 3),
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

func TestEvaluateArgumentsSpreadExprErr(t *testing.T) {
	t.Parallel()

	input := &ast.FunctionCall{
		Namespace:    "",
		FunctionName: "printf",
		Arguments: []ast.ExprNode{
			&ast.StringLiteral{
				Value: "test %g\n",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			&ast.SpreadExpr{
				Expression: &ast.NumberLiteral{
					Value: "invalid",
					Range: ast.Range{
						Start: ast.Position{Offset: 10, Line: 0, Column: 0},
						End:   ast.Position{Offset: 11, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 10, Line: 0, Column: 0},
					End:   ast.Position{Offset: 11, Line: 0, Column: 0},
				},
			},
		},
		Range: ast.Range{
			Start: ast.Position{Offset: 0, Line: 0, Column: 0},
			End:   ast.Position{Offset: 11, Line: 0, Column: 0},
		},
	}

	ev := NewEvaluator(io.Discard)
	_, err := ev.evaluateFunctionCall(input)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if errors.Unwrap(err).Error() != "invalid syntax" {
		t.Errorf(
			"expected \"invalid syntax\", got \"%s\"",
			errors.Unwrap(err).Error(),
		)
	}
}

func TestEvaluateArgumentsSpreadTuple(t *testing.T) {
	t.Parallel()

	functionRegistry := stdlib.GetFunctionRegistry()
	functionRegistry[""].Functions["spreadTest"] = function.Info{ //nolint:exhaustruct
		FunctionType: function.FunctionTypeFixed,
		Parameters:   []function.ArgInfo{},
		Handler: func(
			_ function.EvaluatorInterface,
			_ []datavalue.Value,
		) (datavalue.Value, error) {
			return datavalue.Tuple(
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			), nil
		},
	}
	functionRegistry["spreadTest"] = function.PackageInfo{
		Functions: map[string]function.Info{
			"spreadFunctionHandlerError": { //nolint:exhaustruct
				FunctionType: function.FunctionTypeFixed,
				Parameters: []function.ArgInfo{
					{Type: datatype.DataTypeString}, //nolint:exhaustruct
				},
				Handler: func(
					_ function.EvaluatorInterface,
					_ []datavalue.Value,
				) (datavalue.Value, error) {
					return datavalue.Null(), errorutil.NewErrorAt(
						errorutil.StageEvaluate,
						"handler error",
						ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					)
				},
			},
		},
	}

	input := &ast.FunctionCall{
		Namespace:    "",
		FunctionName: "printf",
		Arguments: []ast.ExprNode{
			&ast.StringLiteral{
				Value: "testing tuple spread: %g %g %g\n",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			&ast.SpreadExpr{
				Expression: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "spreadTest",
					Arguments:    []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 10, Line: 0, Column: 0},
						End:   ast.Position{Offset: 15, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 10, Line: 0, Column: 0},
					End:   ast.Position{Offset: 15, Line: 0, Column: 0},
				},
			},
		},
		Range: ast.Range{
			Start: ast.Position{Offset: 0, Line: 0, Column: 0},
			End:   ast.Position{Offset: 15, Line: 0, Column: 0},
		},
	}

	ev := NewEvaluator(io.Discard)
	_, err := ev.evaluateFunctionCall(input)

	if err != nil {
		t.Fatalf("expected no error, got: %s", err.Error())
	}

	expected := "testing tuple spread: 1 2 3\n"

	if ev.buf.String() != expected {
		t.Errorf("expected \"%s\", got \"%s\"", expected, ev.buf.String())
	}
}

func TestValidateVariadicArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        []ast.ExprNode
		functionInfo function.Info
		expected     []datavalue.Value
	}{
		{
			name: "no parameters",
			input: []ast.ExprNode{
				&ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
			},
			functionInfo: function.Info{
				Documentation: function.Documentation{
					Name:        "test",
					Description: "",
					Since:       "",
					Examples:    []string{},
					DeprecationInfo: function.DeprecationInfo{
						IsDeprecated: false,
						Description:  "",
						Version:      "",
					},
				},
				FunctionType: function.FunctionTypeVariadic,
				PackageName:  "",
				IsBuiltin:    false,
				Parameters:   []function.ArgInfo{},
				ReturnValues: []function.ArgInfo{
					{
						Type:        datatype.DataTypeNumber,
						Name:        "result",
						Description: "",
					},
				},
				Handler: nil,
			},
			expected: []datavalue.Value{datavalue.Number(1)},
		},
		{
			name: "single argument",
			input: []ast.ExprNode{
				&ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
			},
			functionInfo: function.Info{
				Documentation: function.Documentation{
					Name:        "test",
					Description: "",
					Since:       "",
					Examples:    []string{},
					DeprecationInfo: function.DeprecationInfo{
						IsDeprecated: false,
						Description:  "",
						Version:      "",
					},
				},
				FunctionType: function.FunctionTypeVariadic,
				PackageName:  "",
				IsBuiltin:    false,
				Parameters: []function.ArgInfo{
					{
						Name:        "a",
						Type:        datatype.DataTypeNumber,
						Description: "",
					},
					{
						Name:        "b",
						Type:        datatype.DataTypeAny,
						Description: "",
					},
				},
				ReturnValues: []function.ArgInfo{
					{
						Type:        datatype.DataTypeNumber,
						Name:        "result",
						Description: "",
					},
				},
				Handler: nil,
			},
			expected: []datavalue.Value{datavalue.Number(1)},
		},
		{
			name: "any argument",
			input: []ast.ExprNode{
				&ast.AnyLiteral{
					Value: &ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
			},
			functionInfo: function.Info{
				Documentation: function.Documentation{
					Name:        "test",
					Description: "",
					Since:       "",
					Examples:    []string{},
					DeprecationInfo: function.DeprecationInfo{
						IsDeprecated: false,
						Description:  "",
						Version:      "",
					},
				},
				FunctionType: function.FunctionTypeVariadic,
				PackageName:  "",
				IsBuiltin:    false,
				Parameters: []function.ArgInfo{
					{
						Name:        "a",
						Type:        datatype.DataTypeAny,
						Description: "",
					},
				},
				ReturnValues: []function.ArgInfo{
					{
						Type:        datatype.DataTypeNumber,
						Name:        "result",
						Description: "",
					},
				},
				Handler: nil,
			},
			expected: []datavalue.Value{datavalue.Number(1)},
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
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
			)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
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
				&ast.StringLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
			},
			functionInfo: function.Info{ //nolint:exhaustruct
				FunctionType: function.FunctionTypeVariadic,
				Parameters: []function.ArgInfo{
					{Type: datatype.DataTypeNumber}, //nolint:exhaustruct
					{Type: datatype.DataTypeAny},    //nolint:exhaustruct
				},
				Handler: nil,
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
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
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
