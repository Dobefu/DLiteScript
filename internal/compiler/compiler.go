// Package compiler provides a compiler for DLiteScript.
package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

// Compiler is the compiler for DLiteScript.
type Compiler struct {
	bytecode []byte
}

// NewCompiler creates a new compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		bytecode: make([]byte, 0),
	}
}

// Compile compiles the given AST node into bytecode.
func (c *Compiler) Compile(node ast.ExprNode) ([]byte, error) {
	return nil, nil
}
