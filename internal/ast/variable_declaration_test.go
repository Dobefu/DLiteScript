package ast

import (
	"testing"
)

func TestVariableDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *VariableDeclaration
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "variable declaration with value",
			input: &VariableDeclaration{
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
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"var x int = 1", "1", "1"},
			continueOn:       "",
		},
		{
			name: "variable declaration without value",
			input: &VariableDeclaration{
				Name:  "x",
				Type:  "int",
				Value: nil,
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"var x int"},
			continueOn:       "",
		},
		{
			name: "walk early return after declaration node",
			input: &VariableDeclaration{
				Name: "y",
				Type: "string",
				Value: &StringLiteral{
					Value: "hello",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 2, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
			},
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"var y string = \"hello\""},
			continueOn:       "var y string = \"hello\"",
		},
		{
			name: "walk early return after value",
			input: &VariableDeclaration{
				Name: "y",
				Type: "string",
				Value: &StringLiteral{
					Value: "hello",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 2, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 2, Line: 0, Column: 0},
				},
			},
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"var y string = \"hello\"", "\"hello\""},
			continueOn:       "\"hello\"",
		},
		{
			name: "variable declaration with different type",
			input: &VariableDeclaration{
				Name: "z",
				Type: "bool",
				Value: &BoolLiteral{
					Value: "true",
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
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"var z bool = true", "true", "true"},
			continueOn:       "",
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
