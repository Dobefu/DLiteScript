package compiler

import (
	"encoding/binary"

	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitJmpImmediateIfLess(addr uint64) (int, error) {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeJmpImmediateIfLess))

	offset := len(c.bytecode)
	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, addr)
	c.bytecode = append(c.bytecode, addrBytes...)

	return offset, nil
}
