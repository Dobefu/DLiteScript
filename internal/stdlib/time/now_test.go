package time

import (
	"testing"
	"time"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetNowFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []datavalue.Value
		expected datavalue.Value
	}{
		{
			name:     "now",
			input:    []datavalue.Value{},
			expected: datavalue.Number(float64(time.Now().Unix())),
		},
	}

	nowFunc := getNowFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := nowFunc.Handler(nil, test.input)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			expectedNumber, _ := test.expected.AsNumber()
			number, err := result.AsNumber()

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if number != expectedNumber {
				t.Fatalf(
					"expected %g, got: %g",
					expectedNumber,
					number,
				)
			}
		})
	}
}
