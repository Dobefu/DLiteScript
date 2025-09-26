package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/internal/repl"
	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{ //nolint:exhaustruct
	Use:   "repl",
	Short: "Start a DLiteScript REPL shell",
	Run:   runReplCmd,
}

func init() {
	rootCmd.AddCommand(replCmd)
}

func runReplCmd(_ *cobra.Command, _ []string) {
	var outfile io.Writer = os.Stdout
	var infile io.Reader = os.Stdin
	replInstance := repl.NewREPL(outfile, infile)
	err := replInstance.Run()

	if err != nil {
		slog.Error(fmt.Sprintf("failed to run REPL: %s", err.Error()))
		setExitCode(1)
	}
}
