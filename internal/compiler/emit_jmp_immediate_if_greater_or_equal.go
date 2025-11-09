package compiler

import (
	"encoding/binary"

	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitJmpImmediateIfGreaterOrEqual(addr uint64) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeJmpImmediateIfGreaterOrEqual))

	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, addr)
	c.bytecode = append(c.bytecode, addrBytes...)

	return nil
}
