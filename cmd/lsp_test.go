package cmd

import (
	"testing"
)

func TestLSPCmd(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	defer func() {
		resetExitCode()
		cmdMutex.Unlock()
	}()

	runLSPCmd(lspCmd, nil)

	if getExitCode() == 0 {
		t.Fatalf("Expected non-zero exit code, got %d", getExitCode())
	}
}
