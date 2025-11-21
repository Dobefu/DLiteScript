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

	argRegisters, err := c.compileArguments(fc.Arguments)

	if err != nil {
		return err
	}

	if functionName == "printf" {
		return c.emitHostCall(functionName, argRegisters)
	}

	return c.emitScriptCall(functionName, argRegisters)
}

func (c *Compiler) compileArguments(args []ast.ExprNode) ([]byte, error) {
	var registers []byte

	for _, arg := range args {
		startRegister := c.regCounter
		err := c.compileNode(arg)

		if err != nil {
			return nil, err
		}

		_, isSpread := arg.(*ast.SpreadExpr)

		if isSpread {
			endRegister := c.regCounter

			for r := startRegister; r < endRegister; r++ {
				registers = append(registers, r)
			}
		} else {
			registers = append(registers, c.getLastRegister())
		}
	}

	return registers, nil
}

func (c *Compiler) emitHostCall(name string, argRegisters []byte) error {
	functionIndex := c.addToFunctionPool(name)
	argCount := len(argRegisters)
	startRegister := c.incrementRegCounter()

	for i, srcRegister := range argRegisters {
		destRegister := (startRegister + byte(i)) % vm.NumRegisters

		if destRegister != srcRegister {
			err := c.emitLoadRegister(destRegister, srcRegister)

			if err != nil {
				return err
			}
		}
	}

	c.regCounter = (startRegister + byte(argCount)) % vm.NumRegisters

	return c.emitFunctionCall(functionIndex, startRegister, byte(argCount))
}

func (c *Compiler) emitScriptCall(name string, argRegisters []byte) error {
	callStartRegister := c.regCounter

	if len(argRegisters) > 0 {
		minRegister := argRegisters[0]

		for _, register := range argRegisters {
			if register < minRegister {
				minRegister = register
			}
		}

		callStartRegister = minRegister
	}

	err := c.saveRegisters(callStartRegister)

	if err != nil {
		return err
	}

	for i, srcRegister := range argRegisters {
		err := c.moveRegister(byte(i), srcRegister)

		if err != nil {
			return err
		}
	}

	err = c.emitCallInstruction(name)

	if err != nil {
		return err
	}

	returnCount := c.getReturnCount(name)

	for i := returnCount - 1; i >= 0; i-- {
		src := byte(i)
		dest := callStartRegister + byte(i)

		err := c.moveRegister(dest, src)

		if err != nil {
			return err
		}
	}

	err = c.restoreRegisters(callStartRegister)

	if err != nil {
		return err
	}

	c.regCounter = callStartRegister + byte(returnCount)

	return nil
}

func (c *Compiler) emitCallInstruction(name string) error {
	addr, exists := c.functionAddresses[name]

	if !exists {
		patchPos, err := c.emitCallImmediate(0)

		if err != nil {
			return err
		}

		c.functionCallPatches[name] = append(c.functionCallPatches[name], patchPos)

		return nil
	}

	_, err := c.emitCallImmediate(addr)

	return err
}

func (c *Compiler) moveRegister(dest, src byte) error {
	if dest == src {
		return nil
	}

	return c.emitLoadRegister(dest, src)
}

func (c *Compiler) saveRegisters(limit byte) error {
	for i := range limit {
		err := c.emitPush(i)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) restoreRegisters(limit byte) error {
	for i := int(limit) - 1; i >= 0; i-- {
		err := c.emitPop(byte(i))

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) getReturnCount(name string) int {
	count, hasCount := c.functionReturnCounts[name]

	if hasCount && count > 0 {
		return count
	}

	return 1
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
