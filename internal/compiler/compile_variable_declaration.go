package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileVariableDeclaration(vd *ast.VariableDeclaration) error {
	addr, exists := c.variableMap[vd.Name]

	if !exists {
		addr = uint64(len(c.variableMap))
		c.variableMap[vd.Name] = addr
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
