package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSortStringsFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []datavalue.Value
		expected []datavalue.Value
	}{
		{
			name: "sort strings",
			arr: []datavalue.Value{
				datavalue.String("c"),
				datavalue.String("b"),
				datavalue.String("a"),
			},
			expected: []datavalue.Value{
				datavalue.String("a"),
				datavalue.String("b"),
				datavalue.String("c"),
			},
		},
	}

	sortStringsFunc := getSortStringsFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := sortStringsFunc.Handler(
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

func TestGetSortStringsFunctionErr(t *testing.T) {
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
			expected: "array contains non-strings",
		},
	}

	sortStringsFunc := getSortStringsFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := sortStringsFunc.Handler(
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
