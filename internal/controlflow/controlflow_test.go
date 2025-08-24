package controlflow

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestControlFlow(t *testing.T) {
	t.Parallel()

	regularResult := NewRegularResult(datavalue.Number(1))

	if !regularResult.IsNormalResult() {
		t.Errorf("Expected normal result, got %v", regularResult)
	}

	if !regularResult.Value.Equals(datavalue.Number(1)) {
		t.Errorf("Expected regular result, got %v", regularResult)
	}

	breakResult := NewBreakResult(1)

	if !breakResult.IsBreakResult() {
		t.Errorf("Expected break result, got %v", breakResult)
	}

	continueResult := NewContinueResult(1)

	if !continueResult.IsContinueResult() {
		t.Errorf("Expected continue result, got %v", continueResult)
	}
}
