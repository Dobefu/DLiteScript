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
	folderName := filepath.Join(tempDir, "this", "is", "a", "test", "folder")
	parentFolderName := filepath.Join(tempDir, "this")
	if err := os.MkdirAll(folderName, 0600); err != nil {
		t.Fatalf("unable to create folder %s: %v", folderName, err)
	}

	getCreateDirFunc := getCreateDirFunction()

	result, err := getCreateDirFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(folderName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error == nil {
		t.Fatal("expected error from func, but got nil")
	}

	err = os.RemoveAll(parentFolderName)
	if err != nil {
		t.Fatalf("unable to delete folders with `os.RemoveAll`")
	}

	getCreateDirFunc = getCreateDirFunction()

	result, err = getCreateDirFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(folderName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error != nil {
		t.Fatalf("expected no error from func, but got: %v", result.Error)
	}

	err = os.RemoveAll(parentFolderName)
	if err != nil {
		t.Fatalf("unable to delete folders with `os.RemoveAll`")
	}
}
