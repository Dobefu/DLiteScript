// The main entrypoint of the application.
package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/internal/evaluator"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

// Main is the main entrypoint of the application.
type Main struct {
	args    []string
	onError func(error)
	outFile io.Writer

	result string
}

// Run actually runs the application.
func (m *Main) Run() {
	if len(m.args) <= 1 {
		m.onError(errors.New("usage: go run main.go <file>"))

		return
	}

	fileContent, err := os.ReadFile(m.args[1])

	if err != nil {
		m.onError(err)

		return
	}

	t := tokenizer.NewTokenizer(string(fileContent))
	tokens, err := t.Tokenize()

	if err != nil {
		m.onError(err)

		return
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()

	if err != nil {
		m.onError(err)

		return
	}

	e := evaluator.NewEvaluator()
	_, err = e.Evaluate(ast)

	if err != nil {
		m.onError(err)

		return
	}

	m.result = e.Output()

	// If the output file is io.Discard, we don't need to format the result.
	if m.outFile == io.Discard {
		return
	}

	_, err = fmt.Fprint(m.outFile, m.result)

	if err != nil {
		m.onError(err)
	}
}

func main() {
	(&Main{
		args:    os.Args,
		outFile: os.Stdout,
		onError: func(err error) {
			slog.Error(err.Error())

			os.Exit(1)
		},

		result: "",
	}).Run()
}
