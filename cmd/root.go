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
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func runRootCmd(_ *cobra.Command, args []string) {
	runner := &scriptrunner.ScriptRunner{
		Args:    args,
		OutFile: os.Stdout,
		OnError: func(err error) {
			slog.Error(err.Error())
			os.Exit(1)
		},
	}

	err := runner.Run()

	if err != nil {
		runner.OnError(err)
	}
}
