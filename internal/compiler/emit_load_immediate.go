package compiler

import (
	"encoding/binary"

	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) emitLoadImmediate(dest byte, value int64) error {
	c.bytecode = append(c.bytecode, byte(vm.OpcodeLoadImmediate))
	c.bytecode = append(c.bytecode, dest)

	valBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(valBytes, uint64(value))
	c.bytecode = append(c.bytecode, valBytes...)

	return nil
}
