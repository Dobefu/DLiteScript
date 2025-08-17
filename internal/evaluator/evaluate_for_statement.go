package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateForStatement(
	node *ast.ForStatement,
) (*controlflow.EvaluationResult, error) {
	e.pushBlockScope()
	err := e.declareLoopVariable(node)

	if err != nil {
		e.popBlockScope()

		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	result := controlflow.NewRegularResult(datavalue.Null())

	for {
		shouldBreak, err := e.evaluateNodeCondition(node)

		if err != nil {
			e.popBlockScope()

			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		if shouldBreak {
			break
		}

		result, shouldBreak, shouldContinue, err := e.executeForIteration(node)

		if err != nil {
			e.popBlockScope()

			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		if shouldBreak {
			break
		}

		if shouldContinue {
			continue
		}

		if !result.IsNormalResult() {
			e.popBlockScope()

			return result, nil
		}
	}

	e.popBlockScope()

	return result, nil
}

func (e *Evaluator) declareLoopVariable(node *ast.ForStatement) error {
	if node.DeclaredVariable == "" {
		return nil
	}

	varName := node.DeclaredVariable
	initialValue := datavalue.Number(0)

	if node.IsRange {
		fromResult, err := e.Evaluate(node.RangeFrom)

		if err != nil {
			return err
		}

		initialValue = fromResult.Value
	}

	variable := &Variable{
		Value: initialValue,
		Type:  datatype.DataTypeNumber.AsString(),
	}

	if e.blockScopesLen > 0 {
		e.blockScopes[e.blockScopesLen-1][varName] = variable
	} else {
		e.outerScope[varName] = variable
	}

	return nil
}

func (e *Evaluator) evaluateNodeCondition(
	node *ast.ForStatement,
) (bool, error) {
	if !node.IsRange && node.Condition == nil {
		return false, nil
	}

	if !node.IsRange && node.Condition != nil {
		conditionResult, err := e.evaluateForCondition(node)

		if err != nil {
			return false, err
		}

		return !conditionResult, nil
	}

	if !node.IsRange {
		return false, errorutil.NewError(
			errorutil.ErrorMsgInvalidForStatement,
			node.StartPosition(),
			node.Condition.Expr(),
		)
	}

	varName := node.DeclaredVariable

	var currentVar ScopedValue
	var isVarFound bool

	if e.blockScopesLen > 0 {
		currentVar, isVarFound = e.blockScopes[e.blockScopesLen-1][varName]
	} else {
		currentVar, isVarFound = e.outerScope[varName]
	}

	if !isVarFound {
		return false, errorutil.NewError(
			errorutil.ErrorMsgVariableNotFound,
			varName,
		)
	}

	currentValue, err := currentVar.GetValue().AsNumber()

	if err != nil {
		return false, err
	}

	toResult, err := e.Evaluate(node.RangeTo)

	if err != nil {
		return false, err
	}

	toValue, err := toResult.Value.AsNumber()

	if err != nil {
		return false, err
	}

	return currentValue > toValue, nil
}

func (e *Evaluator) executeForIteration(
	node *ast.ForStatement,
) (*controlflow.EvaluationResult, bool, bool, error) {
	result, err := e.Evaluate(node.Body)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), false, false, err
	}

	shouldBreak, shouldContinue, propagatedResult, err := e.handleForControlFlowResult(result)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), false, false, err
	}

	err = e.incrementLoopVariable(node)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), false, false, err
	}

	if propagatedResult != nil {
		return propagatedResult, shouldBreak, shouldContinue, nil
	}

	return result, shouldBreak, shouldContinue, nil
}

func (e *Evaluator) evaluateForCondition(
	node *ast.ForStatement,
) (bool, error) {
	evaluatedCondition, err := e.Evaluate(node.Condition)

	if err != nil {
		return false, err
	}

	conditionResultBool, err := evaluatedCondition.Value.AsBool()

	if err != nil {
		return false, err
	}

	return conditionResultBool, nil
}

func (e *Evaluator) handleForControlFlowResult(
	result *controlflow.EvaluationResult,
) (shouldBreak bool, shouldContinue bool, propagatedResult *controlflow.EvaluationResult, err error) {
	if result.IsNormalResult() {
		return false, false, nil, nil
	}

	if result.IsBreakResult() {
		if result.Control.Count > 1 {
			return false, false, controlflow.NewBreakResult(result.Control.Count - 1), nil
		}

		return true, false, nil, nil
	}

	if result.IsContinueResult() {
		if result.Control.Count > 1 {
			return false, false, controlflow.NewContinueResult(result.Control.Count - 1), nil
		}

		return false, true, nil, nil
	}

	return false, false, nil, nil
}

func (e *Evaluator) incrementLoopVariable(node *ast.ForStatement) error {
	if node.DeclaredVariable == "" {
		return nil
	}

	var currentVar ScopedValue
	var isVarFound bool

	if e.blockScopesLen > 0 {
		currentVar, isVarFound = e.blockScopes[e.blockScopesLen-1][node.DeclaredVariable]
	} else {
		currentVar, isVarFound = e.outerScope[node.DeclaredVariable]
	}

	if !isVarFound {
		return errorutil.NewError(
			errorutil.ErrorMsgVariableNotFound,
			node.DeclaredVariable,
		)
	}

	currentValue, err := currentVar.GetValue().AsNumber()

	if err != nil {
		return err
	}

	newVarValue := &Variable{
		Value: datavalue.Number(currentValue + 1),
		Type:  currentVar.GetType(),
	}

	if e.blockScopesLen > 0 {
		e.blockScopes[e.blockScopesLen-1][node.DeclaredVariable] = newVarValue
	} else {
		e.outerScope[node.DeclaredVariable] = newVarValue
	}

	return nil
}
