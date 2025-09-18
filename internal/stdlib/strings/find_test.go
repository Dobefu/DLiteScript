package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetFindFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		args     datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "empty search string",
			input:    datavalue.String("some test string"),
			args:     datavalue.String(""),
			expected: datavalue.Number(0),
		},
		{
			name:     "simple search string",
			input:    datavalue.String("some test string"),
			args:     datavalue.String("test"),
			expected: datavalue.Number(5),
		},
		{
			name:     "bogus search string",
			input:    datavalue.String("some test string"),
			args:     datavalue.String("bogus"),
			expected: datavalue.Number(-1),
		},
	}

	findFunc := getFindFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := findFunc.Handler(
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
