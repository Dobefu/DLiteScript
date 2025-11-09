package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (c *Compiler) compileComparison(
	destReg byte,
	leftReg byte,
	rightReg byte,
	op token.Type,
) error {
	err := c.emitLoadImmediate(destReg, 0)

	if err != nil {
		return err
	}

	err = c.emitCMP(leftReg, rightReg)

	if err != nil {
		return err
	}

	return c.emitLoadImmediate(destReg, 1)
}
