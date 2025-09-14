package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetTanFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "positive number (1.5)",
			input:    datavalue.Number(1.5),
			expected: datavalue.Number(math.Tan(1.5)),
		},
		{
			name:     "positive number (1)",
			input:    datavalue.Number(1),
			expected: datavalue.Number(math.Tan(1)),
		},
		{
			name:     "positive number (0)",
			input:    datavalue.Number(0),
			expected: datavalue.Number(math.Tan(0)),
		},
		{
			name:     "negative number (-1.5)",
			input:    datavalue.Number(-1.5),
			expected: datavalue.Number(math.Tan(-1.5)),
		},
		{
			name:     "negative number (-1)",
			input:    datavalue.Number(-1),
			expected: datavalue.Number(math.Tan(-1)),
		},
		{
			name:     "negative number (-0)",
			input:    datavalue.Number(-0),
			expected: datavalue.Number(math.Tan(-0)),
		},
	}

	functions := GetMathFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tanFunc, hasFunction := functions["tan"]

			if !hasFunction {
				t.Fatalf("could not find tan function")
			}

			result, err := tanFunc.Handler(nil, []datavalue.Value{test.input})

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if result.Num != test.expected.Num {
				t.Fatalf("expected %f, got %f", test.expected.Num, result.Num)
			}
		})
	}
}
