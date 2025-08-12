// Package datavalue provides the Value type and related methods.
package datavalue

import (
	"strconv"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// Value represents a value in the language.
type Value struct {
	dataType datatype.DataType

	Num  float64
	Str  string
	Bool bool
}

// DataType returns the data type of the value.
func (v Value) DataType() datatype.DataType {
	return v.dataType
}

// Null creates a new null value.
func Null() Value {
	return Value{
		dataType: datatype.DataTypeNull,

		Num:  0,
		Str:  "",
		Bool: false,
	}
}

// ToString returns the string representation of the value.
func (v Value) ToString() string {
	switch v.dataType {
	case datatype.DataTypeNull:
		return "null"

	case datatype.DataTypeNumber:
		return strconv.FormatFloat(v.Num, 'f', -1, 64)

	case datatype.DataTypeString:
		return v.Str

	case datatype.DataTypeBool:
		return strconv.FormatBool(v.Bool)

	default:
		return errorutil.ErrorMsgTypeUnknownDataType
	}
}

// Number creates a new number value.
func Number(n float64) Value {
	return Value{
		dataType: datatype.DataTypeNumber,

		Num:  n,
		Str:  "",
		Bool: false,
	}
}

// String creates a new string value.
func String(s string) Value {
	return Value{
		dataType: datatype.DataTypeString,

		Num:  0,
		Str:  s,
		Bool: false,
	}
}

// Bool creates a new boolean value.
func Bool(b bool) Value {
	return Value{
		dataType: datatype.DataTypeBool,

		Num:  0,
		Str:  "",
		Bool: b,
	}
}

// AsNumber returns the value as a number.
func (v Value) AsNumber() (float64, error) {
	if v.dataType != datatype.DataTypeNumber {
		return 0, errorutil.NewError(
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeNumber.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Num, nil
}

// AsString returns the value as a string.
func (v Value) AsString() (string, error) {
	if v.dataType != datatype.DataTypeString {
		return "", errorutil.NewError(
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeString.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Str, nil
}

// AsBool returns the value as a boolean.
func (v Value) AsBool() (bool, error) {
	if v.dataType != datatype.DataTypeBool {
		return false, errorutil.NewError(
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeBool.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Bool, nil
}
