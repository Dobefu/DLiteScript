package evaluator

import (
	"math"

	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

type identifierInfo struct {
	handler func() (datavalue.Value, error)
}

var identifierRegistry = map[string]identifierInfo{
	"PI": {
		handler: func() (datavalue.Value, error) {
			return datavalue.Number(math.Pi), nil
		},
	},
	"TAU": {
		handler: func() (datavalue.Value, error) {
			return datavalue.Number(math.Pi * 2), nil
		},
	},
	"E": {
		handler: func() (datavalue.Value, error) {
			return datavalue.Number(math.E), nil
		},
	},
	"PHI": {
		handler: func() (datavalue.Value, error) {
			return datavalue.Number(math.Phi), nil
		},
	},
	"LN2": {
		handler: func() (datavalue.Value, error) {
			return datavalue.Number(math.Ln2), nil
		},
	},
	"LN10": {
		handler: func() (datavalue.Value, error) {
			return datavalue.Number(math.Ln10), nil
		},
	},
}
