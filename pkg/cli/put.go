package cli

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreatePutCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put <url>",
		Short: "put",
		RunE: func(cmd *cobra.Command, args []string) error {
			url := ""
			if len(args) > 0 {
				url = args[0]
			}

			invocation := invoke.NewInvocation("PUT", url)
			if err := usecase.PromptReq(repos, &invocation); err != nil {
				return err
			}
			return usecase.Invoke(repos, &invocation)
		},
	}

	return cmd
}
