package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateNullLiteral() (datavalue.Value, error) {
	return datavalue.Null(), nil
}
