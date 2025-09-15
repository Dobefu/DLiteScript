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
	RunE:  runRootCmd,
}

// Execute executes the root command.
func Execute() int {
	err := rootCmd.Execute()

	if err != nil {
		return 1
	}

	return 0
}

func runRootCmd(_ *cobra.Command, args []string) error {
	runner := &scriptrunner.ScriptRunner{
		Args:    args,
		OutFile: os.Stdout,
	}

	err := runner.Run()

	if err != nil {
		slog.Error(err.Error())

		return fmt.Errorf("failed to run script: %s", err.Error())
	}

	return nil
}
