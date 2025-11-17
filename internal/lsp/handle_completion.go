package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func (h *Handler) handleCompletion(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var completionParams lsptypes.CompletionParams
	err := json.Unmarshal(params, &completionParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	response := []lsptypes.CompletionItem{
		{
			Label: "TODO: Implement completion",
		},
	}

	data, err := json.Marshal(response)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInternalError,
			err.Error(),
			nil,
		)
	}

	return data, nil
}
