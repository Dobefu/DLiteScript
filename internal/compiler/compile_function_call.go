package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileFunctionCall(fc *ast.FunctionCall) error {
	for _, arg := range fc.Arguments {
		err := c.compileNode(arg)

		if err != nil {
			return err
		}
	}

	return nil
}
