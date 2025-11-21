package compiler

import vm "github.com/Dobefu/vee-em"

func (c *Compiler) emitPush(reg byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodePush), reg)

	return nil
}
