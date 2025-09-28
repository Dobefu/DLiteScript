package os

import (
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetSetEnvVariableFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		envName  string
		envValue string
	}{
		{
			name:     "set VARIABLE environment variable",
			envName:  "VARIABLE",
			envValue: "value",
		},
	}

	setEnvVariableFunc := getSetEnvVariableFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := setEnvVariableFunc.Handler(
				nil,
				[]datavalue.Value{
					datavalue.String(test.envName),
					datavalue.String(test.envValue),
				},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			if result.Error != nil {
				t.Fatalf("expected no error, got: \"%s\"", result.Error.Error())
			}

			actualValue := os.Getenv(test.envName)

			if actualValue != test.envValue {
				t.Fatalf(
					"expected env var %s to be \"%s\", got \"%s\"",
					test.envName,
					test.envValue,
					actualValue,
				)
			}
		})
	}
}
