package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetFilterFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		arr                 []datavalue.Value
		expectedFilteredOut []datavalue.Value
		expectedRemaining   []datavalue.Value
	}{
		{
			name: "filter mixed falsy values",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(0),
				datavalue.Number(2),
				datavalue.Bool(false),
				datavalue.Number(3),
				datavalue.String(""),
				datavalue.Number(4),
			},
			expectedFilteredOut: []datavalue.Value{
				datavalue.Number(0),
				datavalue.Bool(false),
				datavalue.String(""),
			},
			expectedRemaining: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
			},
		},
		{
			name: "filter boolean and string falsy values",
			arr: []datavalue.Value{
				datavalue.Bool(true),
				datavalue.Bool(false),
				datavalue.String("hello"),
				datavalue.String(""),
			},
			expectedFilteredOut: []datavalue.Value{
				datavalue.Bool(false),
				datavalue.String(""),
			},
			expectedRemaining: []datavalue.Value{
				datavalue.Bool(true),
				datavalue.String("hello"),
			},
		},
		{
			name: "filter with no falsy values",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
			expectedFilteredOut: []datavalue.Value{},
			expectedRemaining: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
		},
		{
			name:                "filter empty array",
			arr:                 []datavalue.Value{},
			expectedFilteredOut: []datavalue.Value{},
			expectedRemaining:   []datavalue.Value{},
		},
	}

	filterFunc := getFilterFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := filterFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.Array(test.arr...),
				},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			tuple, _ := result.AsTuple()
			filteredOutArr, _ := tuple[0].AsArray()
			remainingArr, _ := tuple[1].AsArray()

			if len(filteredOutArr) != len(test.expectedFilteredOut) {
				t.Fatalf(
					"expected filtered out array length to be %d, got %d",
					len(test.expectedFilteredOut),
					len(filteredOutArr),
				)
			}

			if len(remainingArr) != len(test.expectedRemaining) {
				t.Fatalf(
					"expected remaining array length to be %d, got %d",
					len(test.expectedRemaining),
					len(remainingArr),
				)
			}
		})
	}
}
