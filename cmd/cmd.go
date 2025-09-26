// Package cmd contains the commands for the DLiteScript CLI.
package cmd

import "sync/atomic"

var exitCode atomic.Uint32

func setExitCode(code byte) {
	exitCode.Store(uint32(code))
}

func getExitCode() byte {
	return byte(exitCode.Load())
}

func resetExitCode() {
	exitCode.Store(0)
}
