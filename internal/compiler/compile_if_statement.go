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
	jmpIfZeroPos, err := c.emitJmpImmediateIfZero(conditionRegister, 0)

	if err != nil {
		return err
	}

	err = c.compileNode(node.ThenBlock)

	if err != nil {
		return err
	}

	var elseOffset uint64

	if node.ElseBlock != nil {
		jmpEndPos, err := c.emitJmpImmediate(0)

		if err != nil {
			return err
		}

		elseOffset = c.getCurrentOffset()

		addrBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(addrBytes, elseOffset)
		copy(c.bytecode[jmpIfZeroPos:jmpIfZeroPos+8], addrBytes)

		err = c.compileNode(node.ElseBlock)

		if err != nil {
			return err
		}

		endOffset := c.getCurrentOffset()

		addrBytes = make([]byte, 8)
		binary.BigEndian.PutUint64(addrBytes, endOffset)
		copy(c.bytecode[jmpEndPos:jmpEndPos+8], addrBytes)

		return nil
	}

	elseOffset = c.getCurrentOffset()

	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, elseOffset)
	copy(c.bytecode[jmpIfZeroPos:jmpIfZeroPos+8], addrBytes)

	return nil
}
