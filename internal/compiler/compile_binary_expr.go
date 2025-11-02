package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (c *Compiler) compileBinaryExpr(b *ast.BinaryExpr) error {
	err := c.compileNode(b.Left)

	if err != nil {
		return err
	}

	leftRegister := c.regCounter - 1
	err = c.compileNode(b.Right)

	if err != nil {
		return err
	}

	rightRegister := c.regCounter - 1
	destReg := c.incrementRegCounter()

	switch b.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return c.emitAdd(destReg, leftRegister, rightRegister)

	default:
		return nil
	}
}
