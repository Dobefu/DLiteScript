package compiler

import "github.com/Dobefu/DLiteScript/internal/ast"

func (c *Compiler) compileNullLiteral(_ *ast.NullLiteral) error {
	reg := c.incrementRegCounter()
	err := c.emitLoadImmediate(reg, 2)

	if err != nil {
		return err
	}

	return nil
}
