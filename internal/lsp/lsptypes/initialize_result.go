package lsptypes

// InitializeResult represents the result for the initialize method.
type InitializeResult struct {
	ServerInfo   ServerInfo         `json:"serverInfo"`
	Capabilities ServerCapabilities `json:"capabilities"`
}

// ServerCapabilities represents the capabilities of the server.
type ServerCapabilities struct {
	TextDocumentSync      TextDocumentSync      `json:"textDocumentSync"`
	DefinitionProvider    bool                  `json:"definitionProvider"`
	CompletionProvider    CompletionProvider    `json:"completionProvider"`
	HoverProvider         bool                  `json:"hoverProvider"`
	SignatureHelpProvider SignatureHelpProvider `json:"signatureHelpProvider"`
}

// TextDocumentSync represents the text document sync capabilities.
type TextDocumentSync struct {
	OpenClose bool       `json:"openClose"`
	Change    ChangeType `json:"change"`
}

// CompletionProvider represents the completion provider capabilities.
type CompletionProvider struct {
	TriggerCharacters []string `json:"triggerCharacters"`
}

// ServerInfo represents the server information.
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
