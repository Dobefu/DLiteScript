package compiler

import (
	"encoding/binary"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (c *Compiler) compileForStatement(node *ast.ForStatement) error {
	c.variableScopes = append(c.variableScopes, make(map[string]uint64))

	defer func() {
		c.variableScopes = c.variableScopes[:len(c.variableScopes)-1]
	}()

	if node.IsRange {
		return c.compileRangeLoop(node)
	}

	if node.Condition == nil {
		return c.compileInfiniteLoop(node)
	}

	return c.compileConditionLoop(node)
}

func (c *Compiler) compileInfiniteLoop(node *ast.ForStatement) error {
	loopStart := c.getCurrentOffset()

	c.loopStack = append(c.loopStack, loopInfo{
		continueAddr:    loopStart,
		breakAddr:       0,
		breakPatches:    make([]int, 0),
		continuePatches: make([]int, 0),
	})

	defer func() { c.loopStack = c.loopStack[:len(c.loopStack)-1] }()

	err := c.compileNode(node.Body)

	if err != nil {
		return err
	}

	_, err = c.emitJmpImmediate(loopStart)

	if err != nil {
		return err
	}

	c.patchLoopBreaks(c.getCurrentOffset())
	c.patchLoopContinues(loopStart)

	return nil
}

func (c *Compiler) compileConditionLoop(node *ast.ForStatement) error {
	var varAddr uint64
	var err error

	if node.DeclaredVariable != "" {
		currentScope := c.variableScopes[len(c.variableScopes)-1]
		var numVars uint64

		for _, scope := range c.variableScopes {
			numVars += uint64(len(scope))
		}

		varAddr = numVars
		currentScope[node.DeclaredVariable] = varAddr
		valRegister := c.incrementRegCounter()

		err := c.emitLoadImmediate(valRegister, 0)

		if err != nil {
			return err
		}

		err = c.storeLoopVariable(varAddr, valRegister)

		if err != nil {
			return err
		}
	}

	loopStart := c.getCurrentOffset()
	c.loopStack = append(c.loopStack, loopInfo{
		continueAddr:    loopStart,
		breakAddr:       0,
		breakPatches:    make([]int, 0),
		continuePatches: make([]int, 0),
	})

	defer func() { c.loopStack = c.loopStack[:len(c.loopStack)-1] }()

	err = c.compileNode(node.Condition)

	if err != nil {
		return err
	}

	conditionRegister := c.getLastRegister()
	jmpEndPos, err := c.emitJmpImmediateIfZero(conditionRegister, 0)

	if err != nil {
		return err
	}

	err = c.compileNode(node.Body)

	if err != nil {
		return err
	}

	if node.DeclaredVariable != "" {
		incrementStart := c.getCurrentOffset()
		c.patchLoopContinues(incrementStart)

		err = c.incrementLoopVariable(varAddr)

		if err != nil {
			return err
		}
	} else {
		c.patchLoopContinues(loopStart)
	}

	_, err = c.emitJmpImmediate(loopStart)

	if err != nil {
		return err
	}

	loopEnd := c.getCurrentOffset()
	c.patchJump(jmpEndPos, loopEnd)
	c.patchLoopBreaks(loopEnd)

	return nil
}

func (c *Compiler) compileRangeLoop(node *ast.ForStatement) error {
	var varAddr uint64
	varName := node.DeclaredVariable

	if varName == "" {
		varName = "__loop_counter"
	}

	currentScope := c.variableScopes[len(c.variableScopes)-1]
	var numVars uint64

	for _, scope := range c.variableScopes {
		numVars += uint64(len(scope))
	}

	varAddr = numVars
	currentScope[varName] = varAddr

	fromExpr := node.RangeFrom

	if !node.HasExplicitFrom {
		fromExpr = &ast.NumberLiteral{Value: "0", Range: node.Range}
	}

	err := c.compileNode(fromExpr)

	if err != nil {
		return err
	}

	fromRegister := c.getLastRegister()

	err = c.storeLoopVariable(varAddr, fromRegister)

	if err != nil {
		return err
	}

	loopStart := c.getCurrentOffset()

	varRegister, err := c.loadLoopVariable(varAddr)

	if err != nil {
		return err
	}

	err = c.compileNode(node.RangeTo)

	if err != nil {
		return err
	}

	toRegister := c.getLastRegister()
	destRegister := c.incrementRegCounter()
	err = c.compileComparison(
		destRegister,
		varRegister,
		toRegister,
		token.TokenTypeGreaterThan,
	)

	if err != nil {
		return err
	}

	jmpEndPos, err := c.emitJmpImmediateIfNotZero(destRegister, 0)

	if err != nil {
		return err
	}

	c.loopStack = append(c.loopStack, loopInfo{
		continueAddr:    c.getCurrentOffset(),
		breakAddr:       0,
		breakPatches:    make([]int, 0),
		continuePatches: make([]int, 0),
	})

	defer func() { c.loopStack = c.loopStack[:len(c.loopStack)-1] }()

	err = c.compileNode(node.Body)

	if err != nil {
		return err
	}

	incrementStart := c.getCurrentOffset()
	c.patchLoopContinues(incrementStart)

	err = c.incrementLoopVariable(varAddr)

	if err != nil {
		return err
	}

	_, err = c.emitJmpImmediate(loopStart)

	if err != nil {
		return err
	}

	loopEnd := c.getCurrentOffset()
	c.patchJump(jmpEndPos, loopEnd)
	c.patchLoopBreaks(loopEnd)

	return nil
}

// --- Helpers ---

func (c *Compiler) loadLoopVariable(addr uint64) (byte, error) {
	addrRegister := c.incrementRegCounter()
	err := c.emitLoadImmediate(addrRegister, int64(addr)) // #nosec: G115

	if err != nil {
		return 0, err
	}

	valRegister := c.incrementRegCounter()
	err = c.emitLoadMemory(valRegister, addrRegister)

	if err != nil {
		return 0, err
	}

	return valRegister, nil
}

func (c *Compiler) storeLoopVariable(addr uint64, valRegister byte) error {
	addrRegister := c.incrementRegCounter()
	err := c.emitLoadImmediate(addrRegister, int64(addr)) // #nosec: G115

	if err != nil {
		return err
	}

	return c.emitStoreMemory(valRegister, addrRegister)
}

func (c *Compiler) incrementLoopVariable(addr uint64) error {
	varRegister, err := c.loadLoopVariable(addr)

	if err != nil {
		return err
	}

	oneRegister := c.incrementRegCounter()

	err = c.emitLoadImmediate(oneRegister, 1)

	if err != nil {
		return err
	}

	incrementedRegister := c.incrementRegCounter()

	err = c.emitAdd(incrementedRegister, varRegister, oneRegister)

	if err != nil {
		return err
	}

	return c.storeLoopVariable(addr, incrementedRegister)
}

func (c *Compiler) patchJump(offset int, targetAddr uint64) {
	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, targetAddr)
	copy(c.bytecode[offset:offset+8], addrBytes)
}

func (c *Compiler) patchLoopBreaks(targetAddr uint64) {
	loopInfo := &c.loopStack[len(c.loopStack)-1]
	loopInfo.breakAddr = targetAddr
	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, targetAddr)

	for _, patchPos := range loopInfo.breakPatches {
		copy(c.bytecode[patchPos:patchPos+8], addrBytes)
	}
}

func (c *Compiler) patchLoopContinues(targetAddr uint64) {
	loopInfo := &c.loopStack[len(c.loopStack)-1]
	loopInfo.continueAddr = targetAddr
	addrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(addrBytes, targetAddr)

	for _, patchPos := range loopInfo.continuePatches {
		copy(c.bytecode[patchPos:patchPos+8], addrBytes)
	}
}
