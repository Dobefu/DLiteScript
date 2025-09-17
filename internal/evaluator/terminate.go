package evaluator

// Terminate stops the evaluator and exits with a given exit code.
func (e *Evaluator) Terminate(code byte) {
	e.exitCode = code
	e.shouldTerminate = true
}
