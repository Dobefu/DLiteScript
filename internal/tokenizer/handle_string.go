package tokenizer

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleString(startPos int) (*token.Token, error) {
	var str strings.Builder
	str.Grow(16)

	var lastChar rune
	isEscaping := false

	for !t.isEOF {
		next, err := t.GetNext()

		if err != nil {
			return nil, err
		}

		if isEscaping {
			switch next {
			case 'n':
				str.WriteRune('\n')
			case 't':
				str.WriteRune('\t')
			case 'r':
				str.WriteRune('\r')
			case '0':
				str.WriteRune('\000')
			case 'b':
				str.WriteRune('\b')
			case 'f':
				str.WriteRune('\f')
			case 'v':
				str.WriteRune('\v')
			default:
				str.WriteRune(next)
			}

			isEscaping = false

			continue
		}

		if next == '"' {
			return token.NewToken(
				str.String(),
				token.TokenTypeString,
				startPos,
				t.expIdx,
			), nil
		}

		lastChar = next

		if lastChar == '\\' {
			isEscaping = true

			continue
		}

		str.WriteRune(next)
	}

	return nil, errorutil.NewErrorAt(errorutil.ErrorMsgUnexpectedEOF, t.byteIdx)
}
