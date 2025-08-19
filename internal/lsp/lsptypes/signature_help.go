package lsptypes

// SignatureHelpParams represents the parameters for a signature help request.
type SignatureHelpParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

// SignatureHelp represents the signature help response.
type SignatureHelp struct {
	Signatures      []SignatureInformation `json:"signatures"`
	ActiveSignature int                    `json:"activeSignature"`
	ActiveParameter int                    `json:"activeParameter"`
}

// SignatureInformation represents information about a function signature.
type SignatureInformation struct {
	Label         string                 `json:"label"`
	Documentation *MarkupContent         `json:"documentation,omitempty"`
	Parameters    []ParameterInformation `json:"parameters,omitempty"`
}

// ParameterInformation represents information about a function parameter.
type ParameterInformation struct {
	Label         string         `json:"label"`
	Documentation *MarkupContent `json:"documentation,omitempty"`
}

// MarkupContent represents documentation content.
type MarkupContent struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}
