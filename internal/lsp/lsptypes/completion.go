package lsptypes

// CompletionParams represents the parameters for a completion request.
type CompletionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

// CompletionItem represents a completion item.
type CompletionItem struct {
	Label string `json:"label"`
}
