package stdlib

import "testing"

func TestStdlib(t *testing.T) {
	t.Parallel()

	functions := GetFunctionRegistry()

	if len(functions) <= 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
