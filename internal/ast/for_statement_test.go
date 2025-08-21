package ast

import "testing"

func TestForStatement(t *testing.T) {
	t.Parallel()

	statement := &ForStatement{
		DeclaredVariable: "",
		Condition:        nil,
		Body: &BlockStatement{
			Statements: []ExprNode{
				&NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
			},
			StartPos: 0,
			EndPos:   1,
		},
		StartPos:      0,
		EndPos:        1,
		RangeVariable: "",
		RangeFrom:     nil,
		RangeTo:       nil,
		IsRange:       false,
	}

	expectedNodes := []string{
		"for { (1) }",
		"(1)",
		"(1)",
		"1",
		"1",
	}

	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf("Expected %d visited nodes, got %d", len(expectedNodes), len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}

	if statement.StartPosition() != 0 {
		t.Fatalf("Expected start position to be 0, got %d", statement.StartPosition())
	}

	if statement.EndPosition() != 1 {
		t.Fatalf("Expected end position to be 1, got %d", statement.EndPosition())
	}
}

func TestForStatementCondition(t *testing.T) {
	t.Parallel()

	statement := &ForStatement{
		DeclaredVariable: "i",
		Condition: &BoolLiteral{
			Value:    "true",
			StartPos: 0,
			EndPos:   1,
		},
		Body: &BlockStatement{
			Statements: []ExprNode{
				&NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
			},
			StartPos: 0,
			EndPos:   1,
		},
		StartPos:      0,
		EndPos:        1,
		RangeVariable: "i",
		RangeFrom:     nil,
		RangeTo:       nil,
		IsRange:       false,
	}

	expectedNodes := []string{
		"for var i true { (1) }",
		"true",
		"true",
		"(1)",
		"(1)",
		"1",
		"1",
	}

	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf("Expected %d visited nodes, got %d", len(expectedNodes), len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}
}

func TestForStatementLiteralCondition(t *testing.T) {
	t.Parallel()

	statement := &ForStatement{
		DeclaredVariable: "",
		Condition: &BoolLiteral{
			Value:    "true",
			StartPos: 0,
			EndPos:   1,
		},
		Body: &BlockStatement{
			Statements: []ExprNode{
				&NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
			},
			StartPos: 0,
			EndPos:   1,
		},
		StartPos:      0,
		EndPos:        1,
		RangeVariable: "",
		RangeFrom:     nil,
		RangeTo:       nil,
		IsRange:       false,
	}

	expectedNodes := []string{
		"for true { (1) }",
		"true",
		"true",
		"(1)",
		"(1)",
		"1",
		"1",
	}

	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf("Expected %d visited nodes, got %d", len(expectedNodes), len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}
}

func TestForStatementRange(t *testing.T) {
	t.Parallel()

	statement := &ForStatement{
		DeclaredVariable: "i",
		Condition:        nil,
		Body: &BlockStatement{
			Statements: []ExprNode{
				&NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
			},
			StartPos: 0,
			EndPos:   1,
		},
		StartPos:      0,
		EndPos:        1,
		RangeVariable: "i",
		RangeFrom: &NumberLiteral{
			Value:    "0",
			StartPos: 0,
			EndPos:   1,
		},
		RangeTo: &NumberLiteral{
			Value:    "10",
			StartPos: 0,
			EndPos:   1,
		},
		IsRange: true,
	}

	expectedNodes := []string{
		"for var i from 0 to 10 { (1) }",
		"0",
		"0",
		"10",
		"10",
		"(1)",
		"(1)",
		"1",
		"1",
	}

	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf("Expected %d visited nodes, got %d", len(expectedNodes), len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}
}

func TestForStatementRangeTo(t *testing.T) {
	t.Parallel()

	statement := &ForStatement{
		DeclaredVariable: "i",
		Condition:        nil,
		Body: &BlockStatement{
			Statements: []ExprNode{
				&NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
			},
			StartPos: 0,
			EndPos:   1,
		},
		StartPos:      0,
		EndPos:        1,
		RangeVariable: "i",
		RangeFrom:     nil,
		RangeTo: &NumberLiteral{
			Value:    "10",
			StartPos: 0,
			EndPos:   1,
		},
		IsRange: true,
	}

	expectedNodes := []string{
		"for var i to 10 { (1) }",
		"10",
		"10",
		"(1)",
		"(1)",
		"1",
		"1",
	}

	visitedNodes := []string{}

	statement.Walk(func(node ExprNode) bool {
		visitedNodes = append(visitedNodes, node.Expr())

		return true
	})

	if len(visitedNodes) != len(expectedNodes) {
		t.Fatalf("Expected %d visited nodes, got %d", len(expectedNodes), len(visitedNodes))
	}

	for idx, node := range visitedNodes {
		if node != expectedNodes[idx] {
			t.Fatalf("Expected \"%s\", got \"%s\"", expectedNodes[idx], node)
		}
	}
}
