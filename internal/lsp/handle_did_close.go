package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func (h *Handler) handleDidClose(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var didCloseParams lsptypes.DidCloseParams
	err := json.Unmarshal(params, &didCloseParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	delete(h.documents, didCloseParams.TextDocument.URI)

	return json.RawMessage("null"), nil
}
