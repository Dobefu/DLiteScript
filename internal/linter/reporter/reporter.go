// Package reporter provides a reporter for linting issues.
package reporter

import (
	"fmt"
	"io"
	"log/slog"
)

// Reporter represents a reporter for linting issues.
type Reporter struct {
	issues  []*Issue
	outFile io.Writer
}

// NewReporter creates a new reporter.
func NewReporter(outFile io.Writer) *Reporter {
	return &Reporter{
		issues:  make([]*Issue, 0),
		outFile: outFile,
	}
}

// AddIssue adds an issue to the reporter.
func (r *Reporter) AddIssue(issue *Issue) {
	r.issues = append(r.issues, issue)
}

// GetIssues returns all issues.
func (r *Reporter) GetIssues() []*Issue {
	return r.issues
}

// HasIssues returns true if there are any issues.
func (r *Reporter) HasIssues() bool {
	return len(r.issues) > 0
}

// PrintIssues prints all issues to the output file.
func (r *Reporter) PrintIssues(filename string) {
	if r.outFile == io.Discard {
		return
	}

	if len(r.issues) == 0 {
		_, err := fmt.Fprintf(r.outFile, "No issues found in %s\n", filename)

		if err != nil {
			slog.Error(err.Error())
		}

		return
	}

	_, err := fmt.Fprintf(r.outFile, "Linting %s:\n", filename)

	if err != nil {
		slog.Error(err.Error())
	}

	for _, issue := range r.issues {
		pos := issue.Range.Start

		_, err = fmt.Fprintf(r.outFile, "%s:%d:%d: %s: %s (%s)\n",
			filename,
			pos.Line+1,
			pos.Column+1,
			issue.Severity.String(),
			issue.Message,
			issue.Rule,
		)

		if err != nil {
			slog.Error(err.Error())
		}
	}
}
