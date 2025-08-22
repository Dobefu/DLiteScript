package datatype

import "testing"

func TestDatatype(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input         DataType
		expectedValue string
	}{
		{
			input:         DataTypeNull,
			expectedValue: "null",
		},
		{
			input:         DataTypeNumber,
			expectedValue: "number",
		},
		{
			input:         DataTypeString,
			expectedValue: "string",
		},
		{
			input:         DataTypeBool,
			expectedValue: "bool",
		},
		{
			input:         DataType(-1),
			expectedValue: "unknown",
		},
	}

	for _, test := range tests {
		if test.input.AsString() != test.expectedValue {
			t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.AsString())
		}
	}
}
