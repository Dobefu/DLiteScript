package io

import (
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetReadFileStringFunction(t *testing.T) {
	t.Parallel()

	fileName := "data.txt"
	expectedContent := "This is a simple test text, nothing more."
	if err := os.WriteFile(fileName, []byte(expectedContent), 0644); err != nil {
		t.Fatalf("unable to create / write to file %s: %v", fileName, err)
	}
	defer os.Remove(fileName)

	getReadFileStringFunc := getReadFileStringFunction()

	result, err := getReadFileStringFunc.Handler(
		nil,
		[]datavalue.Value{
			datavalue.String(fileName),
		},
	)
	if err != nil {
		t.Fatalf("expected no error from handler, got: %v", err)
	}

	actualContent, err := result.AsString()
	if err != nil {
		t.Fatalf("expected result to be a string value, got conversion error: %v", err)
	}

	if actualContent != expectedContent {
		t.Fatalf("expected file content to be %q, got %q", expectedContent, actualContent)
	}
}
