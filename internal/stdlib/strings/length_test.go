package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetLengthFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		args     datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "empty string",
			input:    datavalue.String(""),
			args:     datavalue.Number(1),
			expected: datavalue.Number(0),
		},
		{
			name:     "simple string",
			input:    datavalue.String("test"),
			args:     datavalue.Number(1),
			expected: datavalue.Number(4),
		},
	}

	lengthFunc := getLengthFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := lengthFunc.Handler(
				nil,
				[]datavalue.Value{test.input, test.args},
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
					"expected %g, got: %g",
					expectedNumber,
					number,
				)
			}
		})
	}
}
