package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetAbsFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "negative number",
			input:    datavalue.Number(-1),
			expected: datavalue.Number(1),
		},
		{
			name:     "positive number",
			input:    datavalue.Number(1),
			expected: datavalue.Number(1),
		},
		{
			name:     "zero",
			input:    datavalue.Number(0),
			expected: datavalue.Number(0),
		},
		{
			name:     "negative zero",
			input:    datavalue.Number(-0),
			expected: datavalue.Number(0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			functions := GetMathFunctions()
			absFunc, hasFunction := functions["abs"]

			if !hasFunction {
				t.Fatalf("could not find abs function")
			}

			result, err := absFunc.Handler(nil, []datavalue.Value{test.input})

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if result.Num != test.expected.Num {
				t.Fatalf("expected %f, got %f", test.expected.Num, result.Num)
			}
		})
	}
}
