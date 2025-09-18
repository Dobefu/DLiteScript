package datatype

import (
	"testing"
)

func TestDatatype(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         DataType
		expectedValue string
	}{
		{
			name:          "null",
			input:         DataTypeNull,
			expectedValue: "null",
		},
		{
			name:          "number",
			input:         DataTypeNumber,
			expectedValue: "number",
		},
		{
			name:          "string",
			input:         DataTypeString,
			expectedValue: "string",
		},
		{
			name:          "bool",
			input:         DataTypeBool,
			expectedValue: "bool",
		},
		{
			name:          "function",
			input:         DataTypeFunction,
			expectedValue: "function",
		},
		{
			name:          "tuple",
			input:         DataTypeTuple,
			expectedValue: "tuple",
		},
		{
			name:          "array",
			input:         DataTypeArray,
			expectedValue: "array",
		},
		{
			name:          "error",
			input:         DataTypeError,
			expectedValue: "error",
		},
		{
			name:          "any",
			input:         DataTypeAny,
			expectedValue: "any",
		},
		{
			name:          "unknown",
			input:         DataType(-1),
			expectedValue: "unknown",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.AsString() != test.expectedValue {
				t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.AsString())
			}
		})
	}
}
