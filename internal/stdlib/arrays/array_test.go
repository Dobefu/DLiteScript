package arrays

import (
	"testing"
)

func TestGetArrayFunctions(t *testing.T) {
	t.Parallel()

	functions := GetArrayFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
