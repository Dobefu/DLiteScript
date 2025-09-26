package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetReverseFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name: "numbers array",
			input: datavalue.Array(
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			),
			expected: datavalue.Array(
				datavalue.Number(3),
				datavalue.Number(2),
				datavalue.Number(1),
			),
		},
		{
			name: "strings array",
			input: datavalue.Array(
				datavalue.String("a"),
				datavalue.String("b"),
				datavalue.String("c"),
			),
			expected: datavalue.Array(
				datavalue.String("c"),
				datavalue.String("b"),
				datavalue.String("a"),
			),
		},
	}

	reverseFunc := getReverseFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := reverseFunc.Handler(
				nil,
				[]datavalue.Value{test.input},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedArray, _ := test.expected.AsArray()
			array, err := result.AsArray()

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			for i, val := range array {
				if val.ToString() != expectedArray[i].ToString() {
					t.Fatalf(
						"expected %v, got %v",
						expectedArray[i].ToString(),
						val.ToString(),
					)
				}
			}
		})
	}
}
