package global

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/function"
)

type testEvaluator struct {
	buf strings.Builder
}

func (e *testEvaluator) AddToBuffer(format string, args ...any) {
	fmt.Fprintf(&e.buf, format, args...)
}

func TestGetPrintfFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []datavalue.Value
		expected string
	}{
		{
			name: "single argument",
			input: []datavalue.Value{
				datavalue.String("test"),
			},
			expected: "test",
		},
		{
			name: "string argument",
			input: []datavalue.Value{
				datavalue.String("test %s"),
				datavalue.String("output"),
			},
			expected: "test output",
		},
		{
			name: "number argument",
			input: []datavalue.Value{
				datavalue.String("test %d"),
				datavalue.Number(1),
			},
			expected: "test 1",
		},
		{
			name: "boolean argument",
			input: []datavalue.Value{
				datavalue.String("test %t"),
				datavalue.Bool(true),
			},
			expected: "test true",
		},
		{
			name: "null argument",
			input: []datavalue.Value{
				datavalue.String("test %s"),
				datavalue.Null(),
			},
			expected: "test null",
		},
		{
			name: "function argument",
			input: []datavalue.Value{
				datavalue.String("test %s"),
				datavalue.Function(&ast.FuncDeclarationStatement{
					Name: "test",
					Args: []ast.FuncParameter{
						{Name: "a", Type: "number"},
					},
					Body: &ast.NumberLiteral{
						Value:    "1",
						StartPos: 0,
						EndPos:   3,
					},
					StartPos:        0,
					EndPos:          3,
					ReturnValues:    []string{"number"},
					NumReturnValues: 1,
				}),
			},
			expected: "test function",
		},
		{
			name: "tuple argument",
			input: []datavalue.Value{
				datavalue.String("test %s %d"),
				datavalue.Tuple(
					datavalue.String("test"),
					datavalue.Number(1),
				),
			},
			expected: "test test 1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := &testEvaluator{buf: strings.Builder{}}
			functions := GetGlobalFunctions()

			if _, ok := functions["printf"]; !ok {
				t.Fatalf("expected printf function, got %v", functions)
			}

			printf := functions["printf"]

			if printf.FunctionType != function.FunctionTypeMixedVariadic {
				t.Fatalf("expected mixed variadic function, got %v", printf.FunctionType)
			}

			result, err := printf.Handler(
				ev,
				test.input,
			)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result.DataType() != datatype.DataTypeNull {
				t.Fatalf("expected null, got %v", result.DataType())
			}
		})
	}
}
