package token

// IsDataType checks if the token is a data type.
func (t *Token) IsDataType() bool {
	switch t.TokenType {
	case
		TokenTypeTypeNumber,
		TokenTypeTypeString,
		TokenTypeTypeBool,
		TokenTypeTypeArray,
		TokenTypeTypeError,
		TokenTypeTypeAny:
		return true

	default:
		return false
	}
}
