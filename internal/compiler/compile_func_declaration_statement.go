package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) compileFuncDeclarationStatement(node *ast.FuncDeclarationStatement) error {
	c.addToFunctionPool(node.Name)

	jmpOverPos, err := c.emitJmpImmediate(0)

	if err != nil {
		return err
	}

	funcStart := c.getCurrentOffset()
	c.functionAddresses[node.Name] = funcStart
	patches, hasPatches := c.functionCallPatches[node.Name]

	if hasPatches {
		for _, patchPos := range patches {
			c.patchJump(patchPos, funcStart)
		}

		delete(c.functionCallPatches, node.Name)
	}

	oldRegCounter := c.regCounter
	c.regCounter = 0
	c.variableScopes = append(c.variableScopes, make(map[string]uint64))
	c.regCounter = byte(len(node.Args))

	for i, arg := range node.Args {
		varAddr := c.allocateVariable(arg.Name)
		argReg := byte(i)
		err := c.storeLoopVariable(varAddr, argReg)

		if err != nil {
			return err
		}
	}

	err = c.compileNode(node.Body)

	if err != nil {
		return err
	}

	c.bytecode = append(c.bytecode, byte(vm.OpcodeReturn))
	c.variableScopes = c.variableScopes[:len(c.variableScopes)-1]
	c.regCounter = oldRegCounter

	funcEnd := c.getCurrentOffset()
	c.patchJump(jmpOverPos, funcEnd)

	return nil
}

func (c *Compiler) allocateVariable(name string) uint64 {
	var numVars uint64

	for _, scope := range c.variableScopes {
		numVars += uint64(len(scope))
	}

	c.variableScopes[len(c.variableScopes)-1][name] = numVars

	return numVars
}
