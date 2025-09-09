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
	// TokenTypeOperationSpread represents the spread operator.
	TokenTypeOperationSpread
	// TokenTypeNumber represents a number literal.
	TokenTypeNumber
	// TokenTypeString represents a string literal.
	TokenTypeString
	// TokenTypeBool represents a boolean value.
	TokenTypeBool
	// TokenTypeIdentifier represents an identifier.
	TokenTypeIdentifier
	// TokenTypeNull represents the null keyword.
	TokenTypeNull
	// TokenTypeLParen represents a left parenthesis.
	TokenTypeLParen
	// TokenTypeRParen represents a right parenthesis.
	TokenTypeRParen
	// TokenTypeFunction represents a function name.
	TokenTypeFunction
	// TokenTypeDot represents a single dot.
	TokenTypeDot
	// TokenTypeComma represents a comma separator.
	TokenTypeComma
	// TokenTypeNewline represents a newline separator.
	TokenTypeNewline
	// TokenTypeAssign represents the assignment operator.
	TokenTypeAssign
	// TokenTypeEqual represents the equality operator.
	TokenTypeEqual
	// TokenTypeNotEqual represents the inequality operator.
	TokenTypeNotEqual
	// TokenTypeGreaterThan represents the greater than operator.
	TokenTypeGreaterThan
	// TokenTypeGreaterThanOrEqual represents the greater than or equal to operator.
	TokenTypeGreaterThanOrEqual
	// TokenTypeLessThan represents the less than operator.
	TokenTypeLessThan
	// TokenTypeLessThanOrEqual represents the less than or equal to operator.
	TokenTypeLessThanOrEqual
	// TokenTypeLogicalAnd represents the logical and operator.
	TokenTypeLogicalAnd
	// TokenTypeLogicalOr represents the logical or operator.
	TokenTypeLogicalOr
	// TokenTypeNot represents the logical not operator.
	TokenTypeNot
	// TokenTypeIf represents the if keyword.
	TokenTypeIf
	// TokenTypeElse represents the else keyword.
	TokenTypeElse
	// TokenTypeLBrace represents a left brace.
	TokenTypeLBrace
	// TokenTypeRBrace represents a right brace.
	TokenTypeRBrace
	// TokenTypeLBracket represents a left bracket.
	TokenTypeLBracket
	// TokenTypeRBracket represents a right bracket.
	TokenTypeRBracket
	// TokenTypeVar represents the 'var' keyword.
	TokenTypeVar
	// TokenTypeConst represents the 'const' keyword.
	TokenTypeConst
	// TokenTypeFor represents the 'for' keyword.
	TokenTypeFor
	// TokenTypeBreak represents the 'break' keyword.
	TokenTypeBreak
	// TokenTypeContinue represents the 'continue' keyword.
	TokenTypeContinue
	// TokenTypeFrom represents the 'from' keyword.
	TokenTypeFrom
	// TokenTypeTo represents the 'to' keyword.
	TokenTypeTo
	// TokenTypeFunc represents the 'func' keyword.
	TokenTypeFunc
	// TokenTypeReturn represents the 'return' keyword.
	TokenTypeReturn

	// TokenTypeTypeNumber represents the 'number' type keyword.
	TokenTypeTypeNumber
	// TokenTypeTypeString represents the 'string' type keyword.
	TokenTypeTypeString
	// TokenTypeTypeBool represents the 'bool' type keyword.
	TokenTypeTypeBool
)
