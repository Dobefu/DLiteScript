package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

func (h *Handler) handleDidClose(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var didCloseParams DidCloseParams
	err := json.Unmarshal(params, &didCloseParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	// Set the document to an empty string to indicate that it is closed.
	// The document is not deleted from the map, to prevent memory fragmentation.
	h.documents[didCloseParams.TextDocument.URI] = ""

	return json.RawMessage("null"), nil
}
