package tokenizer

import (
	"strings"
	"unicode"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleIdentifier(firstChar rune) (*token.Token, error) {
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
