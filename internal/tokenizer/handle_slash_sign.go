package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleSlashSign(startPos int) (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next != '/' {
		return token.NewToken(
			"/",
			token.TokenTypeOperationDiv,
			startPos,
			t.expIdx,
		), nil
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
