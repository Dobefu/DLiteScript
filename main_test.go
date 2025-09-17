package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	t.Parallel()

	oldOsArgs := os.Args
	os.Args = []string{"DLiteScript", "-q", "examples/00_simple/main.dl"}
	defer func() { os.Args = oldOsArgs }()

	main()
}
