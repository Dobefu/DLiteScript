package tokenizer

import (
	"strings"
	"unicode"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

// Tokenize analyzes the expression string and turns it into tokens.
func (t *Tokenizer) Tokenize() ([]*token.Token, error) {
	approxNumTokens := (t.expLen / 3)
	tokens := make([]*token.Token, 0, approxNumTokens)

	for !t.isEOF {
		next, err := t.GetNext()

		if err != nil {
			return tokens, err
		}

		switch next {
		case ' ', '\t', '\r':
			continue

		case '\n':
			tokens = append(tokens, t.tokenPool.GetToken("\n", token.TokenTypeNewline))

			continue

		// Numeric characters.
		case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newToken, err := t.parseNumber(next)

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, newToken)

		case '+':
			tokens = append(tokens, t.tokenPool.GetToken(
				"+",
				token.TokenTypeOperationAdd,
			))

		case '-':
			tokens = append(tokens, t.tokenPool.GetToken(
				"-",
				token.TokenTypeOperationSub,
			))

		case '*':
			token, err := t.handleAsterisk()

			if err != nil {
				return nil, err
			}

			tokens = append(tokens, token)

		case '%':
			tokens = append(tokens, t.tokenPool.GetToken(
				"%",
				token.TokenTypeOperationMod,
			))

		case '/':
			tokens = append(tokens, t.tokenPool.GetToken(
				"/",
				token.TokenTypeOperationDiv,
			))

		case '(':
			tokens = append(tokens, t.tokenPool.GetToken(
				"(",
				token.TokenTypeLParen,
			))

		case ')':
			tokens = append(tokens, t.tokenPool.GetToken(
				")",
				token.TokenTypeRParen,
			))

		case ',':
			tokens = append(tokens, t.tokenPool.GetToken(
				",",
				token.TokenTypeComma,
			))

		default:
			newToken, err := t.parseUnknownChar(next)

			if err != nil {
				return nil, err
			}

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

func (t *Tokenizer) parseIdentifier(firstChar rune) (*token.Token, error) {
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

	return t.tokenPool.GetToken(identifier.String(), token.TokenTypeIdentifier), nil
}

func (t *Tokenizer) parseUnknownChar(next rune) (*token.Token, error) {
	if unicode.IsLetter(rune(next)) || next == '_' {
		return t.parseIdentifier(rune(next))
	}

	return nil, errorutil.NewErrorAt(
		errorutil.ErrorMsgUnexpectedChar,
		t.expIdx,
		string(next),
	)
}
