package cmd

import (
	"testing"
)

func TestLSPCmd(t *testing.T) {
	t.Parallel()

	runLSPCmd(lspCmd, nil)
}
