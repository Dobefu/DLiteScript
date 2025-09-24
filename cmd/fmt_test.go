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
				exitCode = 0
				cmdMutex.Unlock()
			}()

			err := fmtCmd.ValidateArgs([]string{test.input})

			if err != nil {
				t.Fatalf("Expected no error, got: \"%s\"", err.Error())
			}

			fmtCmd.SetArgs([]string{test.input})
			runFmtCmd(fmtCmd, []string{test.input})

			if exitCode != 0 {
				t.Fatalf("Expected exit code 0, got %d", exitCode)
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			cmdMutex.Lock()
			defer func() {
				exitCode = 0
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

			if exitCode == 0 {
				t.Fatalf("expected non-zero exit code, got 0")
			}
		})
	}
}
