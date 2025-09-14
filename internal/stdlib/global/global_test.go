package global

import (
	"testing"
)

func TestGetGlobalFunctions(t *testing.T) {
	t.Parallel()

	functions := GetGlobalFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
