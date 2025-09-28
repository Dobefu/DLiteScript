package reporter

import "github.com/Dobefu/DLiteScript/internal/errorutil"

// Severity represents the severity level of a linting issue.
type Severity int

const (
	// SeverityError represents an error-level issue.
	SeverityError Severity = iota
	// SeverityWarning represents a warning-level issue.
	SeverityWarning
	// SeverityInfo represents an info-level issue.
	SeverityInfo
)

// String returns the string representation of the severity.
func (s Severity) String() string {
	switch s {
	case SeverityError:
		return "error"

	case SeverityWarning:
		return "warning"

	case SeverityInfo:
		return "info"

	default:
		return "unknown"
	}
}

// ToErrorutilStage converts severity to errorutil stage for consistent error reporting.
func (s Severity) ToErrorutilStage() errorutil.Stage {
	switch s {
	case SeverityError:
		return errorutil.StageParse

	case SeverityWarning, SeverityInfo:
		return errorutil.StageParse

	default:
		return errorutil.StageParse
	}
}
