// Package scriptrunner provides a script runner for DLiteScript.
package scriptrunner

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Dobefu/DLiteScript/internal/evaluator"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

// ScriptRunner handles the execution of DLiteScript files.
type ScriptRunner struct {
	Args    []string
	OutFile io.Writer

	result string
}

// Run executes the DLiteScript file processing.
func (r *ScriptRunner) Run() error {
	if len(r.Args) == 0 {
		return errors.New("no file specified")
	}

	fileContent, err := os.ReadFile(r.Args[0])

	if err != nil {
		return fmt.Errorf("failed to read file: %s", err.Error())
	}

	t := tokenizer.NewTokenizer(string(fileContent))
	tokens, err := t.Tokenize()

	if err != nil {
		return fmt.Errorf("failed to tokenize file: %s", err.Error())
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()

	if err != nil {
		return fmt.Errorf("failed to parse file: %s", err.Error())
	}

	e := evaluator.NewEvaluator(r.OutFile)
	_, err = e.Evaluate(ast)

	if err != nil {
		return fmt.Errorf("failed to evaluate file: %s", err.Error())
	}

	r.result = e.Output()

	// If the output file is io.Discard, we don't need to format the result.
	if r.OutFile == io.Discard {
		return nil
	}

	_, err = fmt.Fprint(r.OutFile, r.result)

	if err != nil {
		return fmt.Errorf("failed to write to output file: %s", err.Error())
	}

	return nil
}

// Output returns the result of the execution.
func (r *ScriptRunner) Output() string {
	return r.result
}
