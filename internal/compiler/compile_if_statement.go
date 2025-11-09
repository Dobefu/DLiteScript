package compiler

import (
	"encoding/binary"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileIfStatement(node *ast.IfStatement) error {
	err := c.compileNode(node.Condition)

	if err != nil {
		return err
	}

	conditionRegister := c.getLastRegister()
	jmpIfZeroPos := c.getCurrentOffset()
	err = c.emitJmpImmediateIfZero(conditionRegister, 0)

	if err != nil {
		return err
	}

	err = c.compileNode(node.ThenBlock)

	if err != nil {
		return err
	}

	var elseOffset uint64

	if node.ElseBlock != nil {
		jmpEndPos := c.getCurrentOffset()

		err = c.emitJmpImmediate(0)

		if err != nil {
			return err
		}

		elseOffset = c.getCurrentOffset()

		addrBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(addrBytes, elseOffset)
		jmpIfZeroBytecodePos := c.instructionsStart + int(jmpIfZeroPos) // #nosec: G115
		copy(c.bytecode[jmpIfZeroBytecodePos+2:jmpIfZeroBytecodePos+10], addrBytes)

		err = c.compileNode(node.ElseBlock)

		if err != nil {
			return err
		}

		endOffset := c.getCurrentOffset()

		addrBytes = make([]byte, 8)
		binary.BigEndian.PutUint64(addrBytes, endOffset)
		jmpEndBytecodePos := c.instructionsStart + int(jmpEndPos) // #nosec: G115
		copy(c.bytecode[jmpEndBytecodePos+1:jmpEndBytecodePos+9], addrBytes)

		return nil
	}

	elseOffset = c.getCurrentOffset()

	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, elseOffset)
	jmpIfZeroBytecodePos := c.instructionsStart + int(jmpIfZeroPos) // #nosec: G115
	copy(c.bytecode[jmpIfZeroBytecodePos+2:jmpIfZeroBytecodePos+10], addrBytes)

	return nil
}
