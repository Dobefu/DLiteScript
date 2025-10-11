package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/Dobefu/DLiteScript/scriptrunner"
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

func init() {
	rootCmd.Flags().BoolP("quiet", "q", false, "Don't print any messages to the output")
}

// Execute executes the root command.
func Execute() byte {
	err := rootCmd.Execute()

	if err != nil {
		setExitCode(1)
	}

	return getExitCode()
}

func runRootCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		slog.Error("no file specified")
		setExitCode(1)

		return
	}

	isQuiet, _ := cmd.Flags().GetBool("quiet")
	var outfile io.Writer = os.Stdout

	if isQuiet {
		outfile = io.Discard
	}

	runner := &scriptrunner.ScriptRunner{
		OutFile: outfile,
	}

	var err error
	code, err := runner.RunScript(args[0])
	setExitCode(code)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to run script: %s", err.Error()))
		setExitCode(1)
	}
}
