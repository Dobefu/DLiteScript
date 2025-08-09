package token

// Type defines an enum of possible token types.
type Type int

const (
	// TokenTypeOperationAdd represents the addition operation.
	TokenTypeOperationAdd = iota
	// TokenTypeOperationSub represents the subtraction operation.
	TokenTypeOperationSub
	// TokenTypeOperationMul represents the multiplication operation.
	TokenTypeOperationMul
	// TokenTypeOperationDiv represents the division operation.
	TokenTypeOperationDiv
	// TokenTypeOperationPow represents the power operation.
	TokenTypeOperationPow
	// TokenTypeOperationMod represents the modulo operation.
	TokenTypeOperationMod
	// TokenTypeNumber represents a number literal.
	TokenTypeNumber
	// TokenTypeIdentifier represents an identifier.
	TokenTypeIdentifier
	// TokenTypeLParen represents a left parenthesis.
	TokenTypeLParen
	// TokenTypeRParen represents a right parenthesis.
	TokenTypeRParen
	// TokenTypeFunction represents a function name.
	TokenTypeFunction
	// TokenTypeComma represents a comma separator.
	TokenTypeComma
	// TokenTypeNewline represents a newline separator.
	TokenTypeNewline
)
