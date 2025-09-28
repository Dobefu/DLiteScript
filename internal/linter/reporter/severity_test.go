package reporter

import (
	"testing"
)

func TestSeverity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		severity Severity
		expected string
	}{
		{
			name:     "error",
			severity: SeverityError,
			expected: "error",
		},
		{
			name:     "warning",
			severity: SeverityWarning,
			expected: "warning",
		},
		{
			name:     "info",
			severity: SeverityInfo,
			expected: "info",
		},
		{
			name:     "unknown",
			severity: Severity(-1),
			expected: "unknown",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.severity.String() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, test.severity.String())
			}
		})
	}
}
