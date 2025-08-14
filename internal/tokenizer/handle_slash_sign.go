package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleSlashSign() (*token.Token, error) {
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
