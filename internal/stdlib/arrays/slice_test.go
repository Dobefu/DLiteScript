package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSliceFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []datavalue.Value
		start    int
		end      int
		expected []datavalue.Value
	}{
		{
			name: "slice middle portion",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start: 1,
			end:   4,
			expected: []datavalue.Value{
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
			},
		},
		{
			name: "slice from beginning",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start: 0,
			end:   3,
			expected: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
		},
		{
			name: "slice with same start and end",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:    2,
			end:      2,
			expected: []datavalue.Value{},
		},
		{
			name: "slice with start and end both 0",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:    0,
			end:      0,
			expected: []datavalue.Value{},
		},
		{
			name: "slice with negative start",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:    -1,
			end:      0,
			expected: []datavalue.Value{},
		},
		{
			name: "slice with negative end",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start: 0,
			end:   -1,
			expected: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
			},
		},
		{
			name: "slice with start greater than array length",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start:    6,
			end:      1,
			expected: []datavalue.Value{},
		},
		{
			name: "slice with end greater than array length",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
			start: 0,
			end:   6,
			expected: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
				datavalue.Number(5),
			},
		},
	}

	sliceFunc := getSliceFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := sliceFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.Array(test.arr...),
					datavalue.Number(float64(test.start)),
					datavalue.Number(float64(test.end)),
				},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			resultArr, _ := result.AsArray()

			if len(resultArr) != len(test.expected) {
				t.Fatalf(
					"expected array length %d, got %d",
					len(test.expected),
					len(resultArr),
				)
			}

			for i, expected := range test.expected {
				if resultArr[i].ToString() != expected.ToString() {
					t.Fatalf(
						"expected element %d to be \"%s\", got \"%s\"",
						i, expected.ToString(), resultArr[i].ToString(),
					)
				}
			}
		})
	}
}
