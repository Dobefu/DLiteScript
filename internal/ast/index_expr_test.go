package ast

import (
	"testing"
)

func TestIndexExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "index expression",
			input: &IndexExpr{
				Array: &Identifier{
					Value: "array",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Index: &NumberLiteral{
					Value: "0",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    "array[0]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"array[0]", "array", "array", "0", "0"},
			continueOn:       "",
		},
		{
			name: "index expression with nil array",
			input: &IndexExpr{
				Array: nil,
				Index: &NumberLiteral{
					Value: "1",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"", "1", "1"},
			continueOn:       "",
		},
		{
			name: "index expression with nil index",
			input: &IndexExpr{
				Array: &Identifier{
					Value: "list",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 4, Line: 0, Column: 0},
					},
				},
				Index: nil,
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"", "list", "list"},
			continueOn:       "",
		},
		{
			name: "walk early return after index expression",
			input: &IndexExpr{
				Array: &Identifier{
					Value: "data",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 4, Line: 0, Column: 0},
					},
				},
				Index: &NumberLiteral{
					Value: "2",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    "data[2]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"data[2]"},
			continueOn:       "data[2]",
		},
		{
			name: "walk early return after array",
			input: &IndexExpr{
				Array: &Identifier{
					Value: "items",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Index: &NumberLiteral{
					Value: "3",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    "items[3]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"items[3]", "items"},
			continueOn:       "items",
		},
		{
			name: "walk early return after index",
			input: &IndexExpr{
				Array: &Identifier{
					Value: "values",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 6, Line: 0, Column: 0},
					},
				},
				Index: &NumberLiteral{
					Value: "4",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    "values[4]",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"values[4]", "values", "values", "4"},
			continueOn:       "4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Fatalf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Fatalf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.GetRange().Start.Offset,
				)
			}

			if test.input.GetRange().End.Offset != test.expectedEndPos {
				t.Fatalf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.GetRange().End.Offset,
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
