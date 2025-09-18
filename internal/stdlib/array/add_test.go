package array

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetAddFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		args     datavalue.Value
		expected datavalue.Value
	}{
		{
			name: "number argument",
			input: datavalue.Array(
				datavalue.Number(1),
			),
			args: datavalue.Number(2),
			expected: datavalue.Array(
				datavalue.Number(1),
				datavalue.Number(2),
			),
		},
		{
			name: "array argument",
			input: datavalue.Array(
				datavalue.Number(1),
			),
			args: datavalue.Array(
				datavalue.Number(2),
				datavalue.Number(3),
			),
			expected: datavalue.Array(
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			),
		},
	}

	addFunc := getAddFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := addFunc.Handler(
				nil,
				[]datavalue.Value{test.input, test.args},
			)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			expectedArray, _ := test.expected.AsArray()
			array, err := result.AsArray()

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
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
