package ast

import (
	"testing"
)

func TestStringLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *StringLiteral
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "string literal",
			input: &StringLiteral{
				Value: "test",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    `"test"`,
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{`"test"`},
			continueOn:       "",
		},
		{
			name: "walk early return after string node",
			input: &StringLiteral{
				Value: "hello",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
			},
			expectedValue:    `"hello"`,
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{`"hello"`},
			continueOn:       `"hello"`,
		},
		{
			name: "empty string literal",
			input: &StringLiteral{
				Value: "",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    `""`,
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{`""`},
			continueOn:       "",
		},
		{
			name: "string literal with special characters",
			input: &StringLiteral{
				Value: "hello\nworld",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 3, Line: 0, Column: 0},
				},
			},
			expectedValue:    `"hello\nworld"`,
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{`"hello\nworld"`},
			continueOn:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Fatalf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
			}

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedStartPos,
					test.input.GetRange().Start.Offset,
				)
			}

			if test.input.GetRange().End.Offset != test.expectedEndPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedEndPos,
					test.input.GetRange().End.Offset,
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
