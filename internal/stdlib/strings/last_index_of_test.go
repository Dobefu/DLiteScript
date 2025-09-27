package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetLastIndexOfFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		substr   string
		expected int
	}{
		{
			name:     "Hello World Hello contains Hello at last index 12",
			str:      "Hello World Hello",
			substr:   "Hello",
			expected: 12,
		},
		{
			name:     "Hello World Hello contains World at index 6",
			str:      "Hello World Hello",
			substr:   "World",
			expected: 6,
		},
		{
			name:     "Hello World Hello does not contain xyz",
			str:      "Hello World Hello",
			substr:   "xyz",
			expected: -1,
		},
		{
			name:     "Hello World Hello contains empty string at last index 17",
			str:      "Hello World Hello",
			substr:   "",
			expected: 17,
		},
	}

	lastIndexOfFunc := getLastIndexOfFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := lastIndexOfFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.String(test.str),
					datavalue.String(test.substr),
				},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			resultNum, _ := result.AsNumber()

			if int(resultNum) != test.expected {
				t.Fatalf(
					"expected result to be %d, got %d",
					test.expected,
					int(resultNum),
				)
			}
		})
	}
}
