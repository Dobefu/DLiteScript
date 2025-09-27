package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSplitFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     datavalue.Value
		delimiter datavalue.Value
		expected  datavalue.Value
	}{
		{
			name:      "comma separated values",
			input:     datavalue.String("a,b,c"),
			delimiter: datavalue.String(","),
			expected: datavalue.Array(
				datavalue.String("a"),
				datavalue.String("b"),
				datavalue.String("c"),
			),
		},
		{
			name:      "space separated words",
			input:     datavalue.String("a b c"),
			delimiter: datavalue.String(" "),
			expected: datavalue.Array(
				datavalue.String("a"),
				datavalue.String("b"),
				datavalue.String("c"),
			),
		},
		{
			name:      "no delimiter found",
			input:     datavalue.String("a b c"),
			delimiter: datavalue.String(","),
			expected: datavalue.Array(
				datavalue.String("a b c"),
			),
		},
	}

	splitFunc := getSplitFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := splitFunc.Handler(
				nil,
				[]datavalue.Value{test.input, test.delimiter},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedArray, _ := test.expected.AsArray()
			array, _ := result.AsArray()

			if len(array) != len(expectedArray) {
				t.Fatalf("expected length %d, got %d", len(expectedArray), len(array))
			}

			for i, val := range array {
				if val.ToString() != expectedArray[i].ToString() {
					t.Fatalf(
						"expected %s, got %s",
						expectedArray[i].ToString(),
						val.ToString(),
					)
				}
			}
		})
	}
}
