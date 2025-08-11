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

func (e *Evaluator) pushBlockScope() {
	newScope := make(map[string]ScopedValue)
	e.blockScopes = append(e.blockScopes, newScope)
}

func (e *Evaluator) popBlockScope() {
	if len(e.blockScopes) > 0 {
		e.blockScopes = e.blockScopes[:len(e.blockScopes)-1]
	}
}
