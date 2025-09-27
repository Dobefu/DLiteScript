package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetLog10Function(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "log10(1)",
			input:    datavalue.Number(1),
			expected: datavalue.Number(0),
		},
		{
			name:     "log10(10)",
			input:    datavalue.Number(10),
			expected: datavalue.Number(math.Log10(10)),
		},
		{
			name:     "log10(100)",
			input:    datavalue.Number(100),
			expected: datavalue.Number(math.Log10(100)),
		},
	}

	log10Func := getLog10Function()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := log10Func.Handler(
				nil,
				[]datavalue.Value{test.input},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedNumber, _ := test.expected.AsNumber()
			number, err := result.AsNumber()

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if number != expectedNumber {
				t.Fatalf("expected %g, got %g", expectedNumber, number)
			}
		})
	}
}
