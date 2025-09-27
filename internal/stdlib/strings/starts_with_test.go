package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetStartsWithFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		prefix   string
		expected bool
	}{
		{
			name:     "Hello World starts with Hello",
			str:      "Hello World",
			prefix:   "Hello",
			expected: true,
		},
		{
			name:     "Hello World does not start with World",
			str:      "Hello World",
			prefix:   "World",
			expected: false,
		},
		{
			name:     "Hello World starts with empty string",
			str:      "Hello World",
			prefix:   "",
			expected: true,
		},
		{
			name:     "empty string does not start with Hello",
			str:      "",
			prefix:   "Hello",
			expected: false,
		},
	}

	startsWithFunc := getStartsWithFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := startsWithFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.String(test.str),
					datavalue.String(test.prefix),
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
