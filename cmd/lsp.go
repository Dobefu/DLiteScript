package cmd

import (
	"log/slog"
	"os"

	"github.com/Dobefu/DLiteScript/internal/lsp"
	"github.com/spf13/cobra"
)

var lspCmd = &cobra.Command{ //nolint:exhaustruct
	Use:   "lsp",
	Short: "Start the DLiteScript Language Server",
	Run:   runLSPCmd,
}

func init() {
	lspCmd.Flags().Bool("stdio", false, "Use stdio transport (required for LSP)")

	rootCmd.AddCommand(lspCmd)
}

func runLSPCmd(_ *cobra.Command, _ []string) {
	handler := lsp.NewHandler()
	server := lsp.NewServer(handler)

	slog.Info("Starting DLiteScript LSP server...")
	err := server.Start()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
