package scriptrunner

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

var errWrite = errors.New("write error")

type writerError struct{}

func (f *writerError) Write(_ []byte) (n int, err error) {
	return 0, errWrite
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
				t.Fatalf("expected no error writing script, got: \"%s\"", err.Error())
			}

			defer func() { _ = tmpFile.Close() }()

			outputBuffer := &bytes.Buffer{}

			_, err = (&ScriptRunner{
				OutFile: outputBuffer,
				result:  "",
			}).RunScript(tmpFile.Name())

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
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
		expected   error
	}{
		{
			name:       "cannot read file",
			hasReadErr: true,
			outFile:    &bytes.Buffer{},
			script:     "",
			expected:   os.ErrPermission,
		},
		{
			name:       "invalid utf-8",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "\x80",
			expected:   errFailedToTokenizeFile,
		},
		{
			name:       "unexpected EOF",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "1 +",
			expected:   errFailedToTokenizeFile,
		},
		{
			name:       "function num args",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "printf()",
			expected:   errFailedToEvaluateFile,
		},
		{
			name:       "parsing error",
			hasReadErr: false,
			outFile:    &bytes.Buffer{},
			script:     "1 + }",
			expected:   errFailedToParseFile,
		},
		{
			name:       "write error",
			hasReadErr: false,
			outFile:    &writerError{},
			script:     "1 + 1",
			expected:   errWrite,
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

				test.expected = os.ErrPermission
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

			if !strings.Contains(err.Error(), test.expected.Error()) {
				t.Fatalf(
					"expected error to be \"%s\", got \"%s\"",
					test.expected,
					err,
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
		t.Fatalf("expected no error, got: \"%s\"", err.Error())
	}

	defer func() { _ = os.Remove(tmpFile.Name()) }()

	scriptRunner = &ScriptRunner{
		OutFile: io.Discard,
		result:  "test",
	}

	exitCode, err := scriptRunner.RunScript(tmpFile.Name())

	if err != nil {
		t.Fatalf("expected no error, got: \"%s\"", err.Error())
	}

	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	if scriptRunner.Output() != "" {
		t.Fatalf("expected empty string, got \"%s\"", scriptRunner.Output())
	}
}
