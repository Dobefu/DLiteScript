package compiler

import vm "github.com/Dobefu/vee-em"

func (c *Compiler) emitPop(reg byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodePop), reg)

	return nil
}
