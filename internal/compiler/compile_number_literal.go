package compiler

import (
	"fmt"
	"strconv"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileNumberLiteral(n *ast.NumberLiteral) error {
	val, err := strconv.ParseFloat(n.Value, 64)

	if err != nil {
		return fmt.Errorf("failed to parse number literal: %s", err.Error())
	}

	return c.emitLoadImmediate(c.incrementRegCounter(), int64(val))
}
