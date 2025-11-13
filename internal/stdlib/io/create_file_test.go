package io

import (
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetCreateFileFunction(t *testing.T) {
	t.Parallel()

	fileName := "data.txt"
	if err := os.WriteFile(fileName, []byte(""), 0644); err != nil {
		t.Fatalf("unable to create / write to file %s: %v", fileName, err)
	}

	getCreateFileFunc := getCreateFileFunction()

	result, err := getCreateFileFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	if result.Error == nil {
		t.Fatal("expected error from func, but got nil")
	}

	err = os.Remove(fileName)
	if err != nil {
		t.Fatalf("unable to delete file with `os.Remove`")
	}

	getCreateFileFunc = getCreateFileFunction()

	result, err = getCreateFileFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
		},
	)

	if result.Error != nil {
		t.Fatalf("expected no error from func, but got: %v", result.Error)
	}

	err = os.Remove(fileName)
	if err != nil {
		t.Fatalf("unable to delete file with `os.Remove`")
	}
}
