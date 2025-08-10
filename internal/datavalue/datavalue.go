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

	Num float64
}

// DataType returns the data type of the value.
func (v Value) DataType() datatype.DataType {
	return v.dataType
}

// Number creates a new number value.
func Number(n float64) Value {
	return Value{
		dataType: datatype.DataTypeNumber,
		Num:      n,
	}
}

// Null creates a new null value.
func Null() Value {
	return Value{
		dataType: datatype.DataTypeNull,
		Num:      0,
	}
}

// String returns the string representation of the value.
func (v Value) String() string {
	switch v.dataType {
	case datatype.DataTypeNumber:
		return strconv.FormatFloat(v.Num, 'f', -1, 64)

	case datatype.DataTypeNull:
		return "null"

	default:
		return "unknown type"
	}
}

// AsNumber returns the value as a number.
func (v Value) AsNumber() (float64, error) {
	if v.dataType != datatype.DataTypeNumber {
		return 0, errorutil.NewError(
			errorutil.ErrorMsgTypeExpectedNumber,
		)
	}

	return v.Num, nil
}
