package cmd

import (
	"fmt"
	"log/slog"

	"github.com/Dobefu/DLiteScript/internal/lsp"
	"github.com/spf13/cobra"
)

var lspCmd = &cobra.Command{ //nolint:exhaustruct
	Use:   "lsp",
	Short: "Start the DLiteScript Language Server",
	Run:   runLSPCmd,
}

func init() {
	lspCmd.Flags().Bool("stdio", false, "Use stdio transport")
	lspCmd.Flags().Bool("debug", false, "Enable debug mode")

	rootCmd.AddCommand(lspCmd)
}

func runLSPCmd(cmd *cobra.Command, _ []string) {
	isDebugMode, err := cmd.Flags().GetBool("debug")

	if err != nil {
		slog.Error(fmt.Sprintf("could not parse flag: %s", err.Error()))
		setExitCode(1)

		return
	}

	handler := lsp.NewHandler(isDebugMode)
	server := lsp.NewServer(handler)

	slog.Info("Starting DLiteScript LSP server...")

	code, err := server.Start()
	setExitCode(code)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to start LSP server: %s", err.Error()))
		setExitCode(1)
	}
}
