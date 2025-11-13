package io

import (
	"testing"
)

func TestGetIOFunctions(t *testing.T) {
	t.Parallel()

	functions := GetIOFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
