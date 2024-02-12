package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateDeleteCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete request",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}