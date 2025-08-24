package scriptrunner

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

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
		script     string
		expected   string
	}{
		{
			name:       "no file",
			hasFile:    false,
			hasReadErr: false,
			script:     "",
			expected:   "no file specified",
		},
		{
			name:       "cannot read file",
			hasFile:    true,
			hasReadErr: true,
			script:     "",
			expected:   "open: permission denied",
		},
		{
			name:       "invalid utf-8",
			hasFile:    true,
			hasReadErr: false,
			script:     "\x80",
			expected:   string(errorutil.ErrorMsgInvalidUTF8Char) + " at position 0",
		},
		{
			name:       "unexpected EOF",
			hasFile:    true,
			hasReadErr: false,
			script:     "1 +",
			expected:   errorutil.ErrorMsgUnexpectedEOF + " at position 2",
		},
		{
			name:       "function num args",
			hasFile:    true,
			hasReadErr: false,
			script:     "min(1)",
			expected:   fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "min", 2, 1) + " at position 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			args := []string{}

			tmpFile, err := os.CreateTemp("", test.name)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if test.hasFile {
				args = append(args, tmpFile.Name())
			}

			defer func() { _ = os.Remove(tmpFile.Name()) }()

			if test.hasReadErr {
				_ = os.Chmod(tmpFile.Name(), 0000)
				test.expected = fmt.Sprintf("open %s: permission denied", tmpFile.Name())
			}

			_, _ = tmpFile.WriteString(test.script)
			defer func() { _ = tmpFile.Close() }()

			outputBuffer := &bytes.Buffer{}

			err = (&ScriptRunner{
				Args:    args,
				OutFile: outputBuffer,
				result:  "",
			}).Run()

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
