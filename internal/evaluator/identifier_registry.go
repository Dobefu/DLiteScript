package evaluator

import (
	"math"
)

type identifierInfo struct {
	handler func() (float64, error)
}

var identifierRegistry = map[string]identifierInfo{
	"PI": {
		handler: func() (float64, error) {
			return math.Pi, nil
		},
	},
	"TAU": {
		handler: func() (float64, error) {
			return math.Pi * 2, nil
		},
	},
	"E": {
		handler: func() (float64, error) {
			return math.E, nil
		},
	},
	"PHI": {
		handler: func() (float64, error) {
			return math.Phi, nil
		},
	},
	"LN2": {
		handler: func() (float64, error) {
			return math.Ln2, nil
		},
	},
	"LN10": {
		handler: func() (float64, error) {
			return math.Ln10, nil
		},
	},
}
