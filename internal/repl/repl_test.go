package repl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

type errReader struct {
	err error
}

func (r *errReader) Read(_ []byte) (n int, err error) {
	return 0, r.err
}

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		outFile  io.Writer
		expected string
	}{
		{
			name:     "basic",
			input:    "1 + 1",
			outFile:  &bytes.Buffer{},
			expected: "dlitescript> => 2",
		},
		{
			name:     "multiple expressions",
			input:    "1 + 1\n2 + 2",
			outFile:  &bytes.Buffer{},
			expected: "dlitescript> => 2\ndlitescript> => 4",
		},
		{
			name:     "multiline continuation",
			input:    "1 + 1\\\n\n2 + 2\\\n3 + 3",
			outFile:  &bytes.Buffer{},
			expected: "dlitescript>   >   >   > => 6\ndlitescript> ",
		},
		{
			name:     "help command",
			input:    ".help",
			outFile:  &bytes.Buffer{},
			expected: "Available commands:",
		},
		{
			name:     "quit command",
			input:    ".quit",
			outFile:  &bytes.Buffer{},
			expected: "dlitescript> ",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repl := NewREPL(test.outFile, strings.NewReader(test.input))
			err := repl.Run()

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			if test.expected != "" && test.outFile != io.Discard {
				outputBuffer := test.outFile.(*bytes.Buffer)

				if !strings.Contains(outputBuffer.String(), test.expected) {
					t.Fatalf(
						"expected output to contain \"%s\", got \"%s\"",
						test.expected,
						outputBuffer.String(),
					)
				}
			}
		})
	}
}

func TestRunErr(t *testing.T) {
	t.Parallel()

	// Trigger a scan error
	repl := NewREPL(io.Discard, &errReader{err: errors.New("scan error")})
	err := repl.Run()

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

}

func TestEvaluateInputErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		outFile  io.Writer
		expected string
	}{
		{
			name:    "tokenization error",
			input:   "1 + 1 +",
			outFile: &bytes.Buffer{},
			expected: fmt.Sprintf(
				"Tokenization error: %s: %s line 1 at position 8\n",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name:    "evaluation error",
			input:   "1 + x",
			outFile: &bytes.Buffer{},
			expected: fmt.Sprintf(
				"Evaluation error: %s: %s line 1 at position 4\n",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "x"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repl := NewREPL(test.outFile, strings.NewReader(test.input))
			repl.evaluateInput(test.input)

			if test.outFile == nil {
				t.Fatalf("expected error, got nil")
			}

			if test.outFile.(*bytes.Buffer).String() != test.expected {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					test.outFile.(*bytes.Buffer).String(),
				)
			}
		})
	}
}
