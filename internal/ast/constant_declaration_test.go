package ast

import (
	"testing"
)

func TestConstantDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *ConstantDeclaration
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "constant declaration with int type",
			input: &ConstantDeclaration{
				Name: "x",
				Type: "int",
				Value: &NumberLiteral{
					Value: "1",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"const x int = 1", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "constant declaration with string type",
			input: &ConstantDeclaration{
				Name: "y",
				Type: "string",
				Value: &StringLiteral{
					Value: "hello",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedNodes: []string{
				`const y string = "hello"`,
				`"hello"`,
				`"hello"`,
			},
			expectedStartPos: 0,
			expectedEndPos:   5,
			continueOn:       "",
		},
		{
			name: "constant declaration without value",
			input: &ConstantDeclaration{
				Name:  "z",
				Type:  "bool",
				Value: nil,
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{""},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after constant declaration",
			input: &ConstantDeclaration{
				Name: "w",
				Type: "float",
				Value: &NumberLiteral{
					Value: "3.14",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 4, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 4, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"const w float = 3.14"},
			expectedStartPos: 0,
			expectedEndPos:   4,
			continueOn:       "const w float = 3.14",
		},
		{
			name: "walk early return after value",
			input: &ConstantDeclaration{
				Name: "w",
				Type: "float",
				Value: &NumberLiteral{
					Value: "3.14",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 4, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 4, Line: 0, Column: 0},
				},
			},
			expectedNodes:    []string{"const w float = 3.14", "3.14"},
			expectedStartPos: 0,
			expectedEndPos:   4,
			continueOn:       "3.14",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
