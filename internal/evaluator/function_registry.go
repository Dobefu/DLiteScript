package evaluator

import (
	"math"
)

type functionHandler func([]float64) (float64, error)

type functionInfo struct {
	handler  functionHandler
	argCount int
}

var functionRegistry = map[string]functionInfo{
	"abs": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Abs(args[0]), nil
		},
	},
	"sin": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Sin(args[0]), nil
		},
	},
	"cos": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Cos(args[0]), nil
		},
	},
	"tan": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Tan(args[0]), nil
		},
	},
	"sqrt": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Sqrt(args[0]), nil
		},
	},
	"round": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Round(args[0]), nil
		},
	},
	"floor": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Floor(args[0]), nil
		},
	},
	"ceil": {
		argCount: 1,
		handler: func(args []float64) (float64, error) {
			return math.Ceil(args[0]), nil
		},
	},
	"min": {
		argCount: 2,
		handler: func(args []float64) (float64, error) {
			return math.Min(args[0], args[1]), nil
		},
	},
	"max": {
		argCount: 2,
		handler: func(args []float64) (float64, error) {
			return math.Max(args[0], args[1]), nil
		},
	},
}
