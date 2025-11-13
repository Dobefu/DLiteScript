package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetDeleteDirFunction(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	dirName := filepath.Join(tempDir, "test_directory")
	if err := os.MkdirAll(dirName, 0700); err != nil {
		t.Fatalf("unable to create directory %s: %v", dirName, err)
	}

	getDeleteDirFunc := getDeleteDirFunction()

	result, err := getDeleteDirFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(dirName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error != nil {
		t.Fatal("expected no error from func, but got nil")
	}

	getDeleteDirFunc = getDeleteDirFunction()

	result, err = getDeleteDirFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(dirName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error == nil {
		t.Fatal("expected error from func, but got nil")
	}
}
