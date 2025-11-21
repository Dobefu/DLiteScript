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
		vm.GetInstructionLen(vm.OpcodeJmpImmediateIfNotEqual) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate)

	switch op {
	case token.TokenTypeEqual:
		_, err = c.emitJmpImmediateIfNotEqual(jumpOffset)

	case token.TokenTypeNotEqual:
		_, err = c.emitJmpImmediateIfEqual(jumpOffset)

	case token.TokenTypeGreaterThan:
		_, err = c.emitJmpImmediateIfLessOrEqual(jumpOffset)

	case token.TokenTypeGreaterThanOrEqual:
		_, err = c.emitJmpImmediateIfLess(jumpOffset)

	case token.TokenTypeLessThan:
		_, err = c.emitJmpImmediateIfGreaterOrEqual(jumpOffset)

	case token.TokenTypeLessThanOrEqual:
		_, err = c.emitJmpImmediateIfGreater(jumpOffset)
	}

	if err != nil {
		return err
	}

	return c.emitLoadImmediate(destReg, 1)
}
