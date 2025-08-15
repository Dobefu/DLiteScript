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
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgUndefinedFunction,
			fc.Position(),
			fc.FunctionName,
		)
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
			fc.Position(),
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
			fc.Position(),
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
			fc.Position(),
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
			fc.Position(),
			functionName,
			i+1,
			expectedKind.AsString(),
			argValues[i].DataType().AsString(),
		)
	}

	return argValues, nil
}
