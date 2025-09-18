package errors

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetNewFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    datavalue.Value
		expected string
	}{
		{
			name:     "new error",
			input:    datavalue.String("test"),
			expected: "test",
		},
		{
			name:     "empty error",
			input:    datavalue.String(""),
			expected: "",
		},
		{
			name:     "error with spaces",
			input:    datavalue.String("Something went wrong"),
			expected: "Something went wrong",
		},
	}

	newFunc := getNewFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := newFunc.Handler(nil, []datavalue.Value{test.input})

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result.DataType() != datatype.DataTypeError {
				t.Fatalf("expected DataTypeError, got %v", result.DataType())
			}

			if result.ToString() != test.expected {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					result.ToString(),
				)
			}
		})
	}
}
