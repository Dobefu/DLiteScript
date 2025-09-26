package math

import (
	"errors"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetMaxFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []datavalue.Value
		expected datavalue.Value
	}{
		{
			name:  "no input",
			input: []datavalue.Value{},
			expected: datavalue.Tuple(
				datavalue.Null(),
				datavalue.Error(errors.New("max requires at least 2 arguments, got 0")),
			),
		},
		{
			name: "positive numbers (1, 2, 3)",
			input: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
			expected: datavalue.Tuple(datavalue.Number(3), datavalue.Null()),
		},
		{
			name: "positive numbers (1.5, 2.5, 3.5)",
			input: []datavalue.Value{
				datavalue.Number(1.5),
				datavalue.Number(2.5),
				datavalue.Number(3.5),
			},
			expected: datavalue.Tuple(datavalue.Number(3.5), datavalue.Null()),
		},
		{
			name: "negative numbers (-1, -2, -3)",
			input: []datavalue.Value{
				datavalue.Number(-1),
				datavalue.Number(-2),
				datavalue.Number(-3),
			},
			expected: datavalue.Tuple(datavalue.Number(-1), datavalue.Null()),
		},
		{
			name: "negative numbers (-1.5, -2.5, -3.5)",
			input: []datavalue.Value{
				datavalue.Number(-1.5),
				datavalue.Number(-2.5),
				datavalue.Number(-3.5),
			},
			expected: datavalue.Tuple(datavalue.Number(-1.5), datavalue.Null()),
		},
	}

	functions := GetMathFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			maxFunc, hasFunction := functions["max"]

			if !hasFunction {
				t.Fatalf("could not find max function")
			}

			result, err := maxFunc.Handler(nil, test.input)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if result.Num != test.expected.Num {
				t.Fatalf("expected %f, got %f", test.expected.Num, result.Num)
			}
		})
	}
}
