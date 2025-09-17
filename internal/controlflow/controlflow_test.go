package controlflow

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestControlFlowRegularResult(t *testing.T) {
	t.Parallel()

	regularResult := NewRegularResult(datavalue.Number(1))

	if !regularResult.IsNormalResult() {
		t.Errorf("Expected normal result, got %v", regularResult)
	}

	if !regularResult.Value.Equals(datavalue.Number(1)) {
		t.Errorf("Expected regular result, got %v", regularResult)
	}
}

func TestControlFlowBreakResult(t *testing.T) {
	t.Parallel()

	breakResult := NewBreakResult(1)

	if !breakResult.IsBreakResult() {
		t.Errorf("Expected break result, got %v", breakResult)
	}
}

func TestControlFlowContinueResult(t *testing.T) {
	t.Parallel()

	continueResult := NewContinueResult(1)

	if !continueResult.IsContinueResult() {
		t.Errorf("Expected continue result, got %v", continueResult)
	}
}

func TestControlFlowReturnResult(t *testing.T) {
	t.Parallel()

	returnResult := NewReturnResult(datavalue.Number(1))

	if !returnResult.IsReturnResult() {
		t.Errorf("Expected return result, got %v", returnResult)
	}

	if !returnResult.Value.Equals(datavalue.Number(1)) {
		t.Errorf("Expected return result, got %v", returnResult)
	}
}

func TestControlFlowExitResult(t *testing.T) {
	t.Parallel()

	exitResult := NewExitResult(0)

	if !exitResult.IsExitResult() {
		t.Errorf("Expected exit result, got %v", exitResult)
	}

	if exitResult.Control.Count != 0 {
		t.Errorf("Expected exit result, got %v", exitResult)
	}
}
