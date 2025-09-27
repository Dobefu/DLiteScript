package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetTrimRightFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "spaces on both sides",
			input:    datavalue.String(" test "),
			expected: datavalue.String(" test"),
		},
		{
			name:     "spaces on the left",
			input:    datavalue.String(" test"),
			expected: datavalue.String(" test"),
		},
		{
			name:     "spaces on the right",
			input:    datavalue.String("test "),
			expected: datavalue.String("test"),
		},
		{
			name:     "no spaces",
			input:    datavalue.String("hello"),
			expected: datavalue.String("hello"),
		},
		{
			name:     "only spaces",
			input:    datavalue.String("   "),
			expected: datavalue.String(""),
		},
	}

	trimRightFunc := getTrimRightFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := trimRightFunc.Handler(
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
