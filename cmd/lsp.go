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
	RunE:  runLSPCmd,
}

func init() {
	lspCmd.Flags().Bool("stdio", false, "Use stdio transport")
	lspCmd.Flags().Bool("debug", false, "Enable debug mode")

	rootCmd.AddCommand(lspCmd)
}

func runLSPCmd(cmd *cobra.Command, _ []string) error {
	isDebugMode, _ := cmd.Flags().GetBool("debug")

	handler := lsp.NewHandler(isDebugMode)
	server := lsp.NewServer(handler)

	slog.Info("Starting DLiteScript LSP server...")
	err := server.Start()

	if err != nil {
		slog.Error(err.Error())

		return fmt.Errorf("failed to start LSP server: %s", err.Error())
	}

	return nil
}
