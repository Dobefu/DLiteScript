// Package compiler provides a compiler for DLiteScript.
package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	vm "github.com/Dobefu/vee-em"
)

// Compiler is the compiler for DLiteScript.
type Compiler struct {
	// The bytecode of the program.
	bytecode []byte
	// The register counter.
	regCounter byte
	// The constant pool.
	constPool []string
	// The index map of the constant in the constant pool.
	constPoolMap map[string]int
	// The function pool.
	functionPool []string
	// The index map of the function in the function pool.
	functionMap map[string]int
}

// NewCompiler creates a new compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		bytecode:     make([]byte, 0),
		regCounter:   0,
		constPool:    make([]string, 0),
		constPoolMap: make(map[string]int),
		functionPool: make([]string, 0),
		functionMap:  make(map[string]int),
	}
}

// Compile compiles the given AST node into bytecode.
func (c *Compiler) Compile(node ast.ExprNode) ([]byte, error) {
	// Add the magic header.
	c.bytecode = append(c.bytecode, []byte("DLS\x01")...)
	instructionsStart := len(c.bytecode)

	err := c.compileNode(node)

	if err != nil {
		return nil, err
	}

	c.bytecode = append(c.bytecode, byte(vm.OpcodeHalt))

	instructionsEnd := len(c.bytecode)

	program := make([]byte, 0)
	program = append(program, c.bytecode[:instructionsStart]...)
	program = append(program, c.serializeConstPool()...)
	program = append(program, c.serializeFunctionPool()...)
	program = append(program, c.bytecode[instructionsStart:instructionsEnd]...)

	return program, nil
}
