package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetIndexOfFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		substr   string
		expected int
	}{
		{
			name:     "Hello World contains Hello at index 0",
			str:      "Hello World",
			substr:   "Hello",
			expected: 0,
		},
		{
			name:     "Hello World contains World at index 6",
			str:      "Hello World",
			substr:   "World",
			expected: 6,
		},
		{
			name:     "Hello World does not contain xyz",
			str:      "Hello World",
			substr:   "xyz",
			expected: -1,
		},
		{
			name:     "Hello World contains empty string at index 0",
			str:      "Hello World",
			substr:   "",
			expected: 0,
		},
	}

	indexOfFunc := getIndexOfFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := indexOfFunc.Handler(
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
