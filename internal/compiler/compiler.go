// Package compiler provides a compiler for DLiteScript.
package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	vm "github.com/Dobefu/vee-em"
)

// Compiler is the compiler for DLiteScript.
type Compiler struct {
	bytecode   []byte
	regCounter byte
}

// NewCompiler creates a new compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		bytecode:   make([]byte, 0),
		regCounter: 0,
	}
}

// Compile compiles the given AST node into bytecode.
func (c *Compiler) Compile(node ast.ExprNode) ([]byte, error) {
	// Add the magic header.
	c.bytecode = append(c.bytecode, []byte("DLS\x01")...)

	err := c.compileNode(node)

	if err != nil {
		return nil, err
	}

	c.bytecode = append(c.bytecode, byte(vm.OpcodeHalt))

	return c.bytecode, nil
}
