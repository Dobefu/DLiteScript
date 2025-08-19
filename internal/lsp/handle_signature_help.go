package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func (h *Handler) handleSignatureHelp(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var signatureHelpParams lsptypes.SignatureHelpParams
	err := json.Unmarshal(params, &signatureHelpParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	response := lsptypes.SignatureHelp{
		Signatures: []lsptypes.SignatureInformation{
			{
				Label: "TODO: Implement signature help",
				Documentation: &lsptypes.MarkupContent{
					Kind:  "plaintext",
					Value: "",
				},
				Parameters: []lsptypes.ParameterInformation{},
			},
		},
		ActiveSignature: 0,
		ActiveParameter: 0,
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
