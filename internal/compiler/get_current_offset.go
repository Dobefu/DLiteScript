package compiler

func (c *Compiler) getCurrentOffset() uint64 {
	return uint64(len(c.bytecode) - c.instructionsStart) // #nosec: G115
}
