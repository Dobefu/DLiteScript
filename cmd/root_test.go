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
				exitCode = 0
				cmdMutex.Unlock()
			}()

			runRootCmd(rootCmd, []string{file})

			if exitCode != 0 {
				t.Fatalf("Expected non-zero exit code, got %d", exitCode)
			}
		})
	}
}

func TestRootCmdErr(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	defer func() {
		exitCode = 0
		cmdMutex.Unlock()
	}()

	runRootCmd(rootCmd, []string{"bogus"})

	if exitCode == 0 {
		t.Fatalf("expected exit code not to be 0")
	}
}

func TestExecute(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	defer func() {
		exitCode = 0
		cmdMutex.Unlock()
	}()

	err := rootCmd.ValidateArgs([]string{"examples/00_simple/main.dl"})

	if err != nil {
		t.Fatalf("Expected no error, got: \"%s\"", err.Error())
	}

	rootCmd.SetArgs([]string{"../examples/00_simple/main.dl"})
	resultExitCode := Execute()

	if resultExitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", resultExitCode)
	}
}

func TestExecuteErr(t *testing.T) {
	t.Parallel()

	cmdMutex.Lock()
	defer func() {
		exitCode = 0
		cmdMutex.Unlock()
	}()

	rootCmd.SetArgs([]string{})
	resultExitCode := Execute()

	if resultExitCode == 0 {
		t.Fatalf("Expected non-zero exit code, got %d", resultExitCode)
	}
}
