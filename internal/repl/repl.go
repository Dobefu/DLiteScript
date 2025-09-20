// Package repl provides a REPL for DLiteScript.
package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/evaluator"
	"github.com/Dobefu/DLiteScript/internal/parser"
	"github.com/Dobefu/DLiteScript/internal/tokenizer"
)

// REPL represents a REPL for DLiteScript.
type REPL struct {
	OutFile     io.Writer
	InFile      io.Reader
	evaluator   *evaluator.Evaluator
	isMultiline bool
	buf         strings.Builder
}

// NewREPL creates a new REPL.
func NewREPL(outFile io.Writer, inFile io.Reader) *REPL {
	return &REPL{
		OutFile:     outFile,
		InFile:      inFile,
		evaluator:   evaluator.NewEvaluator(outFile),
		isMultiline: false,
		buf:         strings.Builder{},
	}
}

// Run starts the REPL loop.
func (r *REPL) Run() error {
	scanner := bufio.NewScanner(r.InFile)

	r.initialize()

	for {
		prompt := r.getPrompt()
		_, _ = fmt.Fprint(r.OutFile, prompt)

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		shouldExit := r.handleCommand(line)

		if shouldExit {
			break
		}

		r.processInput(line)
	}

	err := scanner.Err()

	if err != nil {
		return fmt.Errorf("scanner error: %s", err.Error())
	}

	return nil
}

func (r *REPL) initialize() {
	var sb strings.Builder

	sb.WriteString("DLiteScript REPL\n")
	sb.WriteString("Type '.help' for commands, '.exit' to quit.\n")
	sb.WriteString("Use '\\' at the end of a line to continue on the next line.\n\n")

	_, _ = fmt.Fprint(r.OutFile, sb.String())
}

func (r *REPL) getPrompt() string {
	if r.isMultiline {
		return "  > "
	}

	return "dlitescript> "
}

func (r *REPL) handleCommand(line string) bool {
	switch line {
	case
		".exit",
		".quit":

		return true

	case
		".help":
		r.showHelp()

		return false

	default:
		return false
	}
}

func (r *REPL) showHelp() {
	var sb strings.Builder

	sb.WriteString("Available commands:\n")
	sb.WriteString("  .help  - Show this help\n")
	sb.WriteString("  .exit  - Exit the REPL\n")
	sb.WriteString("  .quit  - Exit the REPL\n")

	_, _ = fmt.Fprint(r.OutFile, sb.String())
}

func (r *REPL) processInput(line string) {
	if strings.HasSuffix(line, "\\") {
		if r.isMultiline {
			r.buf.WriteString(strings.TrimSuffix(line, "\\"))
			r.buf.WriteString("\n")

			return
		}

		r.isMultiline = true
		r.buf.WriteString(strings.TrimSuffix(line, "\\"))
		r.buf.WriteString("\n")

		return
	}

	input := line

	if r.isMultiline {
		r.buf.WriteString(line)
		input = r.buf.String()
		r.isMultiline = false
		r.buf.Reset()
	}

	r.evaluateInput(input)
}

func (r *REPL) evaluateInput(input string) {
	t := tokenizer.NewTokenizer(input)
	tokens, err := t.Tokenize()

	if err != nil {
		_, _ = fmt.Fprintf(r.OutFile, "Tokenization error: %s\n", err.Error())

		return
	}

	p := parser.NewParser(tokens)
	ast, err := p.Parse()

	if err != nil {
		_, _ = fmt.Fprintf(r.OutFile, "Parse error: %s\n", err.Error())

		return
	}

	result, err := r.evaluator.Evaluate(ast)

	if err != nil {
		_, _ = fmt.Fprintf(r.OutFile, "Evaluation error: %s\n", err.Error())

		return
	}

	if result.Value.DataType != datatype.DataTypeNull {
		_, _ = fmt.Fprintf(r.OutFile, "=> %s\n", result.Value.ToString())
	}
}
