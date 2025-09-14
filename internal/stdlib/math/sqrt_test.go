package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSqrtFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "positive number (4)",
			input:    datavalue.Number(4),
			expected: datavalue.Number(2),
		},
		{
			name:     "positive number (16)",
			input:    datavalue.Number(16),
			expected: datavalue.Number(4),
		},
		{
			name:     "positive number (0)",
			input:    datavalue.Number(0),
			expected: datavalue.Number(0),
		},
		{
			name:     "negative number (-1)",
			input:    datavalue.Number(-1),
			expected: datavalue.Null(),
		},
	}

	functions := GetMathFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			sqrtFunc, hasFunction := functions["sqrt"]

			if !hasFunction {
				t.Fatalf("could not find sqrt function")
			}

			result, err := sqrtFunc.Handler(nil, []datavalue.Value{test.input})

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if result.Num != test.expected.Num {
				t.Fatalf("expected %f, got %f", test.expected.Num, result.Num)
			}
		})
	}
}
