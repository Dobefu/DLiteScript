package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) compileReturnStatement(node *ast.ReturnStatement) error {
	regs := make([]byte, 0, len(node.Values))

	for _, value := range node.Values {
		err := c.compileNode(value)

		if err != nil {
			return err
		}

		regs = append(regs, c.getLastRegister())
		c.incrementRegCounter()
	}

	for i, srcReg := range regs {
		destReg := byte(i)

		if srcReg == destReg {
			continue
		}

		err := c.emitLoadRegister(destReg, srcReg)

		if err != nil {
			return err
		}
	}

	c.bytecode = append(c.bytecode, byte(vm.OpcodeReturn))

	return nil
}
