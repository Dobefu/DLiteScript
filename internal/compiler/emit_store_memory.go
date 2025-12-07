package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitStoreMemory(src, addrReg byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeStoreMemory))
	c.bytecode = append(c.bytecode, src, addrReg)

	return nil
}
