package holgerdocs

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "holgersync",
		Usage: "Manage IaC from the CLI.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "debug",
				Usage: "Enable debug mode",
			},
		},
		Commands: []*cli.Command{{
			Name:     "holgerdocs",
			Category: "cli-commands",
			Subcommands: []*cli.Command{{
				Name:      "terraform",
				HelpName:  "terraform",
				Category:  "holgerdocs",
				Usage:     "holgersync hoglerdocs terraform --module-path=<path>",
				UsageText: "Generate documentation for Terraform modules.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "module-path",
						Usage:    "Relative path to the Terraform module you want to perform actions on.",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					holgerdocs(absolutePath(c.String("module-path")), "terraform")
					return nil
				},
			},
			},
		}}}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
