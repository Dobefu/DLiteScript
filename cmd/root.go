package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/internal/scriptrunner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //nolint:exhaustruct
	Use:     "DLiteScript",
	Aliases: []string{"run"},
	Args: cobra.PositionalArgs(func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("usage: %s <file>", cmd.CommandPath())
		}

		return nil
	}),
	Short: "A delightfully simple scripting language",
	Run:   runRootCmd,
}

// Execute executes the root command.
func Execute() int {
	err := rootCmd.Execute()

	if err != nil {
		exitCode = 1
	}

	return exitCode
}

func runRootCmd(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		slog.Error("no file specified")
		exitCode = 1

		return
	}

	runner := &scriptrunner.ScriptRunner{
		OutFile: os.Stdout,
	}

	var err error
	exitCode, err = runner.RunScript(args[0])

	if err != nil {
		slog.Error(fmt.Sprintf("failed to run script: %s", err.Error()))
		exitCode = 1
	}
}
