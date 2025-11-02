package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitLoadMemory(dest byte, addrReg byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeLoadMemory))
	c.bytecode = append(c.bytecode, dest, addrReg)

	return nil
}
