package ast

// AnyLiteral defines a struct for a literal any value.
type AnyLiteral struct {
	Value ExprNode
	Range Range
}

// Expr returns the expression of the any literal.
func (e *AnyLiteral) Expr() string {
	return "any"
}

// GetRange returns the range of the any literal.
func (e *AnyLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the any literal.
func (e *AnyLiteral) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(e)

	if !shouldContinue {
		return
	}

	if e.Value != nil {
		shouldContinue = fn(e.Value)

		if !shouldContinue {
			return
		}

		e.Value.Walk(fn)
	}
}
