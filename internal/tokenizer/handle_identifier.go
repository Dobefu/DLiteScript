package tokenizer

import (
	"strings"
	"unicode"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleIdentifier(
	firstChar rune,
	startPos int,
) (*token.Token, error) {
	var identifier strings.Builder
	identifier.WriteRune(firstChar)

	for !t.isEOF {
		next, err := t.Peek()

		if err != nil {
			return nil, err
		}

		if unicode.IsLetter(next) ||
			next == '_' ||
			unicode.IsDigit(next) {
			_, _ = t.GetNext()
			identifier.WriteRune(next)

			continue
		}

		break
	}

	identifierText := identifier.String()
	tokenType, isKeyword := keywords[identifierText]

	if isKeyword {
		return token.NewToken(identifierText, tokenType, startPos, t.expIdx), nil
	}

	return token.NewToken(
		identifierText,
		token.TokenTypeIdentifier,
		startPos,
		t.expIdx,
	), nil
}
