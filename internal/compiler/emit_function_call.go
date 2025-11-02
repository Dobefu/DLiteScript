package compiler

import (
	"encoding/binary"

	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitFunctionCall(functionIndex int, firstArgReg byte, argCount byte) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeHostCall))

	functionIndexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(functionIndexBytes, uint64(functionIndex)) // #nosec: G115
	c.bytecode = append(c.bytecode, functionIndexBytes...)

	c.bytecode = append(c.bytecode, firstArgReg, argCount)

	return nil
}
