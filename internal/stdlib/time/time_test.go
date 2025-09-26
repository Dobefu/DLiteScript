package time

import (
	"testing"
)

func TestGetTimeFunctions(t *testing.T) {
	t.Parallel()

	functions := GetTimeFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
