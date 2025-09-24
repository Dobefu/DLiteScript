package ast

import (
	"fmt"
	"testing"
)

func TestIsMultiline(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Range
		expected bool
	}{
		{
			name: "single line",
			input: Range{
				Start: Position{Offset: 0, Line: 0, Column: 0},
				End:   Position{Offset: 10, Line: 0, Column: 10},
			},
			expected: false,
		},
		{
			name: "multiple lines",
			input: Range{
				Start: Position{Offset: 0, Line: 0, Column: 0},
				End:   Position{Offset: 5, Line: 4, Column: 5},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := test.input.IsMultiLine()

			if test.expected != result {
				t.Errorf("expected multiline to be %t, got %t", test.expected, result)
			}
		})
	}
}

func TestLineSpan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Range
		expected int
	}{
		{
			name: "single line",
			input: Range{
				Start: Position{Offset: 0, Line: 0, Column: 0},
				End:   Position{Offset: 10, Line: 0, Column: 10},
			},
			expected: 1,
		},
		{
			name: "multiple lines",
			input: Range{
				Start: Position{Offset: 0, Line: 0, Column: 0},
				End:   Position{Offset: 5, Line: 4, Column: 5},
			},
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := test.input.LineSpan()

			if test.expected != result {
				t.Errorf(
					"expected range to span %d line(s), got %d",
					test.expected,
					result,
				)
			}
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Range
		expected string
	}{
		{
			name: "single line",
			input: Range{
				Start: Position{Offset: 0, Line: 0, Column: 0},
				End:   Position{Offset: 10, Line: 0, Column: 10},
			},
			expected: fmt.Sprintf("line %d at position %d", 1, 1),
		},
		{
			name: "multiple lines",
			input: Range{
				Start: Position{Offset: 0, Line: 0, Column: 0},
				End:   Position{Offset: 5, Line: 4, Column: 5},
			},
			expected: fmt.Sprintf(
				"from line %d at position %d to line %d at position %d",
				1,
				1,
				5,
				6,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := test.input.String()

			if test.expected != result {
				t.Errorf(
					"expected the range string to be \"%s\", got \"%s\"",
					test.expected,
					result,
				)
			}
		})
	}
}
