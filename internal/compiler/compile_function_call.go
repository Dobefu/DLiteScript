package compiler

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
	vm "github.com/Dobefu/vee-em"
)

func (c *Compiler) compileFunctionCall(fc *ast.FunctionCall) error {
	functionName := fc.FunctionName

	if fc.Namespace != "" {
		functionName = fmt.Sprintf("%s.%s", fc.Namespace, fc.FunctionName)
	}

	functionIndex := c.addToFunctionPool(functionName)
	argCount := len(fc.Arguments)
	argResultRegisters := make([]byte, 0, argCount)

	for _, arg := range fc.Arguments {
		err := c.compileNode(arg)

		if err != nil {
			return err
		}

		resultRegister := c.getLastRegister()
		argResultRegisters = append(argResultRegisters, resultRegister)
	}

	arg1Register := c.incrementRegCounter()

	for i, srcRegister := range argResultRegisters {
		destRegister := (arg1Register + byte(i)) % vm.NumRegisters

		if destRegister == srcRegister {
			continue
		}

		err := c.emitLoadRegister(destRegister, srcRegister)

		if err != nil {
			return err
		}
	}

	c.regCounter = (arg1Register + byte(argCount)) % vm.NumRegisters
	err := c.emitFunctionCall(functionIndex, arg1Register, byte(argCount))

	if err != nil {
		return err
	}

	return nil
}

func (c *Compiler) addToFunctionPool(name string) int {
	index, hasIndex := c.functionMap[name]

	if hasIndex {
		return index
	}

	index = len(c.functionPool)
	c.functionPool = append(c.functionPool, name)
	c.functionMap[name] = index

	return index
}
