package strings

import (
	"testing"
)

func TestGetStringsFunctions(t *testing.T) {
	t.Parallel()

	functions := GetStringsFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
