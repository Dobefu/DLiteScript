package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) getLastRegister() byte {
	if c.regCounter == 0 {
		return vm.NumRegisters - 1
	}

	return c.regCounter - 1
}
