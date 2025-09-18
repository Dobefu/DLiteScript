package errors

import (
	"testing"
)

func TestGetErrorFunctions(t *testing.T) {
	t.Parallel()

	functions := GetErrorFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
