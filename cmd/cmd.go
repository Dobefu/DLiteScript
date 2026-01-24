// Package cmd contains the commands for the DLiteScript CLI.
package cmd

import "sync/atomic"

var _exitCode atomic.Uint32

func setExitCode(code byte) {
	_exitCode.Store(uint32(code))
}

func exitCode() byte {
	return byte(_exitCode.Load())
}

func resetExitCode() {
	_exitCode.Store(0)
}
