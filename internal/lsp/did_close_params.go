package lsp

// DidCloseParams represents the parameters for a didClose request.
type DidCloseParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
