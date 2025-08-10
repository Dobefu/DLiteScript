package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

// ScopedValue defines an interface for scoped values.
type ScopedValue interface {
	GetValue() datavalue.Value
	GetType() string
	IsConstant() bool
}
