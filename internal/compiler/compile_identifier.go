package compiler

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileIdentifier(id *ast.Identifier) error {
	var addr uint64
	var hasAddr bool

	for i := len(c.variableScopes) - 1; i >= 0; i-- {
		addr, hasAddr = c.variableScopes[i][id.Value]

		if hasAddr {
			break
		}
	}

	if !hasAddr {
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
