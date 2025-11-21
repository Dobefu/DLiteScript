package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileBreakStatement(node *ast.BreakStatement) error {
	if len(c.loopStack) == 0 {
		return nil
	}

	targetLoop := max(len(c.loopStack)-node.Count, 0)
	jmpPos, err := c.emitJmpImmediate(0)

	if err != nil {
		return err
	}

	c.loopStack[targetLoop].breakPatches = append(
		c.loopStack[targetLoop].breakPatches,
		jmpPos,
	)

	return nil
}
