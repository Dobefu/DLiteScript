package math

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSignFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "-5",
			input:    datavalue.Number(-5),
			expected: datavalue.Number(-1),
		},
		{
			name:     "0",
			input:    datavalue.Number(0),
			expected: datavalue.Number(0),
		},
		{
			name:     "0.1",
			input:    datavalue.Number(.1),
			expected: datavalue.Number(1),
		},
		{
			name:     "-0",
			input:    datavalue.Number(-0),
			expected: datavalue.Number(0),
		},
	}

	signFunc := getSignFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := signFunc.Handler(
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
				t.Fatalf(
					"expected %g, got %g",
					expectedNumber,
					number,
				)
			}
		})
	}
}
