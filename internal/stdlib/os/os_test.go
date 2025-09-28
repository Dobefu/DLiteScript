package os

import (
	"testing"
)

func TestGetOSFunctions(t *testing.T) {
	t.Parallel()

	functions := GetOSFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
