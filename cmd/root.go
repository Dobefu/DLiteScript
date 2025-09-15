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
	_ = rootCmd.Execute()

	return exitCode
}

func runRootCmd(_ *cobra.Command, args []string) {
	runner := &scriptrunner.ScriptRunner{
		Args:    args,
		OutFile: os.Stdout,
	}

	var err error
	exitCode, err = runner.Run()

	if err != nil {
		slog.Error(fmt.Sprintf("failed to run script: %s", err.Error()))
	}
}
