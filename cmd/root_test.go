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

			err := runRootCmd(rootCmd, []string{file})

			if err != nil {
				t.Fatalf("Failed to run script: %s", err.Error())
			}
		})
	}
}
