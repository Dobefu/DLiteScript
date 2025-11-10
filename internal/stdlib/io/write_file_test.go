package io

import (
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetWriteFileFunction(t *testing.T) {
	t.Parallel()

	fileName := "data.txt"
	content := "This is a simple test text, nothing more."
	if err := os.WriteFile(fileName, []byte("This should not see light."), 0644); err != nil {
		t.Fatalf("unable to create / write to file %s: %v", fileName, err)
	}
	defer os.Remove(fileName)

	getWriteFileFunc := getWriteFileFunction()

	_, err := getWriteFileFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
			datavalue.String(content),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("unable to read file: %v", err)
	}
	if string(data) != content {
		t.Fatalf("content of file does not match the content that should have been written to the file.")
	}
}
