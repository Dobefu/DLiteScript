package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleEqualSign() (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '=' {
		_, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		return t.tokenPool.GetToken("==", token.TokenTypeEqual), nil
	}

	return t.tokenPool.GetToken("=", token.TokenTypeAssign), nil
}
