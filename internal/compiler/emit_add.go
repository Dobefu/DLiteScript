package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitAdd(dest, src1, src2 byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeAdd))
	c.bytecode = append(c.bytecode, dest, src1, src2)

	return nil
}
