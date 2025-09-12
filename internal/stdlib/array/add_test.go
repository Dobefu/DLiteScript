package array

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
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

	add := getAddFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := add.Handler(
				nil,
				[]datavalue.Value{test.input, test.args},
			)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result.DataType() != datatype.DataTypeArray {
				t.Fatalf("expected array, got %v", result.DataType())
			}

			array, err := result.AsArray()

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if !result.Equals(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}

			expectedArray, err := test.expected.AsArray()

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if len(array) != len(expectedArray) {
				t.Fatalf("expected %d values, got %d", len(expectedArray), len(array))
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
