package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetPowFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		base     datavalue.Value
		exponent datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "2 ** 3",
			base:     datavalue.Number(2),
			exponent: datavalue.Number(3),
			expected: datavalue.Number(8),
		},
		{
			name:     "10 ** 2",
			base:     datavalue.Number(10),
			exponent: datavalue.Number(2),
			expected: datavalue.Number(100),
		},
		{
			name:     "4 ** 0.5",
			base:     datavalue.Number(4),
			exponent: datavalue.Number(0.5),
			expected: datavalue.Number(2),
		},
	}

	functions := GetMathFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			powFunc, hasFunction := functions["pow"]

			if !hasFunction {
				t.Fatalf("could not find pow function")
			}

			result, err := powFunc.Handler(
				nil,
				[]datavalue.Value{test.base, test.exponent},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if result.Num != test.expected.Num {
				t.Fatalf("expected %f, got %f", test.expected.Num, result.Num)
			}
		})
	}
}
