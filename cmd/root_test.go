package cmd

import (
	"path/filepath"
	"sync"
	"testing"
)

var rootCmdMutex sync.Mutex

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

			rootCmdMutex.Lock()
			err := runRootCmd(rootCmd, []string{file})
			rootCmdMutex.Unlock()

			if err != nil {
				t.Fatalf("Failed to run script: %s", err.Error())
			}
		})
	}
}

func TestRootCmdErr(t *testing.T) {
	t.Parallel()

	rootCmdMutex.Lock()
	err := runRootCmd(rootCmd, []string{"bogus"})
	rootCmdMutex.Unlock()

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestExecute(t *testing.T) {
	t.Parallel()

	rootCmdMutex.Lock()
	err := rootCmd.ValidateArgs([]string{"examples/00_simple/main.dl"})
	rootCmdMutex.Unlock()

	if err != nil {
		t.Fatalf("Expected no error, got: \"%s\"", err.Error())
	}

	rootCmdMutex.Lock()
	rootCmd.SetArgs([]string{"../examples/00_simple/main.dl"})
	exitCode := Execute()
	rootCmdMutex.Unlock()

	if exitCode != 0 {
		t.Fatalf("Expected exit code 0, got %d", exitCode)
	}
}

func TestExecuteErr(t *testing.T) {
	t.Parallel()

	rootCmdMutex.Lock()
	rootCmd.SetArgs([]string{})
	exitCode := Execute()
	rootCmdMutex.Unlock()

	if exitCode == 0 {
		t.Fatalf("Expected non-zero exit code, got %d", exitCode)
	}
}
