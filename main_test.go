package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
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

	for _, test := range tests {
		main := &Main{
			args:    []string{os.Args[0], createTmpFile(t, test.input)},
			outFile: io.Discard,
			onError: func(err error) {
				t.Errorf("expected no error, got '%s'", err.Error())
			},

			result: "",
		}

		main.Run()

		if main.result != "" {
			t.Errorf("expected empty output, got '%s'", main.result)
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
			expected: "usage: go run main.go <file>",
		},
		{
			input:    "1 +",
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input:    "\x80",
			expected: string(errorutil.ErrorMsgInvalidUTF8Char),
		},
		{
			input:    "min(1)",
			expected: fmt.Sprintf(errorutil.ErrorMsgFunctionNumArgs, "min", 2, 1),
		},
	}

	for _, test := range tests {
		var mainErr error
		args := []string{os.Args[0]}

		if test.input != "" {
			args = append(args, createTmpFile(t, test.input))
		}

		main := &Main{
			args:    args,
			outFile: io.Discard,
			onError: func(err error) {
				mainErr = errors.Unwrap(err)

				if mainErr == nil {
					mainErr = err
				}
			},

			result: "",
		}

		main.Run()

		if mainErr == nil {
			t.Fatalf("expected error, got none for input '%s'", test.input)
		}

		if mainErr.Error() != test.expected {
			t.Errorf(
				"expected error '%s', got '%s'",
				test.expected,
				mainErr.Error(),
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

	main := &Main{
		args:    []string{os.Args[0], filePath},
		outFile: buf,
		onError: func(err error) {
			mainErr = errors.Unwrap(err)

			if mainErr == nil {
				mainErr = err
			}
		},
		result: "",
	}

	main.Run()

	if mainErr == nil {
		t.Fatalf("expected error, got none")
	}
}

func BenchmarkMain(b *testing.B) {
	filePath := createTmpFile(b, "1 + -2 * 3 / 4")

	for b.Loop() {
		main := &Main{
			args:    []string{os.Args[0], filePath},
			outFile: io.Discard,
			onError: func(err error) {
				b.Errorf("expected no error, got '%s'", err.Error())
			},
			result: "",
		}

		main.Run()
	}
}
