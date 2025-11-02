package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) incrementRegCounter() byte {
	reg := c.regCounter
	c.regCounter = (c.regCounter + 1) % vm.NumRegisters

	return reg
}
