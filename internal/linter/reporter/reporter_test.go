package reporter

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

type discardWriter struct{}

func (w *discardWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func TestReporter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		writer   io.Writer
		issues   []*Issue
		expected string
	}{
		{
			name:     "writer",
			writer:   &discardWriter{},
			issues:   []*Issue{},
			expected: "No issues found in test.dl\n",
		},
		{
			name:     "io.discard writer",
			writer:   io.Discard,
			issues:   []*Issue{},
			expected: "No issues found in test.dl\n",
		},
		{
			name:   "issues",
			writer: &discardWriter{},
			issues: []*Issue{
				{
					Rule:    "test",
					Message: "test",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
					Severity: SeverityWarning,
				},
			},
			expected: "Linting test.dl:\nwarning: test (test)\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			reporter := NewReporter(test.writer)

			for _, issue := range test.issues {
				reporter.AddIssue(issue)
			}

			if len(test.issues) > 0 && !reporter.HasIssues() {
				t.Fatalf("expected issues, got none")
			}

			issues := reporter.GetIssues()

			if len(issues) != len(test.issues) {
				t.Fatalf("expected 1 issue, got %d", len(issues))
			}

			reporter.PrintIssues("test.dl")
		})
	}
}
