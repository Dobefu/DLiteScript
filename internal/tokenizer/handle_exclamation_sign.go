package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleExclamationSign() (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '=' {
		_, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		return t.tokenPool.GetToken("!=", token.TokenTypeNotEqual), nil
	}

	return t.tokenPool.GetToken("!", token.TokenTypeNot), nil
}
