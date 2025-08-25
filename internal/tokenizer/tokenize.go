package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

var keywords = map[string]token.Type{
	"var":      token.TokenTypeVar,
	"const":    token.TokenTypeConst,
	"number":   token.TokenTypeTypeNumber,
	"string":   token.TokenTypeTypeString,
	"bool":     token.TokenTypeTypeBool,
	"true":     token.TokenTypeBool,
	"false":    token.TokenTypeBool,
	"if":       token.TokenTypeIf,
	"else":     token.TokenTypeElse,
	"for":      token.TokenTypeFor,
	"break":    token.TokenTypeBreak,
	"continue": token.TokenTypeContinue,
	"from":     token.TokenTypeFrom,
	"to":       token.TokenTypeTo,
	"null":     token.TokenTypeNull,
	"func":     token.TokenTypeFunc,
	"return":   token.TokenTypeReturn,
}

// Tokenize analyzes the expression string and turns it into tokens.
func (t *Tokenizer) Tokenize() ([]*token.Token, error) {
	approxNumTokens := (t.expLen / 3)
	tokens := make([]*token.Token, 0, approxNumTokens)

	for !t.isEOF {
		startPos := t.expIdx
		next, err := t.GetNext()

		if err != nil {
			return tokens, err
		}

		var newToken *token.Token

		switch next {
		case ' ', '\t', '\r':
			continue

		case '\n':
			newToken = token.NewToken("\n", token.TokenTypeNewline, startPos, t.expIdx)

		case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newToken, err = t.handleNumberOrSpread(next, startPos)

		case '+':
			newToken = token.NewToken("+", token.TokenTypeOperationAdd, startPos, t.expIdx)

		case '-':
			newToken = token.NewToken("-", token.TokenTypeOperationSub, startPos, t.expIdx)

		case '*':
			newToken, err = t.handleAsteriskSign(startPos)

		case '%':
			newToken = token.NewToken("%", token.TokenTypeOperationMod, startPos, t.expIdx)

		case '/':
			newToken, err = t.handleSlashSign(startPos)

		case '(':
			newToken = token.NewToken("(", token.TokenTypeLParen, startPos, t.expIdx)

		case ')':
			newToken = token.NewToken(")", token.TokenTypeRParen, startPos, t.expIdx)

		case '=':
			newToken, err = t.handleEqualSign(startPos)

		case '!':
			newToken, err = t.handleExclamationSign(startPos)

		case '>':
			newToken, err = t.handleGreaterThanSign(startPos)

		case '<':
			newToken, err = t.handleLessThanSign(startPos)

		case '&':
			newToken, err = t.handleAmpersandSign(startPos)

		case '|':
			newToken, err = t.handlePipeSign(startPos)

		case '{':
			newToken = token.NewToken("{", token.TokenTypeLBrace, startPos, t.expIdx)

		case '}':
			newToken = token.NewToken("}", token.TokenTypeRBrace, startPos, t.expIdx)

		case ',':
			newToken = token.NewToken(",", token.TokenTypeComma, startPos, t.expIdx)

		case '"':
			newToken, err = t.handleString(startPos)

		default:
			newToken, err = t.handleUnknownChar(next, startPos)
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
