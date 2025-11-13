package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetDeleteFileFunction(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	fileName := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(fileName, []byte(""), 0600); err != nil {
		t.Fatalf("unable to create / write to file %s: %v", fileName, err)
	}

	getDeleteFileFunc := getDeleteFileFunction()

	result, err := getDeleteFileFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error != nil {
		t.Fatal("expected no error from func, but got nil")
	}

	getDeleteFileFunc = getDeleteFileFunction()

	result, err = getDeleteFileFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
		},
	)

	if result.Error == nil {
		t.Fatal("expected error from func, but got nil")
	}
}
