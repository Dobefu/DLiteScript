package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetEndsWithFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		suffix   string
		expected bool
	}{
		{
			name:     "Hello World ends with World",
			str:      "Hello World",
			suffix:   "World",
			expected: true,
		},
		{
			name:     "Hello World does not end with Hello",
			str:      "Hello World",
			suffix:   "Hello",
			expected: false,
		},
		{
			name:     "Hello World ends with empty string",
			str:      "Hello World",
			suffix:   "",
			expected: true,
		},
		{
			name:     "empty string does not end with World",
			str:      "",
			suffix:   "World",
			expected: false,
		},
	}

	endsWithFunc := getEndsWithFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := endsWithFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.String(test.str),
					datavalue.String(test.suffix),
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
