package compiler

import (
	"encoding/binary"

	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitLoadConstPoolIndex(dest byte, index int) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeLoadImmediate))
	c.bytecode = append(c.bytecode, dest)

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, uint64(index)) // #nosec: G115
	c.bytecode = append(c.bytecode, indexBytes...)

	return nil
}
