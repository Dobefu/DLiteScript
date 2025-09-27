package lsp

import (
	"encoding/json"
	"fmt"
	"os"
)

func (h *Handler) printDebugMessage(
	method string,
	params json.RawMessage,
) {
	fmt.Fprintf(os.Stderr, "Received request: %s\n", method)

	var formattedParams map[string]any
	err := json.Unmarshal(params, &formattedParams)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling params: %s\n", err)
	}

	formattedParamsJSON, _ := json.MarshalIndent(formattedParams, "", "  ")

	fmt.Fprintf(os.Stderr, "Params: %s\n", string(formattedParamsJSON))
	fmt.Fprintf(os.Stderr, "---\n\n")
}
