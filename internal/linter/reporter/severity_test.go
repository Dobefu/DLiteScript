package reporter

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
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

			if test.severity.ToErrorutilStage() != errorutil.StageParse {
				t.Fatalf(
					"expected %s, got %s",
					errorutil.StageParse,
					test.severity.ToErrorutilStage(),
				)
			}
		})
	}
}
