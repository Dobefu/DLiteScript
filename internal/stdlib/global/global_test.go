package global

import (
	"fmt"
	"strings"
	"testing"
)

type testEvaluator struct {
	buf      strings.Builder
	exitCode byte
}

func (e *testEvaluator) AddToBuffer(format string, args ...any) {
	fmt.Fprintf(&e.buf, format, args...)
}

func (e *testEvaluator) Terminate(code byte) {
	e.exitCode = code
}

func TestGetGlobalFunctions(t *testing.T) {
	t.Parallel()

	functions := GetGlobalFunctions()

	if len(functions) == 0 {
		t.Fatalf("expected at least 1 function, got %d", len(functions))
	}
}
