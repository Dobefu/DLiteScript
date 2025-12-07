package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) compileLogicalOr(destReg, leftReg, rightReg byte) error {
	trueLabel := c.getCurrentOffset() +
		vm.GetInstructionLen(vm.OpcodeJmpImmediateIfNotZero) +
		vm.GetInstructionLen(vm.OpcodeJmpImmediateIfNotZero) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate) +
		vm.GetInstructionLen(vm.OpcodeJmpImmediate)

	_, err := c.emitJmpImmediateIfNotZero(leftReg, trueLabel)

	if err != nil {
		return err
	}

	trueLabel = c.getCurrentOffset() +
		vm.GetInstructionLen(vm.OpcodeJmpImmediateIfNotZero) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate) +
		vm.GetInstructionLen(vm.OpcodeJmpImmediate)

	_, err = c.emitJmpImmediateIfNotZero(rightReg, trueLabel)

	if err != nil {
		return err
	}

	err = c.emitLoadImmediate(destReg, 0)

	if err != nil {
		return err
	}

	endLabel := c.getCurrentOffset() +
		vm.GetInstructionLen(vm.OpcodeJmpImmediate) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate)

	_, err = c.emitJmpImmediate(endLabel)

	if err != nil {
		return err
	}

	return c.emitLoadImmediate(destReg, 1)
}
