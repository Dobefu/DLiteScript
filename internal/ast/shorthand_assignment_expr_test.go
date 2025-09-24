package ast

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestShorthandAssignmentExpr(t *testing.T) {
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
			name: "shorthand assignment with addition",
			input: &ShorthandAssignmentExpr{
				Left: &Identifier{
					Value: "x",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "1",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "x += 1",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"x += 1", "x", "x", "1", "1"},
			continueOn:       "",
		},
		{
			name: "shorthand assignment with subtraction",
			input: &ShorthandAssignmentExpr{
				Left: &Identifier{
					Value: "y",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "2",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("-=", token.TokenTypeOperationSubAssign, 0, 1),
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "y -= 2",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"y -= 2", "y", "y", "2", "2"},
			continueOn:       "",
		},
		{
			name: "shorthand assignment with nil left",
			input: &ShorthandAssignmentExpr{
				Left: nil,
				Right: &NumberLiteral{
					Value: "3",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("*=", token.TokenTypeOperationMulAssign, 0, 1),
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"", "3", "3"},
			continueOn:       "",
		},
		{
			name: "shorthand assignment with nil right",
			input: &ShorthandAssignmentExpr{
				Left: &Identifier{
					Value: "z",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right:    nil,
				Operator: *token.NewToken("/=", token.TokenTypeOperationDivAssign, 0, 1),
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"", "z", "z"},
			continueOn:       "",
		},
		{
			name: "walk early return after shorthand assignment",
			input: &ShorthandAssignmentExpr{
				Left: &Identifier{
					Value: "a",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "5",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("%=", token.TokenTypeOperationModAssign, 0, 1),
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "a %= 5",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"a %= 5"},
			continueOn:       "a %= 5",
		},
		{
			name: "walk early return after left operand",
			input: &ShorthandAssignmentExpr{
				Left: &Identifier{
					Value: "b",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "6",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("**=", token.TokenTypeOperationPowAssign, 0, 1),
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "b **= 6",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"b **= 6", "b"},
			continueOn:       "b",
		},
		{
			name: "walk early return after right operand",
			input: &ShorthandAssignmentExpr{
				Left: &Identifier{
					Value: "c",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "7",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expectedValue:    "c += 7",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"c += 7", "c", "c", "7"},
			continueOn:       "7",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Errorf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.GetRange().Start.Offset,
				)
			}

			if test.input.GetRange().End.Offset != test.expectedEndPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.GetRange().End.Offset,
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
