package cmd

import (
	"context"
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
	rootCmd.AddCommand(lspCmd)
}

func runLSPCmd(_ *cobra.Command, _ []string) {
	ctx := context.Background()

	server := lsp.NewServer()

	slog.Info("DLiteScript LSP server starting...")

	err := server.Start(ctx)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
