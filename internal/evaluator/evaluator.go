// Package evaluator defines logic to evaluate an AST.
package evaluator

import (
	"io"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

// Evaluator defines the actual evaluator struct.
type Evaluator struct {
	outerScope      map[string]ScopedValue
	blockScopes     []map[string]ScopedValue
	blockScopesLen  int
	userFunctions   map[string]*ast.FuncDeclarationStatement
	buf             strings.Builder
	outFile         io.Writer
	shouldTerminate bool
	exitCode        byte
}

// NewEvaluator creates a new evaluator.
func NewEvaluator(outFile io.Writer) *Evaluator {
	return &Evaluator{
		outerScope:      make(map[string]ScopedValue),
		blockScopes:     make([]map[string]ScopedValue, 0),
		blockScopesLen:  0,
		userFunctions:   make(map[string]*ast.FuncDeclarationStatement),
		buf:             strings.Builder{},
		outFile:         outFile,
		shouldTerminate: false,
		exitCode:        0,
	}
}

func (e *Evaluator) pushBlockScope() {
	newScope := make(map[string]ScopedValue)
	e.blockScopes = append(e.blockScopes, newScope)
	e.blockScopesLen = len(e.blockScopes)
}

func (e *Evaluator) popBlockScope() {
	if e.blockScopesLen > 0 {
		e.blockScopes = e.blockScopes[:e.blockScopesLen-1]
		e.blockScopesLen--
	}
}
