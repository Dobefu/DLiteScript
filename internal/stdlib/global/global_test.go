package global

import (
	"fmt"
	"strings"
	"testing"
)

type testEvaluator struct {
	buf strings.Builder
}

func (e *testEvaluator) AddToBuffer(format string, args ...any) {
	fmt.Fprintf(&e.buf, format, args...)
}

func TestGetGlobalFunctions(t *testing.T) {
	t.Parallel()

	functions := GetGlobalFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
