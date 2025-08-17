package lsptypes

// Hover represents the result of a hover request.
type Hover struct {
	Contents string `json:"contents"`
	Range    *Range `json:"range,omitempty"`
}

// Range represents a range in a text document.
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
