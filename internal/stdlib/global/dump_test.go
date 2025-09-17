package global

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestDump(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []datavalue.Value
		expected string
	}{
		{
			name: "string value",
			input: []datavalue.Value{
				datavalue.String("test"),
			},
			expected: "\"test\"\n",
		},
		{
			name: "number value",
			input: []datavalue.Value{
				datavalue.Number(1),
			},
			expected: "1\n",
		},
		{
			name: "boolean value",
			input: []datavalue.Value{
				datavalue.Bool(true),
			},
			expected: "true\n",
		},
		{
			name: "null value",
			input: []datavalue.Value{
				datavalue.Null(),
			},
			expected: "null\n",
		},
		{
			name: "function value",
			input: []datavalue.Value{
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
					ReturnValues:    []string{"number"},
					NumReturnValues: 1,
					StartPos:        0,
					EndPos:          3,
				}),
			},
			expected: "function\n",
		},
		{
			name: "array value",
			input: []datavalue.Value{
				datavalue.Array(
					datavalue.Number(1),
					datavalue.Number(2),
					datavalue.Number(3),
				),
			},
			expected: "array[3]:\n  [0]:   1\n  [1]:   2\n  [2]:   3\n",
		},
		{
			name: "tuple value",
			input: []datavalue.Value{
				datavalue.Tuple(
					datavalue.Number(1),
					datavalue.Number(2),
					datavalue.Number(3),
				),
			},
			expected: "tuple[3]:\n  (0):   1\n  (1):   2\n  (2):   3\n",
		},
		{
			name: "any value",
			input: []datavalue.Value{
				datavalue.Any("test"),
			},
			expected: "test\n",
		},
		{
			name: "multiple values",
			input: []datavalue.Value{
				datavalue.Number(1),
				datavalue.String("test"),
			},
			expected: "1\n\"test\"\n",
		},
	}

	functions := GetGlobalFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := &testEvaluator{buf: strings.Builder{}, exitCode: 0}
			dumpFunc, hasDump := functions["dump"]

			if !hasDump {
				t.Fatalf("expected dump function, got %v", functions)
			}

			_, err := dumpFunc.Handler(evaluator, test.input)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if evaluator.buf.String() != test.expected {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					evaluator.buf.String(),
				)
			}
		})
	}
}
