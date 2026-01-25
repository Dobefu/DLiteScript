// Package scriptrunner provides a script runner for DLiteScript.
package scriptrunner

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/evaluator"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

var errNoFile = errors.New("no file specified")
var errFailedToReadFile = errors.New("failed to read file")
var errFailedToTokenizeString = errors.New("failed to tokenize string")
var errFailedToTokenizeFile = errors.New("failed to tokenize file")
var errFailedToParseString = errors.New("failed to parse string")
var errFailedToParseFile = errors.New("failed to parse file")
var errFailedToEvaluateFile = errors.New("failed to evaluate file")
var errFailedToWriteToOutputFile = errors.New("failed to write to output file")
var errFailedToLoadBytecode = errors.New("failed to load bytecode")
var errFailedToRunBytecode = errors.New("failed to run bytecode")
var errFailedToReadBytecodeFile = errors.New("failed to read bytecode file")

// ScriptRunner handles the execution of DLiteScript files.
type ScriptRunner struct {
	OutFile io.Writer

	result string
}

// ReadFileFromArgs reads a file from the arguments and returns its content.
func (r *ScriptRunner) ReadFileFromArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errNoFile
	}

	fileContent, err := os.ReadFile(args[0])

	if err != nil {
		return "", fmt.Errorf("%q: %w", errFailedToReadFile, err)
	}

	return string(fileContent), nil
}

// ParseString parses a string to an AST.
func (r *ScriptRunner) ParseString(str string) (ast.ExprNode, error) {
	t := tokenizer.NewTokenizer(str)
	tokens, err := t.Tokenize()

	if err != nil {
		return nil, fmt.Errorf("%q: %w", errFailedToTokenizeString, err)
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()

	if err != nil {
		return nil, fmt.Errorf("%q: %w", errFailedToParseString, err)
	}

	return ast, nil
}

// RunString executes a DLiteScript script from a string.
func (r *ScriptRunner) RunString(str string, filePath ...string) (byte, error) {
	t := tokenizer.NewTokenizer(str)
	tokens, err := t.Tokenize()

	if err != nil {
		return 1, fmt.Errorf("%q: %w", errFailedToTokenizeFile, err)
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()

	if err != nil {
		return 1, fmt.Errorf("%q: %w", errFailedToParseFile, err)
	}

	e := evaluator.NewEvaluator(r.OutFile)

	if len(filePath) > 0 && filePath[0] != "" {
		e.SetCurrentFilePath(filePath[0])
	}

	result, err := e.Evaluate(ast)

	if err != nil {
		return 1, fmt.Errorf("%q: %w", errFailedToEvaluateFile, err)
	}

	r.result = e.Output()

	// If the output file is io.Discard, we don't need to format the result.
	if r.OutFile == io.Discard {
		return 0, nil
	}

	_, err = fmt.Fprint(r.OutFile, r.result)

	if err != nil {
		return 1, fmt.Errorf("%q: %w", errFailedToWriteToOutputFile, err)
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
		return 1, fmt.Errorf("%q: %w", errFailedToReadFile, err)
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
		return 1, fmt.Errorf("%q: %w", errFailedToLoadBytecode, err)
	}

	err = rt.run(r.OutFile)

	if err != nil {
		return 1, fmt.Errorf("%q: %w", errFailedToRunBytecode, err)
	}

	return 0, nil
}

// RunCompiledScript executes a compiled DLiteScript file.
func (r *ScriptRunner) RunCompiledScript(file string) (byte, error) {
	bytecode, err := os.ReadFile(filepath.Clean(file))

	if err != nil {
		return 1, fmt.Errorf("%q: %w", errFailedToReadBytecodeFile, err)
	}

	return r.RunBytecode(bytecode)
}
