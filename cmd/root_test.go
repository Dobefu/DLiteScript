package cmd

import (
	"path/filepath"
	"testing"
)

func TestRootCmd(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob(filepath.Join("..", "examples", "*", "main.dl"))

	if err != nil {
		t.Fatalf("Failed to find files: %s", err.Error())
	}

	if len(files) == 0 {
		t.Fatal("No files found in examples directory")
	}

	for _, file := range files {
		t.Run(filepath.Base(filepath.Dir(file)), func(t *testing.T) {
			t.Parallel()

			cmdMutex.Lock()
			defer func() {
				resetExitCode()
				cmdMutex.Unlock()
			}()

			runRootCmd(rootCmd, []string{file})

			if getExitCode() != 0 {
				t.Fatalf("Expected zero exit code, got %d", getExitCode())
			}
		})
	}
}

func TestRootCmdErr(t *testing.T) {
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
				resetExitCode()
				cmdMutex.Unlock()
			}()

			input := []string{}

			if test.input != "" {
				input = append(input, test.input)
			}

			if len(input) == 0 {
				err := rootCmd.ValidateArgs(input)

				if err == nil {
					t.Fatalf("Expected error, got nil")
				}
			}

			rootCmd.SetArgs(input)
			runRootCmd(rootCmd, input)

			if getExitCode() == 0 {
				t.Fatalf("expected non-zero exit code, got 0")
			}
		})
	}
}

func TestExecute(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	defer func() {
		resetExitCode()
		cmdMutex.Unlock()
	}()

	err := rootCmd.ValidateArgs([]string{"examples/00_simple/main.dl"})

	if err != nil {
		t.Fatalf("Expected no error, got: \"%s\"", err.Error())
	}

	rootCmd.SetArgs([]string{"-q", "../examples/00_simple/main.dl"})
	resultExitCode := Execute()

	if resultExitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", resultExitCode)
	}
}

func TestExecuteErr(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	defer func() {
		resetExitCode()
		cmdMutex.Unlock()
	}()

	rootCmd.SetArgs([]string{})
	resultExitCode := Execute()

	if resultExitCode == 0 {
		t.Fatalf("Expected non-zero exit code, got %d", resultExitCode)
	}
}
