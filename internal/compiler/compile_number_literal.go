package compiler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

var errFailedToParseNumberLiteral = errors.New("failed to parse number literal")

func (c *Compiler) compileNumberLiteral(n *ast.NumberLiteral) error {
	val, err := strconv.ParseFloat(n.Value, 64)

	if err != nil {
		return fmt.Errorf("%w: %w", errFailedToParseNumberLiteral, err)
	}

	return c.emitLoadImmediate(c.incrementRegCounter(), int64(val))
}
