package tokenizer

import "github.com/Dobefu/DLiteScript/internal/token"

func (t *Tokenizer) handleNumberOrSpread(next rune, startPos int) (*token.Token, error) {
	if next == '.' {
		afterNext, err := t.Peek()

		if err != nil {
			return nil, err
		}

		if afterNext == '.' {
			return t.parseSpread(startPos)
		}

		switch afterNext {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return t.parseNumber('.', startPos)

		default:
			return token.NewToken(".", token.TokenTypeDot, startPos, t.expIdx), nil
		}
	}

	return t.parseNumber(next, startPos)
}
