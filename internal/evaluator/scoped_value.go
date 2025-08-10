package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

type ScopedValue interface {
	GetValue() datavalue.Value
	GetType() string
	IsConstant() bool
}
