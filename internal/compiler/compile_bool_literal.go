package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileBoolLiteral(b *ast.BoolLiteral) error {
	var val int64

	if b.Value == "true" {
		val = 1
	} else {
		val = 0
	}

	return c.emitLoadImmediate(c.incrementRegCounter(), val)
}
