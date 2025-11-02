package compiler

import (
	"encoding/binary"
)

func (c *Compiler) serializeFunctionPool() []byte {
	pool := make([]byte, 0)

	numBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(numBytes, uint64(len(c.functionPool)))
	pool = append(pool, numBytes...)

	for _, name := range c.functionPool {
		lengthBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(lengthBytes, uint64(len(name)))

		pool = append(pool, lengthBytes...)
		pool = append(pool, []byte(name)...)
	}

	return pool
}
