package scriptrunner

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

type errWriter struct{}

func (f *errWriter) Write(_ []byte) (n int, err error) {
	return 0, errors.New("write error")
}

func TestScriptRunner(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		script   string
		expected string
	}{
		{
			name:     "printf",
			script:   `printf("test\n")`,
			expected: "test\n",
		},
		{
			name:     "exit",
			script:   "exit(0)\nprintf(\"test\")",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tmpFile, err := os.CreateTemp("", "test")

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			defer func() { _ = os.Remove(tmpFile.Name()) }()

			_, err = tmpFile.WriteString(test.script)

			if err != nil {
				t.Fatalf("expected no error writing script, got %v", err)
			}

			defer func() { _ = tmpFile.Close() }()

			outputBuffer := &bytes.Buffer{}

			_, err = (&ScriptRunner{
				OutFile: outputBuffer,
				result:  "",
			}).RunScript(tmpFile.Name())

			if err != nil {
				t.Fatalf("expected no error, got %s", err.Error())
			}

			if outputBuffer.String() != test.expected {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					outputBuffer.String(),
				)
			}
		})
	}
}

func TestScriptRunnerErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		hasReadErr bool
		outFile    io.Writer
		script     string
		expected   string
	}{
		{
			name:       "cannot read file",
			hasReadErr: true,
			outFile:    &bytes.Buffer{},
			script:     "",
			expected:   "open: permission denied",
		},
		{
			name:       "invalid utf-8",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "\x80",
			expected: fmt.Sprintf(
				"failed to tokenize file: %s: %s at position 0",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgInvalidUTF8Char,
			),
		},
		{
			name:       "unexpected EOF",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "1 +",
			expected: fmt.Sprintf(
				"failed to tokenize file: %s: %s at position 3",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name:       "function num args",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "printf()",
			expected: fmt.Sprintf(
				"failed to evaluate file: %s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "printf", 1, 0),
			),
		},
		{
			name:       "parsing error",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "1 + }",
			expected: fmt.Sprintf(
				"failed to parse file: %s: %s at position 3",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "}"),
			),
		},
		{
			name:       "write error",
			hasReadErr: false,
			outFile:    &errWriter{},
			script:     `1 + 1`,
			expected:   "failed to write to output file: write error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tmpFile, _ := os.CreateTemp("", test.name)
			args := []string{tmpFile.Name()}

			defer func() { _ = os.Remove(tmpFile.Name()) }()

			if test.hasReadErr {
				_ = os.Chmod(tmpFile.Name(), 0000)

				test.expected = fmt.Sprintf(
					"failed to read file: open %s: permission denied",
					tmpFile.Name(),
				)
			}

			_, _ = tmpFile.WriteString(test.script)
			defer func() { _ = tmpFile.Close() }()

			_, err := (&ScriptRunner{
				OutFile: test.outFile,
				result:  "",
			}).RunScript(args[0])

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	t.Parallel()

	scriptRunner := &ScriptRunner{
		OutFile: io.Discard,
		result:  "test",
	}

	if scriptRunner.Output() != "test" {
		t.Fatalf("expected \"test\", got \"%s\"", scriptRunner.Output())
	}

	tmpFile, err := os.CreateTemp("", "test")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	defer func() { _ = os.Remove(tmpFile.Name()) }()

	scriptRunner = &ScriptRunner{
		OutFile: io.Discard,
		result:  "test",
	}

	exitCode, err := scriptRunner.RunScript(tmpFile.Name())

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	if scriptRunner.Output() != "" {
		t.Fatalf("expected empty string, got \"%s\"", scriptRunner.Output())
	}
}
