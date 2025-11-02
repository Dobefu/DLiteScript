package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/Dobefu/DLiteScript/internal/compiler"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{ //nolint:exhaustruct
	Use: "compile",
	Args: cobra.PositionalArgs(func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("usage: %s <file>", cmd.CommandPath())
		}

		return nil
	}),
	Short: "Compile a DLiteScript file to bytecode",
	Run:   runCompileCmd,
}

func init() {
	compileCmd.Flags().StringP("output", "o", "", "Output file path")

	rootCmd.AddCommand(compileCmd)
}

func runCompileCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		slog.Error("no file specified")
		setExitCode(1)

		return
	}

	fileContent, err := os.ReadFile(args[0])

	if err != nil {
		slog.Error(fmt.Sprintf("failed to read file: %s", err.Error()))
		setExitCode(1)

		return
	}

	t := tokenizer.NewTokenizer(string(fileContent))
	tokens, err := t.Tokenize()

	if err != nil {
		slog.Error(fmt.Sprintf("failed to tokenize: %s", err.Error()))
		setExitCode(1)

		return
	}

	p := parser.NewParser(tokens)
	astNode, err := p.Parse()

	if err != nil {
		slog.Error(fmt.Sprintf("failed to parse: %s", err.Error()))
		setExitCode(1)

		return
	}

	c := compiler.NewCompiler()
	bytecode, err := c.Compile(astNode)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to compile: %s", err.Error()))
		setExitCode(1)

		return
	}

	outputPath, _ := cmd.Flags().GetString("output")

	if outputPath == "" {
		ext := filepath.Ext(args[0])
		outputPath = fmt.Sprintf("%s.dlc", args[0][:len(args[0])-len(ext)])
	}

	err = os.WriteFile(outputPath, bytecode, 0600)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to write bytecode: %s", err.Error()))
		setExitCode(1)

		return
	}
}
