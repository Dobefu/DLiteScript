// Package datatype provides the DataType type and related methods.
package datatype

// DataType represents the type of a value.
type DataType int

const (
	// DataTypeNull represents a null type.
	DataTypeNull DataType = iota
	// DataTypeNumber represents a number type.
	DataTypeNumber
	// DataTypeString represents a string type.
	DataTypeString
	// DataTypeBool represents a boolean type.
	DataTypeBool
	// DataTypeFunction represents a function type.
	DataTypeFunction
	// DataTypeTuple represents a tuple type.
	DataTypeTuple
)

// AsString provides the string representation of a DataType for error messages.
func (dt DataType) AsString() string {
	switch dt {
	case DataTypeNull:
		return "null"

	case DataTypeNumber:
		return "number"

	case DataTypeString:
		return "string"

	case DataTypeBool:
		return "bool"

	case DataTypeFunction:
		return "function"

	case DataTypeTuple:
		return "tuple"

	default:
		return "unknown"
	}
}
