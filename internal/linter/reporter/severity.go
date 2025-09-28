package reporter

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
