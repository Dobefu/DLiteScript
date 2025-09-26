package math

import (
	"errors"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetModFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		base     datavalue.Value
		exponent datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "2 % 3",
			base:     datavalue.Number(2),
			exponent: datavalue.Number(3),
			expected: datavalue.Tuple(datavalue.Number(2), datavalue.Null()),
		},
		{
			name:     "10 % 2",
			base:     datavalue.Number(10),
			exponent: datavalue.Number(2),
			expected: datavalue.Tuple(datavalue.Number(0), datavalue.Null()),
		},
		{
			name:     "4 % 0.5",
			base:     datavalue.Number(4),
			exponent: datavalue.Number(0.5),
			expected: datavalue.Tuple(datavalue.Number(0), datavalue.Null()),
		},
		{
			name:     "0 % 0",
			base:     datavalue.Number(0),
			exponent: datavalue.Number(0),
			expected: datavalue.Tuple(
				datavalue.Number(0),
				datavalue.Error(errors.New("cannot mod by zero")),
			),
		},
	}

	functions := GetMathFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			modFunc, hasFunction := functions["mod"]

			if !hasFunction {
				t.Fatalf("could not find mod function")
			}

			result, err := modFunc.Handler(
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
