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
	folderName := filepath.Join(tempDir, "test_folder")
	if err := os.MkdirAll(folderName, 0600); err != nil {
		t.Fatalf("unable to create folder %s: %v", folderName, err)
	}

	getDeleteDirFunc := getDeleteDirFunction()

	result, err := getDeleteDirFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(folderName),
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
			datavalue.String(folderName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error == nil {
		t.Fatal("expected error from func, but got nil")
	}
}
