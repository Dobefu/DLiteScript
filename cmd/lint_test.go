package cmd

import (
	"testing"
)

func TestLintCmd(t *testing.T) {
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

			err := lintCmd.ValidateArgs([]string{test.input})

			if err != nil {
				t.Fatalf("Expected no error, got: \"%s\"", err.Error())
			}

			lintCmd.SetArgs([]string{test.input})
			runLintCmd(lintCmd, []string{test.input})

			if getExitCode() != 0 {
				t.Fatalf("Expected exit code 0, got %d", getExitCode())
			}
		})
	}
}

func TestLintCmdErr(t *testing.T) {
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
				err := lintCmd.ValidateArgs(input)

				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
			}

			lintCmd.SetArgs(input)
			runLintCmd(lintCmd, input)

			if getExitCode() == 0 {
				t.Fatalf("expected non-zero exit code, got 0")
			}
		})
	}
}
