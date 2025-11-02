package compiler

import (
	"errors"
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileAssignmentStatement(as *ast.AssignmentStatement) error {
	if as.Left == nil {
		return errors.New("assignment statement has no left side")
	}

	var addr uint64
	var hasAddr bool

	for i := len(c.variableScopes) - 1; i >= 0; i-- {
		addr, hasAddr = c.variableScopes[i][as.Left.Value]

		if hasAddr {
			break
		}
	}

	if !hasAddr {
		return fmt.Errorf("undefined variable: %s", as.Left.Value)
	}

	err := c.compileNode(as.Right)

	if err != nil {
		return err
	}

	valueReg := c.getLastRegister()

	addrReg := c.incrementRegCounter()
	err = c.emitLoadImmediate(addrReg, int64(addr)) // #nosec: G115

	if err != nil {
		return err
	}

	return c.emitStoreMemory(valueReg, addrReg)
}
