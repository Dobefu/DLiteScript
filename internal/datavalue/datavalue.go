// Package datavalue provides the Value type and related methods.
package datavalue

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// Value represents a value in the language.
type Value struct {
	DataType datatype.DataType

	Num    float64
	Str    string
	Bool   bool
	Func   *ast.FuncDeclarationStatement
	Values []Value
	Error  error
	Any    any
}

// Null creates a new null value.
func Null() Value {
	return Value{
		DataType: datatype.DataTypeNull,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: nil,
		Error:  nil,
		Any:    nil,
	}
}

// ToString returns the string representation of the value.
func (v Value) ToString() string {
	switch v.DataType {
	case
		datatype.DataTypeNull:
		return "null"

	case
		datatype.DataTypeNumber:
		return strconv.FormatFloat(v.Num, 'f', -1, 64)

	case
		datatype.DataTypeString:
		return v.Str

	case
		datatype.DataTypeBool:
		return strconv.FormatBool(v.Bool)

	case
		datatype.DataTypeFunction:
		return fmt.Sprintf("func %s", v.Func.Name)

	case
		datatype.DataTypeTuple:
		if len(v.Values) == 0 {
			return "()"
		}

		valueStrings := make([]string, len(v.Values))

		for i, val := range v.Values {
			valueStrings[i] = val.ToString()
		}

		return fmt.Sprintf("(%s)", strings.Join(valueStrings, ", "))

	case
		datatype.DataTypeArray:
		if len(v.Values) == 0 {
			return "[]"
		}

		valueStrings := make([]string, len(v.Values))

		for i, val := range v.Values {
			valueStrings[i] = val.ToString()
		}

		return fmt.Sprintf("[%s]", strings.Join(valueStrings, ", "))

	case
		datatype.DataTypeError:
		if v.Error == nil {
			return "null"
		}

		return v.Error.Error()

	case
		datatype.DataTypeAny:
		return fmt.Sprintf("%v", v.Any)

	default:
		return errorutil.ErrorMsgTypeUnknownDataType
	}
}

// Number creates a new number value.
func Number(n float64) Value {
	return Value{
		DataType: datatype.DataTypeNumber,

		Num:    n,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: nil,
		Error:  nil,
		Any:    nil,
	}
}

// String creates a new string value.
func String(s string) Value {
	return Value{
		DataType: datatype.DataTypeString,

		Num:    0,
		Str:    s,
		Bool:   false,
		Func:   nil,
		Values: nil,
		Error:  nil,
		Any:    nil,
	}
}

// Bool creates a new boolean value.
func Bool(b bool) Value {
	return Value{
		DataType: datatype.DataTypeBool,

		Num:    0,
		Str:    "",
		Bool:   b,
		Func:   nil,
		Values: nil,
		Error:  nil,
		Any:    nil,
	}
}

// Function creates a new function value.
func Function(fn *ast.FuncDeclarationStatement) Value {
	return Value{
		DataType: datatype.DataTypeFunction,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   fn,
		Values: nil,
		Error:  nil,
		Any:    nil,
	}
}

// Tuple creates a new tuple value.
func Tuple(values ...Value) Value {
	return Value{
		DataType: datatype.DataTypeTuple,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: values,
		Error:  nil,
		Any:    nil,
	}
}

// Array creates a new array value.
func Array(values ...Value) Value {
	return Value{
		DataType: datatype.DataTypeArray,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: values,
		Error:  nil,
		Any:    nil,
	}
}

// Error creates a new error value.
func Error(e error) Value {
	return Value{
		DataType: datatype.DataTypeError,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: nil,
		Error:  e,
		Any:    nil,
	}
}

// Any creates a new any value.
func Any(a any) Value {
	return Value{
		DataType: datatype.DataTypeAny,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: nil,
		Error:  nil,
		Any:    a,
	}
}

// AsNumber returns the value as a number.
func (v Value) AsNumber() (float64, error) {
	if v.DataType == datatype.DataTypeNumber {
		return v.Num, nil
	}

	if v.DataType == datatype.DataTypeAny {
		if v.Any == nil {
			return 0, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeNumber.AsString(),
				v.DataType.AsString(),
			)
		}

		switch a := v.Any.(type) {
		case float64:
			return a, nil

		case int:
			return float64(a), nil

		case int8:
			return float64(a), nil

		case int16:
			return float64(a), nil

		case int32:
			return float64(a), nil

		case int64:
			return float64(a), nil

		case uint:
			return float64(a), nil

		case uint8:
			return float64(a), nil

		case uint16:
			return float64(a), nil

		case uint32:
			return float64(a), nil

		case uint64:
			return float64(a), nil

		case float32:
			return float64(a), nil

		default:
			return 0, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeNumber.AsString(),
				v.DataType.AsString(),
			)
		}
	}

	return 0, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeNumber.AsString(),
		v.DataType.AsString(),
	)
}

// AsString returns the value as a string.
func (v Value) AsString() (string, error) {
	if v.DataType == datatype.DataTypeString {
		return v.Str, nil
	}

	if v.DataType == datatype.DataTypeAny {
		if v.Any == nil {
			return "", errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeString.AsString(),
				v.DataType.AsString(),
			)
		}

		str, isString := v.Any.(string)

		if isString {
			return str, nil
		}

		return "", errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeString.AsString(),
			v.DataType.AsString(),
		)
	}

	return "", errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeString.AsString(),
		v.DataType.AsString(),
	)
}

// AsBool returns the value as a boolean.
func (v Value) AsBool() (bool, error) {
	if v.DataType == datatype.DataTypeBool {
		return v.Bool, nil
	}

	if v.DataType == datatype.DataTypeAny {
		if v.Any == nil {
			return false, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeBool.AsString(),
				v.DataType.AsString(),
			)
		}

		boolean, isBoolean := v.Any.(bool)

		if isBoolean {
			return boolean, nil
		}

		return false, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeBool.AsString(),
			v.DataType.AsString(),
		)
	}

	return false, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeBool.AsString(),
		v.DataType.AsString(),
	)
}

// AsFunction returns the value as a function.
func (v Value) AsFunction() (*ast.FuncDeclarationStatement, error) {
	if v.DataType == datatype.DataTypeFunction {
		return v.Func, nil
	}

	if v.DataType == datatype.DataTypeAny {
		if v.Any == nil {
			return nil, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeFunction.AsString(),
				v.DataType.AsString(),
			)
		}

		function, isFunction := v.Any.(*ast.FuncDeclarationStatement)

		if isFunction {
			return function, nil
		}

		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeFunction.AsString(),
			v.DataType.AsString(),
		)
	}

	return nil, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeFunction.AsString(),
		v.DataType.AsString(),
	)
}

// AsArray returns the value as an array.
func (v Value) AsArray() ([]Value, error) {
	if v.DataType == datatype.DataTypeArray {
		return v.Values, nil
	}

	if v.DataType == datatype.DataTypeAny {
		if v.Any == nil {
			return nil, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeArray.AsString(),
				v.DataType.AsString(),
			)
		}

		array, isArray := v.Any.([]Value)

		if isArray {
			return array, nil
		}

		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeArray.AsString(),
			v.DataType.AsString(),
		)
	}

	return nil, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeArray.AsString(),
		v.DataType.AsString(),
	)
}

// AsTuple returns the value as a tuple.
func (v Value) AsTuple() ([]Value, error) {
	if v.DataType == datatype.DataTypeTuple {
		return v.Values, nil
	}

	if v.DataType == datatype.DataTypeAny {
		if v.Any == nil {
			return nil, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeTuple.AsString(),
				v.DataType.AsString(),
			)
		}

		tuple, isTuple := v.Any.([]Value)

		if isTuple {
			return tuple, nil
		}

		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeTuple.AsString(),
			v.DataType.AsString(),
		)
	}

	return nil, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeTuple.AsString(),
		v.DataType.AsString(),
	)
}

// AsError returns the value as an error.
func (v Value) AsError() (error, error) {
	if v.DataType == datatype.DataTypeError {
		return v.Error, nil
	}

	return nil, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeError.AsString(),
		v.DataType.AsString(),
	)
}

// Equals returns if two values are equal.
func (v Value) Equals(other Value) bool {
	if v.DataType != other.DataType {
		return false
	}

	switch v.DataType {
	case
		datatype.DataTypeNull:
		return true

	case
		datatype.DataTypeNumber:
		return v.Num == other.Num

	case
		datatype.DataTypeString:
		return v.Str == other.Str

	case
		datatype.DataTypeBool:
		return v.Bool == other.Bool

	case
		datatype.DataTypeFunction:
		return v.Func == other.Func

	case
		datatype.DataTypeTuple,
		datatype.DataTypeArray:
		if len(v.Values) != len(other.Values) {
			return false
		}

		for i, val := range v.Values {
			if !val.Equals(other.Values[i]) {
				return false
			}
		}

		return true

	case
		datatype.DataTypeError:
		if v.Error == nil {
			return other.Error == nil
		}

		if other.Error == nil {
			return v.Error == nil
		}

		return v.Error.Error() == other.Error.Error()

	case
		datatype.DataTypeAny:
		return true

	default:
		return false
	}
}

// IsTruthy checks if the provided value is truthy.
func (v Value) IsTruthy() bool {
	switch v.DataType {
	case
		datatype.DataTypeBool:
		boolVal, _ := v.AsBool() // This cannot return an error.

		return boolVal

	case
		datatype.DataTypeNumber:
		numVal, _ := v.AsNumber() // This cannot return an error.

		return numVal != 0

	case
		datatype.DataTypeString:
		strVal, _ := v.AsString() // This cannot return an error.

		return strVal != ""

	case
		datatype.DataTypeNull:
		return false

	case
		datatype.DataTypeAny:
		return v.Any != nil

	case
		datatype.DataTypeArray:
		return len(v.Values) > 0

	case
		datatype.DataTypeFunction,
		datatype.DataTypeTuple,
		datatype.DataTypeError:
		return true

	default:
		return false
	}
}
