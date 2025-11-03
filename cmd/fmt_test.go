package cmd

import (
	"testing"
)

func TestFmtCmd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "simple file",
			input: "../examples/00_simple/main.dl",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			cmdMutex.Lock()
			defer func() {
				resetExitCode()
				cmdMutex.Unlock()
			}()

			err := fmtCmd.ValidateArgs([]string{test.input})

			if err != nil {
				t.Fatalf("expected no error, got: \"%s\"", err.Error())
			}

			fmtCmd.SetArgs([]string{test.input})
			runFmtCmd(fmtCmd, []string{test.input})

			if getExitCode() != 0 {
				t.Fatalf("Expected exit code 0, got %d", getExitCode())
			}
		})
	}
}

func TestFmtCmdErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "missing arguments",
			input: "",
		},
		{
			name:  "invalid file path",
			input: "bogus",
		},
		{
			name:  "syntax error",
			input: "./testfiles/syntax_error.dl",
		},
		{
			name:  "parse error",
			input: "./testfiles/parse_error.dl",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			cmdMutex.Lock()
			defer func() {
				resetExitCode()
				cmdMutex.Unlock()
			}()

			input := []string{}

			if test.input != "" {
				input = append(input, test.input)
			}

			if len(input) == 0 {
				err := fmtCmd.ValidateArgs(input)

				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
			}

			fmtCmd.SetArgs(input)
			runFmtCmd(fmtCmd, input)

			if getExitCode() == 0 {
				t.Fatalf("expected non-zero exit code, got 0")
			}
		})
	}
}
