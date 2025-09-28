package reporter

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

// Issue represents a linting issue found in the code.
type Issue struct {
	Rule     string    `json:"rule"`
	Message  string    `json:"message"`
	Range    ast.Range `json:"range"`
	Severity Severity  `json:"severity"`
}
