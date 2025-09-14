package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetCeilFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "positive number (1.5)",
			input:    datavalue.Number(1.5),
			expected: datavalue.Number(2),
		},
		{
			name:     "positive number (1.2)",
			input:    datavalue.Number(1.2),
			expected: datavalue.Number(2),
		},
		{
			name:     "positive number (1)",
			input:    datavalue.Number(1),
			expected: datavalue.Number(1),
		},
		{
			name:     "negative number (-1.5)",
			input:    datavalue.Number(-1.5),
			expected: datavalue.Number(-1),
		},
	}

	functions := GetMathFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ceilFunc, hasFunction := functions["ceil"]

			if !hasFunction {
				t.Fatalf("could not find ceil function")
			}

			result, err := ceilFunc.Handler(nil, []datavalue.Value{test.input})

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if result.Num != test.expected.Num {
				t.Fatalf("expected %f, got %f", test.expected.Num, result.Num)
			}
		})
	}
}
