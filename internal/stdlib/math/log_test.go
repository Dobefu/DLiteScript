package math

import (
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetLogFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "log(1)",
			input:    datavalue.Number(1),
			expected: datavalue.Number(0),
		},
		{
			name:     "log(10)",
			input:    datavalue.Number(10),
			expected: datavalue.Number(math.Log(10)),
		},
	}

	logFunc := getLogFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := logFunc.Handler(
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
