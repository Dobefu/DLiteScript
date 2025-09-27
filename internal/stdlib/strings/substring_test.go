package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSubstringFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		start    float64
		length   float64
		expected string
	}{
		{
			name:     "valid substring",
			str:      "Hello World",
			start:    0,
			length:   5,
			expected: "Hello",
		},
		{
			name:     "valid substring to end",
			str:      "Hello World",
			start:    6,
			length:   5,
			expected: "World",
		},
		{
			name:     "valid substring from end",
			str:      "Hello World",
			start:    -5,
			length:   3,
			expected: "Wor",
		},
		{
			name:     "valid substring from start",
			str:      "Hello World",
			start:    0,
			length:   -6,
			expected: "Hello",
		},
		{
			name:     "valid substring from start and from end",
			str:      "Hello World",
			start:    -5,
			length:   -1,
			expected: "Worl",
		},
	}

	substringFunc := getSubstringFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := substringFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.String(test.str),
					datavalue.Number(test.start),
					datavalue.Number(test.length),
				},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			tuple, _ := result.AsTuple()

			if tuple[1].DataType != datatype.DataTypeNull {
				errVal, _ := tuple[1].AsError()

				t.Fatalf("expected no error, got: \"%s\"", errVal.Error())
			}

			resultStr, _ := tuple[0].AsString()

			if resultStr != test.expected {
				t.Fatalf(
					"expected result to be \"%s\", got \"%s\"",
					test.expected,
					resultStr,
				)
			}
		})
	}
}

func TestGetSubstringFunctionErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		start    float64
		length   float64
		expected string
	}{
		{
			name:     "start greater than string length",
			str:      "Hello World",
			start:    20,
			length:   5,
			expected: "start index out of bounds: 20 >= 11",
		},
		{
			name:     "negative length greater than string length",
			str:      "Hello World",
			start:    0,
			length:   -20,
			expected: "negative length results in empty string: length -20",
		},
		{
			name:     "invalid length",
			str:      "Hello World",
			start:    0,
			length:   20,
			expected: "length exceeds string bounds: requested 20, available 11",
		},
	}

	substringFunc := getSubstringFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := substringFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.String(test.str),
					datavalue.Number(test.start),
					datavalue.Number(test.length),
				},
			)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			tuple, _ := result.AsTuple()

			if tuple[1].DataType == datatype.DataTypeNull {
				t.Fatalf("expected error, got null")
			}

			errVal, _ := tuple[1].AsError()

			if errVal.Error() != test.expected {
				t.Fatalf(
					"expected error to be \"%s\", got \"%s\"",
					test.expected,
					errVal.Error(),
				)
			}
		})
	}
}
