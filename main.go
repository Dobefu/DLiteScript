// The main entrypoint of the application.
package main

import (
	"syscall"

	"github.com/Dobefu/DLiteScript/cmd"
)

func main() {
	syscall.Exit(cmd.Execute())
}
