package tokenizer

import (
	"strings"
	"unicode"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

var keywords = map[string]token.Type{
	"var":    token.TokenTypeVar,
	"const":  token.TokenTypeConst,
	"number": token.TokenTypeTypeNumber,
	"string": token.TokenTypeTypeString,
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
			newToken, err = t.handleAsterisk()

		case '%':
			newToken = t.tokenPool.GetToken("%", token.TokenTypeOperationMod)

		case '/':
			newToken, err = t.handleDivisionOrComment()

		case '(':
			newToken = t.tokenPool.GetToken("(", token.TokenTypeLParen)

		case ')':
			newToken = t.tokenPool.GetToken(")", token.TokenTypeRParen)

		case '=':
			newToken = t.tokenPool.GetToken("=", token.TokenTypeAssign)

		case '{':
			newToken = t.tokenPool.GetToken("{", token.TokenTypeLBrace)

		case '}':
			newToken = t.tokenPool.GetToken("}", token.TokenTypeRBrace)

		case ',':
			newToken = t.tokenPool.GetToken(",", token.TokenTypeComma)

		case '"':
			newToken, err = t.handleString()

		default:
			newToken, err = t.parseUnknownChar(next)
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

func (t *Tokenizer) handleAsterisk() (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '*' {
		_, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		return t.tokenPool.GetToken("**", token.TokenTypeOperationPow), nil
	}

	return t.tokenPool.GetToken("*", token.TokenTypeOperationMul), nil
}

func (t *Tokenizer) handleDivisionOrComment() (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next != '/' {
		return t.tokenPool.GetToken("/", token.TokenTypeOperationDiv), nil
	}

	for !t.isEOF {
		next, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		if next == '\n' {
			break
		}
	}

	return nil, nil
}

func (t *Tokenizer) parseKeywordOrIdentifier(firstChar rune) (*token.Token, error) {
	var identifier strings.Builder
	identifier.WriteRune(firstChar)

	for !t.isEOF {
		next, err := t.Peek()

		if err != nil {
			break
		}

		if unicode.IsLetter(rune(next)) ||
			next == '_' ||
			unicode.IsDigit(rune(next)) {
			_, err = t.GetNext()

			if err != nil {
				return nil, err
			}

			identifier.WriteRune(next)

			continue
		}

		break
	}

	identifierText := identifier.String()

	if tokenType, isKeyword := keywords[identifierText]; isKeyword {
		return t.tokenPool.GetToken(identifierText, tokenType), nil
	}

	return t.tokenPool.GetToken(identifierText, token.TokenTypeIdentifier), nil
}

func (t *Tokenizer) parseUnknownChar(next rune) (*token.Token, error) {
	if unicode.IsLetter(rune(next)) || next == '_' {
		return t.parseKeywordOrIdentifier(rune(next))
	}

	return nil, errorutil.NewErrorAt(
		errorutil.ErrorMsgUnexpectedChar,
		t.expIdx,
		string(next),
	)
}

func (t *Tokenizer) handleString() (*token.Token, error) {
	var str strings.Builder
	str.Grow(16)

	var lastChar rune
	isEscaping := false

	for !t.isEOF {
		next, err := t.GetNext()

		if err != nil {
			return nil, err
		}

		if isEscaping {
			switch next {
			case 'n':
				str.WriteRune('\n')
			case 't':
				str.WriteRune('\t')
			case 'r':
				str.WriteRune('\r')
			case '0':
				str.WriteRune('\000')
			case 'b':
				str.WriteRune('\b')
			case 'f':
				str.WriteRune('\f')
			case 'v':
				str.WriteRune('\v')
			default:
				str.WriteRune(next)
			}

			isEscaping = false

			continue
		}

		if next == '"' {
			return t.tokenPool.GetToken(str.String(), token.TokenTypeString), nil
		}

		lastChar = next

		if lastChar == '\\' {
			isEscaping = true

			continue
		}

		str.WriteRune(next)
	}

	return nil, errorutil.NewErrorAt(errorutil.ErrorMsgUnexpectedEOF, t.byteIdx)
}
