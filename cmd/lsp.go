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
	lspCmd.Flags().Bool("debug", false, "Enable debug mode")

	rootCmd.AddCommand(lspCmd)
}

func runLSPCmd(cmd *cobra.Command, _ []string) {
	isDebugMode, _ := cmd.Flags().GetBool("debug")

	handler := lsp.NewHandler(isDebugMode)
	server := lsp.NewServer(handler)

	slog.Info("Starting DLiteScript LSP server...")
	err := server.Start()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
