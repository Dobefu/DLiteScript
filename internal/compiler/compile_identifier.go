package compiler

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileIdentifier(id *ast.Identifier) error {
	addr, exists := c.variableMap[id.Value]

	if !exists {
		return fmt.Errorf("undefined variable: %s", id.Value)
	}

	addrReg := c.incrementRegCounter()
	destReg := c.incrementRegCounter()

	err := c.emitLoadImmediate(addrReg, int64(addr)) // #nosec: G115

	if err != nil {
		return err
	}

	return c.emitLoadMemory(destReg, addrReg)
}
