package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/token"
	vm "github.com/Dobefu/vee-em"
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

	jumpOffset := c.getCurrentOffset() +
		vm.GetInstructionLen(vm.OpcodeJmpImmediate)

	switch op {
	case token.TokenTypeEqual:
		err = c.emitJmpImmediateIfEqual(jumpOffset)

	case token.TokenTypeNotEqual:
		err = c.emitJmpImmediateIfNotEqual(jumpOffset)

	case token.TokenTypeGreaterThan:
		err = c.emitJmpImmediateIfGreater(jumpOffset)

	case token.TokenTypeGreaterThanOrEqual:
		err = c.emitJmpImmediateIfGreaterOrEqual(jumpOffset)

	case token.TokenTypeLessThan:
		err = c.emitJmpImmediateIfLess(jumpOffset)

	case token.TokenTypeLessThanOrEqual:
		err = c.emitJmpImmediateIfLessOrEqual(jumpOffset)
	}

	if err != nil {
		return err
	}

	return c.emitLoadImmediate(destReg, 1)
}
