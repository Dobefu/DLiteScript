// Package scriptrunner provides a script runner for DLiteScript.
package scriptrunner

import (
	"fmt"
	"io"
	"os"

	"github.com/Dobefu/DLiteScript/internal/evaluator"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

// ScriptRunner handles the execution of DLiteScript files.
type ScriptRunner struct {
	OutFile io.Writer

	result string
}

// RunString executes a DLiteScript script file.
func (r *ScriptRunner) RunString(str string) (int, error) {
	t := tokenizer.NewTokenizer(str)
	tokens, err := t.Tokenize()

	if err != nil {
		return 1, fmt.Errorf("failed to tokenize file: %s", err.Error())
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()

	if err != nil {
		return 1, fmt.Errorf("failed to parse file: %s", err.Error())
	}

	e := evaluator.NewEvaluator(r.OutFile)
	_, err = e.Evaluate(ast)

	if err != nil {
		return 1, fmt.Errorf("failed to evaluate file: %s", err.Error())
	}

	r.result = e.Output()

	// If the output file is io.Discard, we don't need to format the result.
	if r.OutFile == io.Discard {
		return 0, nil
	}

	_, err = fmt.Fprint(r.OutFile, r.result)

	if err != nil {
		return 1, fmt.Errorf("failed to write to output file: %s", err.Error())
	}

	return 0, nil
}

// RunScript executes a DLiteScript script file.
func (r *ScriptRunner) RunScript(file string) (int, error) {
	fileContent, err := os.ReadFile(file)

	if err != nil {
		return 1, fmt.Errorf("failed to read file: %s", err.Error())
	}

	return r.RunString(string(fileContent))
}

// Output returns the result of the execution.
func (r *ScriptRunner) Output() string {
	return r.result
}
