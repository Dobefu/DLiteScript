package ast

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestBinaryExpr(t *testing.T) {
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
			name: "empty binary expression",
			input: &BinaryExpr{
				Left:  nil,
				Right: nil,
				Operator: token.Token{
					Atom:      "",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes:    []string{""},
			continueOn:       "",
		},
		{
			name: "binary expression addition",
			input: &BinaryExpr{
				Left: &NumberLiteral{
					Value: "1",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "1",
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expectedValue:    "(1 + 1)",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes:    []string{"(1 + 1)", "1", "1", "1", "1"},
			continueOn:       "",
		},
		{
			name: "binary expression multiplication",
			input: &BinaryExpr{
				Left: &NumberLiteral{
					Value: "1",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "2",
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "*",
					TokenType: token.TokenTypeOperationMul,
					StartPos:  0,
					EndPos:    0,
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expectedValue:    "(1 * 2)",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes:    []string{"(1 * 2)", "1", "1", "2", "2"},
			continueOn:       "",
		},
		{
			name: "walk early return after binary expression node",
			input: &BinaryExpr{
				Left: &NumberLiteral{
					Value: "3",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "4",
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expectedValue:    "(3 - 4)",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes:    []string{"(3 - 4)"},
			continueOn:       "(3 - 4)",
		},
		{
			name: "walk early return after left operand",
			input: &BinaryExpr{
				Left: &NumberLiteral{
					Value: "3",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "4",
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expectedValue:    "(3 - 4)",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes:    []string{"(3 - 4)", "3"},
			continueOn:       "3",
		},
		{
			name: "walk early return after right operand",
			input: &BinaryExpr{
				Left: &NumberLiteral{
					Value: "3",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &NumberLiteral{
					Value: "4",
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "-",
					TokenType: token.TokenTypeOperationSub,
					StartPos:  0,
					EndPos:    0,
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expectedValue:    "(3 - 4)",
			expectedStartPos: 0,
			expectedEndPos:   0,
			expectedNodes:    []string{"(3 - 4)", "3", "3", "4"},
			continueOn:       "4",
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
