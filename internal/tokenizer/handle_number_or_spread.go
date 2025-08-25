package tokenizer

import "github.com/Dobefu/DLiteScript/internal/token"

func (t *Tokenizer) handleNumberOrSpread(next rune, startPos int) (*token.Token, error) {
	afterNext, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if afterNext == '.' {
		return t.parseSpread(startPos)
	}

	return t.parseNumber(next, startPos)
}
