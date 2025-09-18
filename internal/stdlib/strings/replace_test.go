package strings

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetReplaceFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []datavalue.Value
		expected datavalue.Value
	}{
		{
			name: "empty search string",
			input: []datavalue.Value{
				datavalue.String("some test string"),
				datavalue.String("test"),
				datavalue.String("new"),
			},
			expected: datavalue.String("some new string"),
		},
		{
			name: "bogus search string",
			input: []datavalue.Value{
				datavalue.String("some test string"),
				datavalue.String("bogus"),
				datavalue.String("new"),
			},
			expected: datavalue.String("some test string"),
		},
	}

	replaceFunc := getReplaceFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := replaceFunc.Handler(nil, test.input)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedString, _ := test.expected.AsString()
			str, err := result.AsString()

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if str != expectedString {
				t.Fatalf(
					"expected %s, got: %s",
					expectedString,
					str,
				)
			}
		})
	}
}
