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

	addr, exists := c.variableMap[as.Left.Value]

	if !exists {
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
