package token

func (t *Token) IsDataType() bool {
	switch t.TokenType {
	case
		TokenTypeTypeNumber,
		TokenTypeTypeString:
		return true

	default:
		return false
	}
}
