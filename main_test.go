package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestMain(t *testing.T) {
	t.Parallel()

	originalOsArgs := os.Args
	os.Args = []string{os.Args[0], "1 + 1"}

	main()

	os.Args = originalOsArgs
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
			args:    []string{os.Args[0], test.input},
			outFile: io.Discard,
			onError: func(err error) {
				t.Errorf("expected no error, got '%s'", err.Error())
			},

			result: 0,
		}

		main.Run()

		if main.result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, main.result)
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
			expected: "usage: go run main.go <expression>",
		},
		{
			input:    "1 +",
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input:    " ",
			expected: errorutil.ErrorMsgEmptyExpression,
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
			args = append(args, test.input)
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

			result: 0,
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

	main := &Main{
		args:    []string{os.Args[0], "1 + 1"},
		outFile: buf,
		onError: func(err error) {
			mainErr = errors.Unwrap(err)

			if mainErr == nil {
				mainErr = err
			}
		},
		result: 0,
	}

	main.Run()

	if mainErr == nil {
		t.Fatalf("expected error, got none")
	}
}

func BenchmarkMain(b *testing.B) {
	for b.Loop() {
		main := &Main{
			args:    []string{os.Args[0], "1 + -2 * 3 / 4"},
			outFile: io.Discard,
			onError: func(err error) {
				b.Errorf("expected no error, got '%s'", err.Error())
			},
			result: 0,
		}

		main.Run()
	}
}
