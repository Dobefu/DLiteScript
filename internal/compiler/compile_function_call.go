package compiler

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileFunctionCall(fc *ast.FunctionCall) error {
	functionName := fc.FunctionName

	if fc.Namespace != "" {
		functionName = fmt.Sprintf("%s.%s", fc.Namespace, fc.FunctionName)
	}

	functionIndex := c.addToFunctionPool(functionName)

	for _, arg := range fc.Arguments {
		err := c.compileNode(arg)

		if err != nil {
			return err
		}
	}

	argCount := len(fc.Arguments)
	firstArgReg := c.regCounter - byte(argCount)

	return c.emitFunctionCall(functionIndex, firstArgReg, byte(argCount))
}

func (c *Compiler) addToFunctionPool(name string) int {
	index, exists := c.functionMap[name]

	if exists {
		return index
	}

	index = len(c.functionPool)
	c.functionPool = append(c.functionPool, name)
	c.functionMap[name] = index

	return index
}
