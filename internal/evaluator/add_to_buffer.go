package evaluator

import (
	"fmt"
)

// AddToBuffer adds data to the buffer.
func (e *Evaluator) AddToBuffer(format string, args ...any) {
	fmt.Fprintf(&e.buf, format, args...)
}
