// Package evaluator defines logic to evaluate an AST.
package evaluator

import (
	"strings"
)

// Evaluator defines the actual evaluator struct.
type Evaluator struct {
	buf strings.Builder
}

// NewEvaluator creates a new evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{
		buf: strings.Builder{},
	}
}
