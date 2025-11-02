package compiler

import (
	"strconv"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileNumberLiteral(n *ast.NumberLiteral) error {
	val, err := strconv.ParseFloat(n.Value, 64)

	if err != nil {
		return err
	}

	return c.emitLoadImmediate(c.incrementRegCounter(), int64(val))
}
