package io

import (
	"os"
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetCreateDirFunction(t *testing.T) {
	t.Parallel()

	folderName := "this\\is\\a\\test\\folder"
	if err := os.MkdirAll(folderName, 0644); err != nil {
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

	splitFolder := strings.Split(folderName, "\\")
	os.RemoveAll(splitFolder[0])

	getCreateDirFunc = getCreateDirFunction()

	result, err = getCreateDirFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(folderName),
		},
	)

	if result.Error != nil {
		t.Fatalf("expected no error from func, but got: %v", result.Error)
	}

	os.RemoveAll(splitFolder[0])
}
