package cmd

import (
	"testing"
)

func TestReplCmd(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	defer func() {
		resetExitCode()
		cmdMutex.Unlock()
	}()

	runReplCmd(replCmd, nil)

	if exitCode() != 0 {
		t.Fatalf("Expected exit code 0, got %d", exitCode())
	}
}
