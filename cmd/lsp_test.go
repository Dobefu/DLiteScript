package cmd

import (
	"testing"
)

func TestLSPCmd(t *testing.T) {
	t.Parallel()

	err := runLSPCmd(lspCmd, nil)

	if err != nil {
		t.Fatalf("Failed to run LSP command: %s", err.Error())
	}
}
