package ast

import (
	"fmt"
)

// Position represents a position in the source code.
type Position struct {
	Offset int
	Line   int
	Column int
}

// Range represents a range in the source code.
type Range struct {
	Start Position
	End   Position
}

// IsMultiLine returns true if the range is multi-line.
func (r *Range) IsMultiLine() bool {
	return r.Start.Line != r.End.Line
}

// LineSpan returns the number of lines in the range.
func (r *Range) LineSpan() int {
	return r.End.Line - r.Start.Line + 1
}

// String returns the string representation of the range.
func (r *Range) String() string {
	if r.IsMultiLine() {
		return fmt.Sprintf(
			"from line %d at position %d to line %d at position %d",
			r.Start.Line+1,
			r.Start.Column+1,
			r.End.Line+1,
			r.End.Column+1,
		)
	}

	return fmt.Sprintf("line %d at position %d", r.Start.Line+1, r.Start.Column+1)
}
