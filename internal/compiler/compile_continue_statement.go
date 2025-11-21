package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileContinueStatement(node *ast.ContinueStatement) error {
	if len(c.loopStack) == 0 {
		return nil
	}

	targetLoop := max(len(c.loopStack)-node.Count, 0)
	jmpPos, err := c.emitJmpImmediate(0)

	if err != nil {
		return err
	}

	c.loopStack[targetLoop].continuePatches = append(
		c.loopStack[targetLoop].continuePatches,
		jmpPos,
	)

	return nil
}
