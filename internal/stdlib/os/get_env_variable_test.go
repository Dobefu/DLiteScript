package os

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetGetEnvVariableFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		envName     string
		expectEmpty bool
	}{
		{
			name:        "get PATH environment variable",
			envName:     "PATH",
			expectEmpty: false,
		},
		{
			name:        "get HOME environment variable",
			envName:     "HOME",
			expectEmpty: false,
		},
		{
			name:        "get non-existent environment variable",
			envName:     "BOGUS",
			expectEmpty: true,
		},
	}

	getEnvVariableFunc := getGetEnvVariableFunction()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := getEnvVariableFunc.Handler(
				nil,
				[]datavalue.Value{datavalue.String(test.envName)},
			)

			if err != nil {
				t.Fatalf("expected no error from handler, got: \"%s\"", err.Error())
			}

			resultStr, err := result.AsString()

			if err != nil {
				t.Fatalf("expected string result, got error: \"%s\"", err.Error())
			}

			if test.expectEmpty && resultStr != "" {
				t.Fatalf("expected empty string, got: \"%s\"", resultStr)
			}

			if !test.expectEmpty && resultStr == "" {
				t.Fatalf("expected non-empty string, got empty string")
			}
		})
	}
}
