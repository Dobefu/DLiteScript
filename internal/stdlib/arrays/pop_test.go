package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetPopFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		arr             []datavalue.Value
		expectedElement datavalue.Value
		expectedArr     []datavalue.Value
	}{
		{
			name: "pop from array of numbers",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
			},
			expectedElement: datavalue.Number(4),
			expectedArr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
		},
		{
			name: "pop from array of strings",
			arr: []datavalue.Value{
				datavalue.String("hello"),
				datavalue.String("world"),
			},
			expectedElement: datavalue.String("world"),
			expectedArr:     []datavalue.Value{datavalue.String("hello")},
		},
		{
			name:            "pop from empty array",
			arr:             []datavalue.Value{},
			expectedElement: datavalue.Null(),
			expectedArr:     []datavalue.Value{},
		},
	}

	popFunc := getPopFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := popFunc.Handler(
				nil,
				[]datavalue.Value{datavalue.Array(test.arr...)},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			tuple, _ := result.AsTuple()
			resultArr, _ := tuple[1].AsArray()

			if len(resultArr) != len(test.expectedArr) {
				t.Fatalf(
					"expected array length %d, got %d",
					len(test.expectedArr),
					len(resultArr),
				)
			}
		})
	}
}
