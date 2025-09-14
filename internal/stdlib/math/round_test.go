package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetRoundFunction(t *testing.T) {
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
			expected: datavalue.Number(1),
		},
		{
			name:     "positive number (1)",
			input:    datavalue.Number(1),
			expected: datavalue.Number(1),
		},
		{
			name:     "negative number (-1.5)",
			input:    datavalue.Number(-1.5),
			expected: datavalue.Number(-2),
		},
	}

	functions := GetMathFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			roundFunc, hasFunction := functions["round"]

			if !hasFunction {
				t.Fatalf("could not find round function")
			}

			result, err := roundFunc.Handler(nil, []datavalue.Value{test.input})

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if result.Num != test.expected.Num {
				t.Fatalf("expected %f, got %f", test.expected.Num, result.Num)
			}
		})
	}
}
