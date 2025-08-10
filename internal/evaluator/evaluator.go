// Package evaluator defines logic to evaluate an AST.
package evaluator

import (
	"strings"
)

// Evaluator defines the actual evaluator struct.
type Evaluator struct {
	outerScope  map[string]ScopedValue
	blockScopes []map[string]ScopedValue
	buf         strings.Builder
}

// NewEvaluator creates a new evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{
		outerScope:  make(map[string]ScopedValue),
		blockScopes: make([]map[string]ScopedValue, 0),
		buf:         strings.Builder{},
	}
}
