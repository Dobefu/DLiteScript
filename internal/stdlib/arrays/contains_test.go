package arrays

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetContainsFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []datavalue.Value
		value    datavalue.Value
		expected bool
	}{
		{
			name: "contains number",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
			value:    datavalue.Number(2),
			expected: true,
		},
		{
			name: "does not contain number",
			arr: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
			},
			value:    datavalue.Number(4),
			expected: false,
		},
		{
			name: "contains string",
			arr: []datavalue.Value{
				datavalue.String("hello"),
				datavalue.String("world"),
			},
			value:    datavalue.String("hello"),
			expected: true,
		},
		{
			name: "does not contain string",
			arr: []datavalue.Value{
				datavalue.String("hello"),
				datavalue.String("world"),
			},
			value:    datavalue.String("hi"),
			expected: false,
		},
	}

	containsFunc := getContainsFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := containsFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.Array(test.arr...),
					test.value,
				},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			resultBool, _ := result.AsBool()

			if resultBool != test.expected {
				t.Fatalf("expected result to be %t, got %t", test.expected, resultBool)
			}
		})
	}
}
