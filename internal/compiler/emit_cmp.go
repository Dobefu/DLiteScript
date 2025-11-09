package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitCMP(src1, src2 byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeCMP))
	c.bytecode = append(c.bytecode, src1, src2)

	return nil
}
