package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateOptionsCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "options",
		Short: "make a options request",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")

			invocation, err := usecase.PromptReq(repos, "OPTIONS", url)
			if err != nil {
				return err
			}

			if err := usecase.Invoke(repos, &invocation); err != nil {
				return err
			}
			fmt.Printf("status: %d\n", invocation.Status)

			return nil
		},
	}
	cmd.Flags().String("url", "", "url")

	return cmd
}