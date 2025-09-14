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

		if unicode.IsLetter(rune(next)) ||
			next == '_' ||
			unicode.IsDigit(rune(next)) {
			_, _ = t.GetNext()
			identifier.WriteRune(next)

			continue
		}

		break
	}

	identifierText := identifier.String()

	if tokenType, isKeyword := keywords[identifierText]; isKeyword {
		return token.NewToken(identifierText, tokenType, startPos, t.expIdx), nil
	}

	return token.NewToken(
		identifierText,
		token.TokenTypeIdentifier,
		startPos,
		t.expIdx,
	), nil
}
