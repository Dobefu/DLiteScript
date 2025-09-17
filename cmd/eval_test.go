package cmd

import (
	"testing"
)

func TestEvalCmd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "simple printf",
			input: `printf("testing")`,
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

			err := evalCmd.ValidateArgs([]string{test.input})

			if err != nil {
				t.Fatalf("Expected no error, got: \"%s\"", err.Error())
			}

			evalCmd.SetArgs([]string{test.input})
			runEvalCmd(evalCmd, []string{test.input})

			if exitCode != 0 {
				t.Fatalf("expected exit code 0, got %d", exitCode)
			}
		})
	}
}

func TestEvalCmdErr(t *testing.T) {
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
			name:  "invalid syntax",
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
				err := evalCmd.ValidateArgs(input)

				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
			}

			evalCmd.SetArgs(input)
			runEvalCmd(evalCmd, input)

			if exitCode == 0 {
				t.Fatalf("expected non-zero exit code, got 0")
			}
		})
	}
}
