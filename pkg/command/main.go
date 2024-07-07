package command

import (
	"fmt"
	"time"

	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/usecase"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
)

func New(repos repository.Repos) *cli.App {
	app := &cli.App{
		Name:    "walkhttp",
		Version: "0.0.8",
		Usage:   "A CLI tool to call http endpoint with browser or prompt.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "port",
				Usage:    "Serve port",
				Value:    3000,
			},
			&cli.StringFlag{
				Name:     "config",
				Aliases: []string{"c"},
				Usage:    "config file path",
				Value:    "walkhttp.json",
			},
		},
		Commands: []*cli.Command{
			NewGetCommand(repos),
			NewPostCommand(repos),
			NewPutCommand(repos),
			NewDeleteCommand(repos),
			NewOptionsCommand(repos),
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			configpath := c.String("config")
			conf := usecase.LoadConfig(repos, configpath)

			usecase.PrintBanner(repos)

			go func() {
				time.Sleep(1 * time.Second)
				url := fmt.Sprintf("http://localhost:%d", port)
				browser.OpenURL(url)
			}()

			return invoke.Serve(repos, conf.BaseUrl, port)

		},
		Suggest: true,
	}

	// disable default
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		return err
	}
	app.HideHelpCommand = true
	cli.AppHelpTemplate = `{{.Usage}}

USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[flags]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
FLAGS:{{range .Flags}}
	{{.}}{{end}}
{{end}}`

	return app
}
