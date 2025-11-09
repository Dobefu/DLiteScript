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

	leftRegister := c.getLastRegister()

	err = c.compileNode(b.Right)

	if err != nil {
		return err
	}

	rightRegister := c.getLastRegister()

	destReg := c.incrementRegCounter()

	switch b.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return c.emitAdd(destReg, leftRegister, rightRegister)

	case token.TokenTypeOperationSub:
		return c.emitSub(destReg, leftRegister, rightRegister)

	case token.TokenTypeOperationMul:
		return c.emitMul(destReg, leftRegister, rightRegister)

	case token.TokenTypeOperationDiv:
		return c.emitDiv(destReg, leftRegister, rightRegister)

	case token.TokenTypeOperationMod:
		return c.emitMod(destReg, leftRegister, rightRegister)

	case token.TokenTypeEqual:
		return c.compileComparison(destReg, leftRegister, rightRegister, token.TokenTypeEqual)

	default:
		return nil
	}
}
