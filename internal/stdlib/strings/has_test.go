package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetHasFunction(t *testing.T) {
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
			expected: datavalue.Bool(true),
		},
		{
			name:     "simple search string",
			input:    datavalue.String("some test string"),
			args:     datavalue.String("test"),
			expected: datavalue.Bool(true),
		},
		{
			name:     "bogus search string",
			input:    datavalue.String("some test string"),
			args:     datavalue.String("bogus"),
			expected: datavalue.Bool(false),
		},
	}

	hasFunc := getHasFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := hasFunc.Handler(
				nil,
				[]datavalue.Value{test.input, test.args},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedBool, _ := test.expected.AsBool()
			boolean, err := result.AsBool()

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if boolean != expectedBool {
				t.Fatalf(
					"expected \"%t\", got: \"%t\"",
					expectedBool,
					boolean,
				)
			}
		})
	}
}
