package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
	"github.com/Dobefu/DLiteScript/internal/version"
)

func (h *Handler) handleInitialize() (json.RawMessage, *jsonrpc2.Error) {
	result := lsptypes.InitializeResult{
		ServerInfo: lsptypes.ServerInfo{
			Name:    "DLiteScript",
			Version: version.GetVersion(),
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

	data, _ := json.Marshal(result)

	return data, nil
}
