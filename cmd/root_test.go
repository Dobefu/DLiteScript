package cmd

import "testing"

func TestRootCmd(t *testing.T) {
	t.Parallel()

	runRootCmd(rootCmd, []string{"../examples/00_simple/main.dl"})
}
