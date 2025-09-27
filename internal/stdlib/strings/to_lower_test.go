package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetToLowerFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "lowercase string",
			input:    datavalue.String("TEST"),
			expected: datavalue.String("test"),
		},
		{
			name:     "mixed case string with numbers",
			input:    datavalue.String("TEST 123"),
			expected: datavalue.String("test 123"),
		},
	}

	toLowerFunc := getToLowerFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := toLowerFunc.Handler(
				nil,
				[]datavalue.Value{test.input},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedStr, _ := test.expected.AsString()
			actualStr, _ := result.AsString()

			if actualStr != expectedStr {
				t.Fatalf("expected \"%s\", got \"%s\"", expectedStr, actualStr)
			}
		})
	}
}
