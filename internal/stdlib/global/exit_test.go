package global

import (
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestExit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []datavalue.Value
		expected byte
	}{
		{
			name:     "exit with 0",
			input:    []datavalue.Value{datavalue.Number(0)},
			expected: 0,
		},
	}

	functions := GetGlobalFunctions()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := &testEvaluator{buf: strings.Builder{}, exitCode: 0}
			exitFunc, hasExit := functions["exit"]

			if !hasExit {
				t.Fatalf("expected exit function, got nil")
			}

			_, err := exitFunc.Handler(evaluator, test.input)

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if evaluator.exitCode != test.expected {
				t.Errorf(
					"expected exit code to be %d, got %d",
					test.expected,
					evaluator.exitCode,
				)
			}
		})
	}
}
