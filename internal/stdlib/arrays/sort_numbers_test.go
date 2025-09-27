package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSortNumbersFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []datavalue.Value
		expected []datavalue.Value
		hasError bool
	}{
		{
			name: "sort numbers",
			arr: []datavalue.Value{
				datavalue.Number(3),
				datavalue.Number(1),
				datavalue.Number(4),
				datavalue.Number(1),
				datavalue.Number(5),
			},
			expected: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(1),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			hasError: false,
		},
		{
			name: "sort numbers with multi-digit",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(10),
			},
			expected: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(10),
			},
			hasError: false,
		},
		{
			name: "mixed types returns error",
			arr: []datavalue.Value{
				datavalue.Number(3),
				datavalue.String("hello"),
				datavalue.Number(1),
			},
			expected: nil,
			hasError: true,
		},
	}

	sortNumbersFunc := getSortNumbersFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := sortNumbersFunc.Handler(
				nil,
				[]datavalue.Value{datavalue.Array(test.arr...)},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			tuple, _ := result.AsTuple()
			resultArr, _ := tuple[0].AsArray()

			if len(resultArr) != len(test.expected) {
				t.Fatalf(
					"expected array length to be %d, got %d",
					len(test.expected),
					len(resultArr),
				)
			}

			for i, expectedVal := range test.expected {
				if resultArr[i].ToString() != expectedVal.ToString() {
					t.Fatalf(
						"expected element %d to be \"%s\", got \"%s\"",
						i, expectedVal.ToString(), resultArr[i].ToString(),
					)
				}
			}
		})
	}
}

func TestGetSortNumbersFunctionErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []datavalue.Value
		expected string
	}{
		{
			name: "mixed types returns error",
			arr: []datavalue.Value{
				datavalue.Number(3),
				datavalue.String("hello"),
				datavalue.Number(1),
			},
			expected: "array contains non-numbers",
		},
	}

	sortNumbersFunc := getSortNumbersFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := sortNumbersFunc.Handler(
				nil,
				[]datavalue.Value{datavalue.Array(test.arr...)},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			tuple, _ := result.AsTuple()

			if tuple[1].DataType != datatype.DataTypeError {
				t.Fatalf("expected error, got: %v", tuple[1])
			}
		})
	}
}
