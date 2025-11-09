package compiler

import (
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) compileLogicalAnd(
	destReg byte,
	leftReg byte,
	rightReg byte,
) error {
	falseLabel := c.getCurrentOffset() +
		vm.GetInstructionLen(vm.OpcodeJmpImmediateIfZero) +
		vm.GetInstructionLen(vm.OpcodeJmpImmediateIfZero) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate) +
		vm.GetInstructionLen(vm.OpcodeJmpImmediate) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate)

	err := c.emitJmpImmediateIfZero(leftReg, falseLabel)

	if err != nil {
		return err
	}

	falseLabel = c.getCurrentOffset() +
		vm.GetInstructionLen(vm.OpcodeJmpImmediateIfZero) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate) +
		vm.GetInstructionLen(vm.OpcodeJmpImmediate) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate)

	err = c.emitJmpImmediateIfZero(rightReg, falseLabel)

	if err != nil {
		return err
	}

	err = c.emitLoadImmediate(destReg, 1)

	if err != nil {
		return err
	}

	endLabel := c.getCurrentOffset() +
		vm.GetInstructionLen(vm.OpcodeJmpImmediate) +
		vm.GetInstructionLen(vm.OpcodeLoadImmediate)

	err = c.emitJmpImmediate(endLabel)

	if err != nil {
		return err
	}

	return c.emitLoadImmediate(destReg, 0)
}
