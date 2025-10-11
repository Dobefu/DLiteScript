package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/scriptrunner"
	"github.com/spf13/cobra"
)

var evalCmd = &cobra.Command{ //nolint:exhaustruct
	Use:     "eval",
	Aliases: []string{"ev"},
	Args: cobra.PositionalArgs(func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("usage: %s <expression>", cmd.CommandPath())
		}

		return nil
	}),
	Short: "Evaluate a DLiteScript expression",
	Run:   runEvalCmd,
}

func init() {
	evalCmd.Flags().BoolP("quiet", "q", false, "Don't print any messages to the output")

	rootCmd.AddCommand(evalCmd)
}

func runEvalCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		slog.Error("no code provided")
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
	code, err := runner.RunString(args[0])
	setExitCode(code)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to run script: %s", err.Error()))
		setExitCode(1)
	}
}
