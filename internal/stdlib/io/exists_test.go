package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetExistsFunction(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	fileName := filepath.Join(tempDir, "data.txt")
	expectedContent := "This is a simple test text, nothing more."
	if err := os.WriteFile(fileName, []byte(expectedContent), 0600); err != nil {
		t.Fatalf("unable to create / write to file %s: %v", fileName, err)
	}

	getExistsFunc := getExistsFunction()

	result, err := getExistsFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	resultTuple, _ := result.AsTuple()
	fileExists, _ := resultTuple[0].AsBool()

	if !fileExists {
		t.Fatalf("expected true from func, but got false")
	}

	err = os.Remove(fileName)
	if err != nil {
		t.Fatalf("unable to delete file with `os.Remove`")
	}

	getExistsFunc = getExistsFunction()

	result, err = getExistsFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Bool == true {
		t.Fatalf("expected false from func, but got true")
	}
}
