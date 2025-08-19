package lsptypes

// Range represents a text range in a text document.
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
