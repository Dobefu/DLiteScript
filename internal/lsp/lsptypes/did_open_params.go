package lsptypes

// DidOpenParams represents the parameters for a didOpen request.
type DidOpenParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
