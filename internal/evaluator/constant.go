package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

// Constant defines a constant struct.
type Constant struct {
	Value datavalue.Value
	Type  string
}

// GetValue returns the value of the constant.
func (c *Constant) GetValue() datavalue.Value {
	return c.Value
}

// GetType returns the constant's type.
func (c *Constant) GetType() string {
	return c.Type
}

// IsConstant just returns true.
func (c *Constant) IsConstant() bool {
	return true
}
