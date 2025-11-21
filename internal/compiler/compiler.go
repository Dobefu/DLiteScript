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
	// The map of function names to their addresses in bytecode.
	functionAddresses map[string]uint64
	// The map of function names to patches that need to be resolved.
	functionCallPatches map[string][]int
	// Map of function names to their return value counts.
	functionReturnCounts map[string]int
	// Variable storage stack of scopes.
	variableScopes []map[string]uint64
	// Loop stack for break/continue support.
	loopStack []loopInfo
	// The start position of instructions in bytecode.
	instructionsStart int
}

type loopInfo struct {
	breakAddr       uint64
	continueAddr    uint64
	breakPatches    []int
	continuePatches []int
}

// NewCompiler creates a new compiler.
func NewCompiler() *Compiler {
	pool := []string{"false", "true", "null"}
	poolMap := map[string]int{
		"false": 0,
		"true":  1,
		"null":  2,
	}

	return &Compiler{
		bytecode:             make([]byte, 0),
		regCounter:           0,
		constPool:            pool,
		constPoolMap:         poolMap,
		functionPool:         make([]string, 0),
		functionMap:          make(map[string]int),
		functionAddresses:    make(map[string]uint64),
		functionCallPatches:  make(map[string][]int),
		functionReturnCounts: make(map[string]int),
		variableScopes:       []map[string]uint64{make(map[string]uint64)},
		loopStack:            make([]loopInfo, 0),
		instructionsStart:    0,
	}
}

// Compile compiles the given AST node into bytecode.
func (c *Compiler) Compile(node ast.ExprNode) ([]byte, error) {
	// Add the magic header.
	c.bytecode = append(c.bytecode, []byte("DLS\x01")...)
	c.instructionsStart = len(c.bytecode)

	err := c.compileNode(node)

	if err != nil {
		return nil, err
	}

	c.bytecode = append(c.bytecode, byte(vm.OpcodeHalt))

	instructionsEnd := len(c.bytecode)

	program := make([]byte, 0)
	program = append(program, c.bytecode[:c.instructionsStart]...)
	program = append(program, c.serializeConstPool()...)
	program = append(program, c.serializeFunctionPool()...)
	program = append(program, c.bytecode[c.instructionsStart:instructionsEnd]...)

	return program, nil
}
