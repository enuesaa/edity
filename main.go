package main

import (
	"github.com/spf13/cobra"
	"github.com/enuesaa/walkin/internal/handler"
)

func main() {
	cmd := &cobra.Command{
		Use:     "walkin",
		Short:   "Instant web server.",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			handler.Handle()
		},
	}

	// disable default
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.Execute()
}
