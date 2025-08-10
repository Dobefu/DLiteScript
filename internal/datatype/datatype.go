// Package datatype provides the DataType type and related methods.
package datatype

// DataType represents the type of a value.
type DataType int

const (
	// DataTypeNull represents a null type.
	DataTypeNull DataType = iota

	// DataTypeNumber represents a number type.
	DataTypeNumber
)

// AsString provides the string representation of a DataType for error messages.
func (dt DataType) AsString() string {
	switch dt {
	case DataTypeNull:
		return "null"

	case DataTypeNumber:
		return "number"

	default:
		return "unknown"
	}
}
