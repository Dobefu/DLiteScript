package evaluator

import (
	"math"

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
	result, err := e.findNamespaceFunction(fc)

	if result != nil || err != nil {
		return result, err
	}

	result, err = e.findRegistryFunction(fc)

	if result != nil || err != nil {
		return result, err
	}

	result, err = e.findUserFunction(fc)

	if result != nil || err != nil {
		return result, err
	}

	if fc.Namespace != "" {
		_, hasPkg := functionRegistry[fc.Namespace]

		if hasPkg {
			return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgUndefinedFunction,
				fc.GetRange().Start.Offset,
				fc.FunctionName,
			)
		}

		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgUndefinedNamespace,
			fc.GetRange().Start.Offset,
			fc.Namespace,
		)
	}

	return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgUndefinedFunction,
		fc.GetRange().Start.Offset,
		fc.FunctionName,
	)
}

func (e *Evaluator) findNamespaceFunction(fc *ast.FunctionCall) (*controlflow.EvaluationResult, error) {
	namespaceFunctions, hasNamespace := e.namespaceFunctions[fc.Namespace]

	if !hasNamespace {
		return nil, nil
	}

	userFunction, hasFunction := namespaceFunctions[fc.FunctionName]

	if !hasFunction {
		return nil, nil
	}

	return e.evaluateUserFunctionCall(fc, userFunction)
}

func (e *Evaluator) findRegistryFunction(fc *ast.FunctionCall) (*controlflow.EvaluationResult, error) {
	pkg, hasPkg := functionRegistry[fc.Namespace]

	if !hasPkg {
		return nil, nil
	}

	function, hasFunction := pkg.Functions[fc.FunctionName]

	if !hasFunction {
		return nil, nil
	}

	argValues, err := e.evaluateArguments(fc.Arguments, function, fc.FunctionName, fc)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	handlerResult, err := function.Handler(e, argValues)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	return controlflow.NewRegularResult(handlerResult), nil
}

func (e *Evaluator) findUserFunction(
	fc *ast.FunctionCall,
) (*controlflow.EvaluationResult, error) {
	userFunction, hasUserFunction := e.userFunctions[fc.FunctionName]

	if !hasUserFunction {
		return nil, nil
	}

	return e.evaluateUserFunctionCall(fc, userFunction)
}

func (e *Evaluator) evaluateUserFunctionCall(
	fc *ast.FunctionCall,
	userFunction *ast.FuncDeclarationStatement,
) (*controlflow.EvaluationResult, error) {
	if len(fc.Arguments) != len(userFunction.Args) {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgFunctionNumArgs,
			fc.GetRange().Start.Offset,
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
		err := e.validateReturnValues(result.Value, userFunction, fc)

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		if result.Value.DataType == datatype.DataTypeTuple {
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

		if !isSpreadArg {
			val, err := e.Evaluate(arg)

			if err != nil {
				return nil, err
			}

			argValues = append(argValues, val.Value)

			continue
		}

		spreadValue, err := e.Evaluate(spreadArg.Expression)

		if err != nil {
			return nil, err
		}

		if spreadValue.Value.DataType == datatype.DataTypeTuple {
			argValues = append(argValues, spreadValue.Value.Values...)

			continue
		}

		if spreadValue.Value.DataType == datatype.DataTypeArray {
			arrayValues, err := spreadValue.Value.AsArray()

			if err != nil {
				return nil, errorutil.NewErrorAt(
					errorutil.StageEvaluate,
					errorutil.ErrorMsgTypeExpected,
					spreadArg.GetRange().Start.Offset,
					"array",
					spreadValue.Value.DataType.AsString(),
				)
			}

			argValues = append(argValues, arrayValues...)

			continue
		}

		return nil, errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			spreadArg.GetRange().Start.Offset,
			"tuple or array",
			spreadValue.Value.DataType.AsString(),
		)
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
	expectedCount := len(functionInfo.Parameters)

	if len(argValues) != expectedCount {
		return nil, errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgFunctionNumArgs,
			fc.GetRange().Start.Offset,
			getFullFunctionName(fc),
			expectedCount,
			len(argValues),
		)
	}

	return e.validateArgTypesMatch(
		argValues,
		functionInfo.Parameters,
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
	if len(function.Parameters) == 0 {
		return argValues, nil
	}

	expectedType := function.Parameters[0]

	for i, arg := range argValues {
		if expectedType.Type == datatype.DataTypeAny {
			continue
		}

		if arg.DataType == expectedType.Type {
			continue
		}

		return nil, errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgFunctionArgType,
			fc.GetRange().Start.Offset,
			functionName,
			i+1,
			expectedType.Type.AsString(),
			arg.DataType.AsString(),
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
	if len(function.Parameters) == 0 {
		return argValues, nil
	}

	requiredParams := int(math.Max(float64(len(function.Parameters)-1), 0))

	if len(argValues) < requiredParams {
		return nil, errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgFunctionNumArgs,
			fc.GetRange().Start.Offset,
			getFullFunctionName(fc),
			requiredParams,
			len(argValues),
		)
	}

	_, err := e.validateArgTypesMatch(
		argValues[:requiredParams],
		function.Parameters[:requiredParams],
		functionName,
		fc,
	)

	if err != nil {
		return nil, err
	}

	if len(argValues) > requiredParams {
		variadicType := function.Parameters[len(function.Parameters)-1]

		for i := requiredParams; i < len(argValues); i++ {
			if variadicType.Type == datatype.DataTypeAny {
				continue
			}

			if argValues[i].DataType == variadicType.Type {
				continue
			}

			return nil, errorutil.NewErrorAt(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgFunctionArgType,
				fc.GetRange().Start.Offset,
				functionName,
				i+1,
				variadicType.Type.AsString(),
				argValues[i].DataType.AsString(),
			)
		}
	}

	return argValues, nil
}

func (e *Evaluator) validateArgTypesMatch(
	argValues []datavalue.Value,
	expectedTypes []function.ArgInfo,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	for i, expectedType := range expectedTypes {
		if expectedType.Type == datatype.DataTypeAny {
			continue
		}

		if argValues[i].DataType == expectedType.Type {
			continue
		}

		return nil, errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgFunctionArgType,
			fc.GetRange().Start.Offset,
			functionName,
			i+1,
			expectedType.Type.AsString(),
			argValues[i].DataType.AsString(),
		)
	}

	return argValues, nil
}

func (e *Evaluator) validateReturnValues(
	returnValue datavalue.Value,
	userFunction *ast.FuncDeclarationStatement,
	fc *ast.FunctionCall,
) error {
	expectedNumValues := userFunction.NumReturnValues

	if expectedNumValues > 1 {
		if returnValue.DataType != datatype.DataTypeTuple {
			return errorutil.NewErrorAt(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgFunctionReturnCount,
				fc.GetRange().Start.Offset,
				userFunction.Name,
				expectedNumValues,
				1,
			)
		}

		numValues := len(returnValue.Values)

		if numValues != expectedNumValues {
			return errorutil.NewErrorAt(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgFunctionReturnCount,
				fc.GetRange().Start.Offset,
				userFunction.Name,
				expectedNumValues,
				numValues,
			)
		}
	}

	return nil
}
