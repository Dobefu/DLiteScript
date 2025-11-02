package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitLoadRegister(dest byte, src byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeLoadRegister))
	c.bytecode = append(c.bytecode, dest, src)

	return nil
}
