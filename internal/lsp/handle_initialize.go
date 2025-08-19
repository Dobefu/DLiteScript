package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func (h *Handler) handleInitialize() (json.RawMessage, *jsonrpc2.Error) {
	result := lsptypes.InitializeResult{
		ServerInfo: lsptypes.ServerInfo{
			Name:    "DLiteScript",
			Version: "0.1.0",
		},
		Capabilities: lsptypes.ServerCapabilities{
			TextDocumentSync: lsptypes.TextDocumentSync{
				OpenClose: true,
				Change:    lsptypes.ChangeTypeIncremental,
			},
			DefinitionProvider: false,
			CompletionProvider: lsptypes.CompletionProvider{
				TriggerCharacters: []string{},
			},
			HoverProvider: true,
			SignatureHelpProvider: lsptypes.SignatureHelpProvider{
				TriggerCharacters: []string{"(", ","},
			},
		},
	}

	data, err := json.Marshal(result)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInternalError,
			err.Error(),
			nil,
		)
	}

	return data, nil
}
