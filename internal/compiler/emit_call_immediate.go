package compiler

import (
	"encoding/binary"

	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitCallImmediate(addr uint64) (int, error) {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeCallImmediate))
	addrOffset := len(c.bytecode)

	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, addr)
	c.bytecode = append(c.bytecode, addrBytes...)

	return addrOffset, nil
}
