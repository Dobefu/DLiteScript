package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetJoinFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     datavalue.Value
		delimiter datavalue.Value
		expected  datavalue.Value
	}{
		{
			name: "comma separated values",
			input: datavalue.Array(
				datavalue.String("a"),
				datavalue.String("b"),
				datavalue.String("c"),
			),
			delimiter: datavalue.String(","),
			expected:  datavalue.String("a,b,c"),
		},
		{
			name: "space separated values",
			input: datavalue.Array(
				datavalue.String("a"),
				datavalue.String("b"),
				datavalue.String("c"),
			),
			delimiter: datavalue.String(" "),
			expected:  datavalue.String("a b c"),
		},
		{
			name: "single element",
			input: datavalue.Array(
				datavalue.String("a"),
			),
			delimiter: datavalue.String(","),
			expected:  datavalue.String("a"),
		},
	}

	joinFunc := getJoinFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := joinFunc.Handler(
				nil,
				[]datavalue.Value{test.input, test.delimiter},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedStr, _ := test.expected.AsString()
			actualStr, _ := result.AsString()

			if actualStr != expectedStr {
				t.Fatalf("expected \"%s\", got \"%s\"", expectedStr, actualStr)
			}
		})
	}
}
