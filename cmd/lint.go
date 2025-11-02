package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/internal/linter"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
	"github.com/Dobefu/DLiteScript/scriptrunner"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{ //nolint:exhaustruct
	Use: "lint",
	Args: cobra.PositionalArgs(func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("usage: %s <file>", cmd.CommandPath())
		}

		return nil
	}),
	Short: "Lint DLiteScript files for common issues",
	Run:   runLintCmd,
}

func init() {
	rootCmd.AddCommand(lintCmd)
}

func runLintCmd(_ *cobra.Command, args []string) {
	var outfile io.Writer = os.Stdout

	isQuiet, _ := rootCmd.Flags().GetBool("quiet")

	if isQuiet {
		outfile = io.Discard
	}

	runner := &scriptrunner.ScriptRunner{
		OutFile: outfile,
	}

	fileContent, err := runner.ReadFileFromArgs(args)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to read file: %s", err.Error()))
		setExitCode(1)

		return
	}

	t := tokenizer.NewTokenizer(fileContent)
	tokens, err := t.Tokenize()

	if err != nil {
		slog.Error(fmt.Sprintf("failed to tokenize file: %s", err.Error()))
		setExitCode(1)

		return
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()

	if err != nil {
		slog.Error(fmt.Sprintf("failed to parse file: %s", err.Error()))
		setExitCode(1)

		return
	}

	l := linter.New(os.Stdout)
	l.Lint(ast)
	l.PrintIssues(args[0])

	if l.HasIssues() {
		setExitCode(1)
	}
}
