package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (c *Compiler) compilePrefixExpr(p *ast.PrefixExpr) error {
	err := c.compileNode(p.Operand)

	if err != nil {
		return err
	}

	operandReg := c.getLastRegister()

	switch p.Operator.TokenType {
	case token.TokenTypeNot:
		zeroRegister := c.incrementRegCounter()
		err = c.emitLoadImmediate(zeroRegister, 0)

		if err != nil {
			return err
		}

		destReg := c.incrementRegCounter()

		return c.compileComparison(destReg, operandReg, zeroRegister, token.TokenTypeEqual)

	case token.TokenTypeOperationSub:
		zeroReg := c.incrementRegCounter()
		err = c.emitLoadImmediate(zeroReg, 0)

		if err != nil {
			return err
		}

		destReg := c.incrementRegCounter()

		return c.emitSub(destReg, zeroReg, operandReg)

	default:
		return nil
	}
}
