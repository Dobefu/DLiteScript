package compiler

import (
	"encoding/binary"
)

func (c *Compiler) serializeConstPool() []byte {
	pool := make([]byte, 0)

	numBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(numBytes, uint64(len(c.constPool)))
	pool = append(pool, numBytes...)

	for _, str := range c.constPool {
		lengthBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(lengthBytes, uint64(len(str)))

		pool = append(pool, lengthBytes...)
		pool = append(pool, []byte(str)...)
	}

	return pool
}
