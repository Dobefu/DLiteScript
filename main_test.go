package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob(filepath.Join("examples", "*", "main.dl"))

	if err != nil {
		t.Fatalf("Failed to find files: %s", err.Error())
	}

	if len(files) == 0 {
		t.Fatal("No files found in examples directory")
	}

	for _, file := range files {
		t.Run(filepath.Base(filepath.Dir(file)), func(t *testing.T) {
			t.Parallel()

			oldOsArgs := os.Args
			os.Args = []string{"DLiteScript", file}

			defer func() {
				os.Args = oldOsArgs
			}()

			main()
		})
	}
}
