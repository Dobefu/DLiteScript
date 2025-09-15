package cmd

import (
	"testing"
)

func TestLSPCmd(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	runLSPCmd(lspCmd, nil)

	if exitCode == 0 {
		t.Fatalf("Expected non-zero exit code, got %d", exitCode)
	}
	cmdMutex.Unlock()
}
