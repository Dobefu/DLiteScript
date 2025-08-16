package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/scriptrunner"
)

func createTmpFile(t testing.TB, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "tmp.dl")

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("error creating tmp file: %v", err)
	}

	return path
}

func TestMainRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected float64
	}{
		{
			input:    "1",
			expected: 1,
		},
		{
			input:    "1 + 1",
			expected: 2,
		},
		{
			input:    "(1 + 1)",
			expected: 2,
		},
		{
			input:    "1 + 2 * 3",
			expected: 7,
		},
		{
			input:    "(1 + 2) * 3",
			expected: 9,
		},
		{
			input:    "1 + 2 * 3 / 4",
			expected: 2.5,
		},
		{
			input:    "8 * 5 % 3",
			expected: 1,
		},
	}

	errMsgUnexpectedErr := "expected no error, got '%s'"

	for _, test := range tests {
		runner := &scriptrunner.ScriptRunner{
			Args:    []string{createTmpFile(t, test.input)},
			OutFile: io.Discard,
			OnError: func(err error) {
				t.Errorf(errMsgUnexpectedErr, err.Error())
			},
		}

		err := runner.Run()

		if err != nil {
			t.Errorf(errMsgUnexpectedErr, err.Error())
		}

		if runner.Output() != "" {
			t.Errorf("expected empty output, got '%s'", runner.Output())
		}
	}
}

func TestMainErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "no file specified",
		},
		{
			input:    "1 +",
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 2",
		},
		{
			input:    "\x80",
			expected: string(errorutil.ErrorMsgInvalidUTF8Char) + " at position 0",
		},
		{
			input:    "min(1)",
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "min", 2, 1) + " at position 0",
		},
	}

	for _, test := range tests {
		var mainErr error
		var args []string

		if test.input != "" {
			args = append(args, createTmpFile(t, test.input))
		}

		runner := &scriptrunner.ScriptRunner{
			Args:    args,
			OutFile: io.Discard,
			OnError: func(err error) {
				mainErr = errors.Unwrap(err)

				if mainErr == nil {
					mainErr = err
				}
			},
		}

		err := runner.Run()

		if err == nil {
			t.Fatalf("expected error, got none for input '%s'", test.input)
		}

		actualErr := err

		if mainErr != nil {
			actualErr = mainErr
		}

		if actualErr.Error() != test.expected {
			t.Errorf(
				"expected error '%s', got '%s'",
				test.expected,
				actualErr.Error(),
			)
		}
	}
}

func TestMainWriteError(t *testing.T) {
	t.Parallel()

	buf, _ := os.OpenFile("/some/bogus/file.txt", os.O_RDONLY, 0)
	defer func() { _ = buf.Close() }()

	var mainErr error

	filePath := createTmpFile(t, "1 + 1")

	runner := &scriptrunner.ScriptRunner{
		Args:    []string{filePath},
		OutFile: buf,
		OnError: func(err error) {
			mainErr = errors.Unwrap(err)

			if mainErr == nil {
				mainErr = err
			}
		},
	}

	err := runner.Run()

	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func BenchmarkMain(b *testing.B) {
	filePath := createTmpFile(b, "1 + -2 * 3 / 4")
	errMsgUnexpectedErr := "expected no error, got '%s'"

	for b.Loop() {
		runner := &scriptrunner.ScriptRunner{
			Args:    []string{filePath},
			OutFile: io.Discard,
			OnError: func(err error) {
				b.Errorf(errMsgUnexpectedErr, err.Error())
			},
		}

		err := runner.Run()

		if err != nil {
			b.Errorf(errMsgUnexpectedErr, err.Error())
		}
	}
}
