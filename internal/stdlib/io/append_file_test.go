package io

import (
	"os"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestGetAppendFileFunction(t *testing.T) {
	t.Parallel()

	fileName := "data.txt"
	content := "\nThis is a simple test text, nothing more."
	defaultContent := "This should remain undisturbed."
	if err := os.WriteFile(fileName, []byte(defaultContent), 0644); err != nil {
		t.Fatalf("unable to create / write to file %s: %v", fileName, err)
	}
	defer os.Remove(fileName)

	getAppendFileFunc := getAppendFileFunction()

	_, err := getAppendFileFunc.Handler(
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

	expectedData := string(defaultContent) + content
	if string(data) != expectedData {
		t.Fatalf("content of file does not match the content that was expected")
	}
}
