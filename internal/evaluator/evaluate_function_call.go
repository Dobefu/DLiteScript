package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/function"
	"github.com/Dobefu/DLiteScript/internal/stdlib"
)

var functionRegistry = stdlib.GetFunctionRegistry()

func (e *Evaluator) evaluateFunctionCall(
	fc *ast.FunctionCall,
) (*controlflow.EvaluationResult, error) {
	_, hasNamespace := functionRegistry[fc.Namespace]

	if !hasNamespace {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgUndefinedNamespace,
			fc.StartPosition(),
			fc.Namespace,
		)
	}

	function, hasFunction := functionRegistry[fc.Namespace].Functions[fc.FunctionName]

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

	handlerResult, err := function.Handler(
		e,
		argValues,
	)

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
			getFullFunctionName(fc),
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
		if result.Value.DataType() == datatype.DataTypeTuple {
			return result, nil
		}

		return result, nil
	}

	return result, nil
}

func (e *Evaluator) evaluateArguments(
	args []ast.ExprNode,
	function function.Info,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	argValues := make([]datavalue.Value, 0, len(args))

	for _, arg := range args {
		spreadArg, isSpreadArg := arg.(*ast.SpreadExpr)

		if isSpreadArg {
			spreadValue, err := e.Evaluate(spreadArg.Expression)

			if err != nil {
				return nil, err
			}

			if spreadValue.Value.DataType() == datatype.DataTypeTuple {
				argValues = append(argValues, spreadValue.Value.Values...)

				continue
			}

			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgTypeExpected,
				spreadArg.StartPosition(),
				datatype.DataTypeTuple.AsString(),
				spreadValue.Value.DataType().AsString(),
			)
		}

		val, err := e.Evaluate(arg)

		if err != nil {
			return nil, err
		}

		argValues = append(argValues, val.Value)
	}

	return e.validateArgumentTypes(argValues, function, functionName, fc)
}

func (e *Evaluator) validateArgumentTypes(
	argValues []datavalue.Value,
	functionInfo function.Info,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	switch functionInfo.FunctionType {
	case function.FunctionTypeFixed:
		return e.validateFixedArgs(argValues, functionInfo, functionName, fc)

	case function.FunctionTypeVariadic:
		return e.validateVariadicArgs(argValues, functionInfo, functionName, fc)

	case function.FunctionTypeMixedVariadic:
		return e.validateMixedVariadicArgs(argValues, functionInfo, functionName, fc)

	default:
		return argValues, nil
	}
}

func (e *Evaluator) validateFixedArgs(
	argValues []datavalue.Value,
	functionInfo function.Info,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	expectedCount := len(functionInfo.ArgKinds)

	if len(argValues) != expectedCount {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionNumArgs,
			fc.StartPosition(),
			getFullFunctionName(fc),
			expectedCount,
			len(argValues),
		)
	}

	return e.validateArgTypesMatch(
		argValues,
		functionInfo.ArgKinds,
		functionName,
		fc,
	)
}

func (e *Evaluator) validateVariadicArgs(
	argValues []datavalue.Value,
	function function.Info,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	if len(function.ArgKinds) == 0 {
		return argValues, nil
	}

	expectedType := function.ArgKinds[0]

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
	function function.Info,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	minRequired := len(function.ArgKinds)

	if len(argValues) < minRequired {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionNumArgs,
			fc.StartPosition(),
			getFullFunctionName(fc),
			minRequired,
			len(argValues),
		)
	}

	_, err := e.validateArgTypesMatch(
		argValues[:minRequired],
		function.ArgKinds,
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
