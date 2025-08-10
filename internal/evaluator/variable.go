package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

// Variable defines a variable struct.
type Variable struct {
	Value datavalue.Value
	Type  string
}

// GetValue returns the value of the variable.
func (v *Variable) GetValue() datavalue.Value {
	return v.Value
}

// GetType returns the variable's type.
func (v *Variable) GetType() string {
	return v.Type
}

// IsConstant just returns false.
func (v *Variable) IsConstant() bool {
	return false
}
