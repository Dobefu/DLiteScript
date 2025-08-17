package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

func (h *Handler) handleInitialize() (json.RawMessage, *jsonrpc2.Error) {
	result := InitializeResult{
		ServerInfo: ServerInfo{
			Name:    "DLiteScript",
			Version: "0.1.0",
		},
		Capabilities: ServerCapabilities{
			TextDocumentSync: TextDocumentSync{
				OpenClose: true,
				Change:    ChangeTypeFull,
			},
			DefinitionProvider: false,
			CompletionProvider: CompletionProvider{
				TriggerCharacters: []string{},
			},
			HoverProvider: true,
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
