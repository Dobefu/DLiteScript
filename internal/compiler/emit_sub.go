package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitSub(dest, src1, src2 byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeSub))
	c.bytecode = append(c.bytecode, dest, src1, src2)

	return nil
}
