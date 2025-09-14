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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tmpFile, err := os.CreateTemp("", "test")

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			defer func() { _ = os.Remove(tmpFile.Name()) }()

			_, err = tmpFile.WriteString(test.script)

			if err != nil {
				t.Fatalf("expected no error writing script, got %v", err)
			}

			defer func() { _ = tmpFile.Close() }()

			outputBuffer := &bytes.Buffer{}

			err = (&ScriptRunner{
				Args:    []string{tmpFile.Name()},
				OutFile: outputBuffer,
				result:  "",
			}).Run()

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
		hasFile    bool
		hasReadErr bool
		outFile    io.Writer
		script     string
		expected   string
	}{
		{
			name:       "no file",
			hasFile:    false,
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "",
			expected:   "no file specified",
		},
		{
			name:       "cannot read file",
			hasFile:    true,
			hasReadErr: true,
			outFile:    &bytes.Buffer{},
			script:     "",
			expected:   "open: permission denied",
		},
		{
			name:       "invalid utf-8",
			hasFile:    true,
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
			hasFile:    true,
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
			hasFile:    true,
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
			hasFile:    true,
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
			hasFile:    true,
			hasReadErr: false,
			outFile:    &errWriter{},
			script:     `1 + 1`,
			expected:   "failed to write to output file: write error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			args := []string{}
			tmpFile, _ := os.CreateTemp("", test.name)

			if test.hasFile {
				args = append(args, tmpFile.Name())
			}

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

			err := (&ScriptRunner{
				Args:    args,
				OutFile: test.outFile,
				result:  "",
			}).Run()

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
		Args:    []string{},
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
		Args:    []string{tmpFile.Name()},
		OutFile: io.Discard,
		result:  "test",
	}

	err = scriptRunner.Run()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if scriptRunner.Output() != "" {
		t.Fatalf("expected empty string, got \"%s\"", scriptRunner.Output())
	}
}
