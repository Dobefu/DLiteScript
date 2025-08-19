package lsptypes

// SignatureHelpProvider represents the signature help provider capabilities.
type SignatureHelpProvider struct {
	TriggerCharacters []string `json:"triggerCharacters"`
}
