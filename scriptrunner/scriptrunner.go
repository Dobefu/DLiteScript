// Package scriptrunner provides a script runner for DLiteScript.
package scriptrunner

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Dobefu/DLiteScript/internal/evaluator"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

// ScriptRunner handles the execution of DLiteScript files.
type ScriptRunner struct {
	OutFile io.Writer

	result string
}

// ReadFileFromArgs reads a file from the arguments and returns its content.
func (r *ScriptRunner) ReadFileFromArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("no file specified")
	}

	fileContent, err := os.ReadFile(args[0])

	if err != nil {
		return "", fmt.Errorf("failed to read file: %s", err.Error())
	}

	return string(fileContent), nil
}

// RunString executes a DLiteScript script from a string.
func (r *ScriptRunner) RunString(str string, filePath ...string) (byte, error) {
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

	if len(filePath) > 0 && filePath[0] != "" {
		e.SetCurrentFilePath(filePath[0])
	}

	result, err := e.Evaluate(ast)

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

	if result.IsExitResult() {
		return byte(result.Control.Count & 0xFF), nil
	}

	return 0, nil
}

// RunScript executes a DLiteScript script file.
func (r *ScriptRunner) RunScript(file string) (byte, error) {
	fileContent, err := os.ReadFile(filepath.Clean(file))

	if err != nil {
		return 1, fmt.Errorf("failed to read file: %s", err.Error())
	}

	fileHeader := fileContent[:4]

	if string(fileHeader) != "DLS\x01" {
		return r.RunString(string(fileContent), file)
	}

	return r.RunBytecode(fileContent)
}

// Output returns the result of the execution.
func (r *ScriptRunner) Output() string {
	return r.result
}

// RunBytecode executes compiled DLiteScript bytecode.
func (r *ScriptRunner) RunBytecode(bytecode []byte) (byte, error) {
	rt := &vmRuntime{
		constPool:    make([]string, 0),
		functionPool: make([]string, 0),
		program:      make([]byte, 0),
	}

	err := rt.loadBytecode(bytecode)

	if err != nil {
		return 1, fmt.Errorf("failed to load bytecode: %s", err.Error())
	}

	err = rt.run(r.OutFile)

	if err != nil {
		return 1, fmt.Errorf("failed to run bytecode: %s", err.Error())
	}

	return 0, nil
}

// RunCompiledScript executes a compiled DLiteScript file.
func (r *ScriptRunner) RunCompiledScript(file string) (byte, error) {
	bytecode, err := os.ReadFile(filepath.Clean(file))

	if err != nil {
		return 1, fmt.Errorf("failed to read bytecode file: %s", err.Error())
	}

	return r.RunBytecode(bytecode)
}
