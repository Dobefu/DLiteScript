package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetCreateDirFunction(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	dirName := filepath.Join(tempDir, "this", "is", "a", "test", "directory")
	parentDirName := filepath.Join(tempDir, "this")
	if err := os.MkdirAll(dirName, 0700); err != nil {
		t.Fatalf("unable to create directory %s: %v", dirName, err)
	}

	getCreateDirFunc := getCreateDirFunction()

	result, err := getCreateDirFunc.Handler(
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

	err = os.RemoveAll(parentDirName)
	if err != nil {
		t.Fatalf("unable to delete directories with `os.RemoveAll`")
	}

	getCreateDirFunc = getCreateDirFunction()

	result, err = getCreateDirFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(dirName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error != nil {
		t.Fatalf("expected no error from func, but got: %v", result.Error)
	}

	err = os.RemoveAll(parentDirName)
	if err != nil {
		t.Fatalf("unable to delete directories with `os.RemoveAll`")
	}
}
