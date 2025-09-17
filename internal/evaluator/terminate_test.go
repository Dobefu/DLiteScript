package evaluator

import (
	"io"
	"testing"
)

func TestTerminate(t *testing.T) {
	t.Parallel()

	evaluator := NewEvaluator(io.Discard)
	evaluator.Terminate(1)

	if evaluator.exitCode != 1 {
		t.Errorf("expected exit code to be 1, got %d", evaluator.exitCode)
	}

	if !evaluator.shouldTerminate {
		t.Errorf("expected should terminate to be true, got false")
	}
}
