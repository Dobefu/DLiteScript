package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/internal/evaluator"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
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

// ScriptRunner handles the execution of DLiteScript files.
type ScriptRunner struct {
	Args    []string
	OnError func(error)
	OutFile io.Writer

	result string
}

// Run executes the DLiteScript file processing.
func (r *ScriptRunner) Run() error {
	if len(r.Args) == 0 {
		return fmt.Errorf("no file specified")
	}

	fileContent, err := os.ReadFile(r.Args[0])

	if err != nil {
		return err
	}

	t := tokenizer.NewTokenizer(string(fileContent))
	tokens, err := t.Tokenize()
	if err != nil {
		return err
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()
	if err != nil {
		return err
	}

	e := evaluator.NewEvaluator(r.OutFile)
	_, err = e.Evaluate(ast)

	if err != nil {
		return err
	}

	r.result = e.Output()

	// If the output file is io.Discard, we don't need to format the result.
	if r.OutFile == io.Discard {
		return nil
	}

	_, err = fmt.Fprint(r.OutFile, r.result)

	return err
}

// Output returns the result of the execution.
func (r *ScriptRunner) Output() string {
	return r.result
}

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func runRootCmd(_ *cobra.Command, args []string) {
	runner := &ScriptRunner{
		Args:    args,
		OutFile: os.Stdout,
		OnError: func(err error) {
			slog.Error(err.Error())
			os.Exit(1)
		},
		result: "",
	}

	err := runner.Run()

	if err != nil {
		runner.OnError(err)
	}
}
