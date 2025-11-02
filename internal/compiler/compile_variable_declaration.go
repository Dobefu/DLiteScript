package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileVariableDeclaration(vd *ast.VariableDeclaration) error {
	currentScope := c.variableScopes[len(c.variableScopes)-1]
	addr, hasAddress := currentScope[vd.Name]

	if !hasAddress {
		var numVars uint64

		for _, scope := range c.variableScopes {
			numVars += uint64(len(scope))
		}

		addr = numVars
		currentScope[vd.Name] = addr
	}

	if vd.Value != nil {
		err := c.compileNode(vd.Value)

		if err != nil {
			return err
		}

		valueRegister := c.getLastRegister()

		addrRegister := c.incrementRegCounter()
		err = c.emitLoadImmediate(addrRegister, int64(addr)) // #nosec: G115

		if err != nil {
			return err
		}

		return c.emitStoreMemory(valueRegister, addrRegister)
	}

	return nil
}
