package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/internal/formatter"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
	"github.com/Dobefu/DLiteScript/scriptrunner"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{ //nolint:exhaustruct
	Use: "fmt",
	Args: cobra.PositionalArgs(func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("usage: %s <file>", cmd.CommandPath())
		}

		return nil
	}),
	Short: "Format DLiteScript code",
	Run:   runFmtCmd,
}

func init() {
	rootCmd.AddCommand(fmtCmd)
}

func runFmtCmd(_ *cobra.Command, args []string) {
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

	formatter := formatter.New()
	formattedCode := formatter.Format(ast)

	_, _ = fmt.Fprintln(os.Stdout, formattedCode)
}
