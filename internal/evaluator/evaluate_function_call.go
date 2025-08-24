package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateFunctionCall(
	fc *ast.FunctionCall,
) (*controlflow.EvaluationResult, error) {
	function, hasFunction := functionRegistry[fc.FunctionName]

	if !hasFunction {
		userFunction, hasUserFunction := e.userFunctions[fc.FunctionName]

		if !hasUserFunction {
			return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
				errorutil.ErrorMsgUndefinedFunction,
				fc.StartPosition(),
				fc.FunctionName,
			)
		}

		return e.evaluateUserFunctionCall(fc, userFunction)
	}

	argValues, err := e.evaluateArguments(
		fc.Arguments,
		function,
		fc.FunctionName,
		fc,
	)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	handlerResult, err := function.handler(e, argValues)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	return controlflow.NewRegularResult(handlerResult), nil
}

func (e *Evaluator) evaluateUserFunctionCall(
	fc *ast.FunctionCall,
	userFunction *ast.FuncDeclarationStatement,
) (*controlflow.EvaluationResult, error) {
	if len(fc.Arguments) != len(userFunction.Args) {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionNumArgs,
			fc.StartPosition(),
			fc.FunctionName,
			len(userFunction.Args),
			len(fc.Arguments),
		)
	}

	argValues := make([]datavalue.Value, len(fc.Arguments))

	for i, arg := range fc.Arguments {
		val, err := e.Evaluate(arg)

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		argValues[i] = val.Value
	}

	e.pushBlockScope()
	defer e.popBlockScope()

	for i, param := range userFunction.Args {
		e.blockScopes[e.blockScopesLen-1][param.Name] = &Variable{
			Value: argValues[i],
			Type:  param.Type,
		}
	}

	result, err := e.Evaluate(userFunction.Body)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if result.IsReturnResult() {
		return result, nil
	}

	return result, nil
}

func (e *Evaluator) evaluateArguments(
	args []ast.ExprNode,
	function functionInfo,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	argValues := make([]datavalue.Value, len(args))

	for i, arg := range args {
		val, err := e.Evaluate(arg)

		if err != nil {
			return nil, err
		}

		argValues[i] = val.Value
	}

	return e.validateArgumentTypes(argValues, function, functionName, fc)
}

func (e *Evaluator) validateArgumentTypes(
	argValues []datavalue.Value,
	function functionInfo,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	switch function.functionType {
	case functionTypeFixed:
		return e.validateFixedArgs(argValues, function, functionName, fc)

	case functionTypeVariadic:
		return e.validateVariadicArgs(argValues, function, functionName, fc)

	case functionTypeMixedVariadic:
		return e.validateMixedVariadicArgs(argValues, function, functionName, fc)

	default:
		return argValues, nil
	}
}

func (e *Evaluator) validateFixedArgs(
	argValues []datavalue.Value,
	function functionInfo,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	expectedCount := len(function.argKinds)

	if len(argValues) != expectedCount {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionNumArgs,
			fc.StartPosition(),
			functionName,
			expectedCount,
			len(argValues),
		)
	}

	return e.validateArgTypesMatch(argValues, function.argKinds, functionName, fc)
}

func (e *Evaluator) validateVariadicArgs(
	argValues []datavalue.Value,
	function functionInfo,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	if len(function.argKinds) == 0 {
		return argValues, nil
	}

	expectedType := function.argKinds[0]

	for i, arg := range argValues {
		if arg.DataType() == expectedType {
			continue
		}

		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionArgType,
			fc.StartPosition(),
			functionName,
			i+1,
			expectedType.AsString(),
			arg.DataType().AsString(),
		)
	}

	return argValues, nil
}

func (e *Evaluator) validateMixedVariadicArgs(
	argValues []datavalue.Value,
	function functionInfo,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	minRequired := len(function.argKinds)

	if len(argValues) < minRequired {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionNumArgs,
			fc.StartPosition(),
			functionName,
			minRequired,
			len(argValues),
		)
	}

	_, err := e.validateArgTypesMatch(
		argValues[:minRequired],
		function.argKinds,
		functionName,
		fc,
	)

	if err != nil {
		return nil, err
	}

	return argValues, nil
}

func (e *Evaluator) validateArgTypesMatch(
	argValues []datavalue.Value,
	expectedTypes []datatype.DataType,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	for i, expectedKind := range expectedTypes {
		if argValues[i].DataType() == expectedKind {
			continue
		}

		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionArgType,
			fc.StartPosition(),
			functionName,
			i+1,
			expectedKind.AsString(),
			argValues[i].DataType().AsString(),
		)
	}

	return argValues, nil
}
