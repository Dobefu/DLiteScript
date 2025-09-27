package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSpliceFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		arr             []datavalue.Value
		start           int
		deleteCount     int
		items           []datavalue.Value
		expectedRemoved []datavalue.Value
		expectedResult  []datavalue.Value
	}{
		{
			name: "remove elements without insertion",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:       1,
			deleteCount: 2,
			items:       []datavalue.Value{},
			expectedRemoved: []datavalue.Value{
				datavalue.Number(2),
				datavalue.Number(3),
			},
			expectedResult: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(4),
				datavalue.Number(5),
			},
		},
		{
			name: "remove and insert elements",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:       1,
			deleteCount: 2,
			items: []datavalue.Value{
				datavalue.Number(6),
				datavalue.Number(7),
			},
			expectedRemoved: []datavalue.Value{
				datavalue.Number(2),
				datavalue.Number(3),
			},
			expectedResult: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(6),
				datavalue.Number(7),
				datavalue.Number(4),
				datavalue.Number(5),
			},
		},
		{
			name: "insert without removing",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
			start:       0,
			deleteCount: 0,
			items: []datavalue.Value{
				datavalue.Number(0),
			},
			expectedRemoved: []datavalue.Value{},
			expectedResult: []datavalue.Value{
				datavalue.Number(0),
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
		},
		{
			name: "remove negative index",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:       -1,
			deleteCount: 1,
			items:       []datavalue.Value{},
			expectedRemoved: []datavalue.Value{
				datavalue.Number(5),
			},
			expectedResult: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
			},
		},
		{
			name: "remove index greater than array length",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:           6,
			deleteCount:     0,
			items:           []datavalue.Value{},
			expectedRemoved: []datavalue.Value{},
			expectedResult: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
		},
		{
			name: "remove negative delete count",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:           0,
			deleteCount:     -1,
			items:           []datavalue.Value{},
			expectedRemoved: []datavalue.Value{},
			expectedResult: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
		},
		{
			name: "remove delete count greater than array length",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:       0,
			deleteCount: 6,
			items:       []datavalue.Value{},
			expectedRemoved: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			expectedResult: []datavalue.Value{},
		},
	}

	spliceFunc := getSpliceFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			args := make([]datavalue.Value, 0, 3+len(test.items))
			args = append(args, datavalue.Array(test.arr...))
			args = append(args, datavalue.Number(float64(test.start)))
			args = append(args, datavalue.Number(float64(test.deleteCount)))
			args = append(args, test.items...)

			result, err := spliceFunc.Handler(nil, args)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			tuple, _ := result.AsTuple()

			removedArr, _ := tuple[0].AsArray()
			resultArr, _ := tuple[1].AsArray()

			if len(removedArr) != len(test.expectedRemoved) {
				t.Fatalf(
					"expected removed array length %d, got %d",
					len(test.expectedRemoved),
					len(removedArr),
				)
			}

			if len(resultArr) != len(test.expectedResult) {
				t.Fatalf(
					"expected result array length %d, got %d",
					len(test.expectedResult),
					len(resultArr),
				)
			}
		})
	}
}
