package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

var keywords = map[string]token.Type{
	"var":    token.TokenTypeVar,
	"const":  token.TokenTypeConst,
	"number": token.TokenTypeTypeNumber,
	"string": token.TokenTypeTypeString,
	"bool":   token.TokenTypeTypeBool,
	"true":   token.TokenTypeBool,
	"false":  token.TokenTypeBool,
	"if":     token.TokenTypeIf,
	"else":   token.TokenTypeElse,
	"for":    token.TokenTypeFor,
	"null":   token.TokenTypeNull,
}

// Tokenize analyzes the expression string and turns it into tokens.
func (t *Tokenizer) Tokenize() ([]*token.Token, error) {
	approxNumTokens := (t.expLen / 3)
	tokens := make([]*token.Token, 0, approxNumTokens)

	for !t.isEOF {
		next, err := t.GetNext()

		if err != nil {
			return tokens, err
		}

		var newToken *token.Token

		switch next {
		case ' ', '\t', '\r':
			continue

		case '\n':
			newToken = t.tokenPool.GetToken("\n", token.TokenTypeNewline)

		case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newToken, err = t.parseNumber(next)

		case '+':
			newToken = t.tokenPool.GetToken("+", token.TokenTypeOperationAdd)

		case '-':
			newToken = t.tokenPool.GetToken("-", token.TokenTypeOperationSub)

		case '*':
			newToken, err = t.handleAsteriskSign()

		case '%':
			newToken = t.tokenPool.GetToken("%", token.TokenTypeOperationMod)

		case '/':
			newToken, err = t.handleSlashSign()

		case '(':
			newToken = t.tokenPool.GetToken("(", token.TokenTypeLParen)

		case ')':
			newToken = t.tokenPool.GetToken(")", token.TokenTypeRParen)

		case '=':
			newToken, err = t.handleEqualSign()

		case '!':
			newToken, err = t.handleExclamationSign()

		case '>':
			newToken, err = t.handleGreaterThanSign()

		case '<':
			newToken, err = t.handleLessThanSign()

		case '&':
			newToken, err = t.handleAmpersandSign()

		case '|':
			newToken, err = t.handlePipeSign()

		case '{':
			newToken = t.tokenPool.GetToken("{", token.TokenTypeLBrace)

		case '}':
			newToken = t.tokenPool.GetToken("}", token.TokenTypeRBrace)

		case ',':
			newToken = t.tokenPool.GetToken(",", token.TokenTypeComma)

		case '"':
			newToken, err = t.handleString()

		default:
			newToken, err = t.handleUnknownChar(next)
		}

		if err != nil {
			return nil, err
		}

		if newToken != nil {
			tokens = append(tokens, newToken)
		}
	}

	return tokens, nil
}
