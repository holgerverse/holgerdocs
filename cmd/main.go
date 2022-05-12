package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "holgerdocs",
		Usage: "Generate documentation for your Infrastructure as Code configuration.",
		Commands: []*cli.Command{{
			Name:  "generate",
			Usage: "Generate documentation.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "module-path",
					Aliases:  []string{"path"},
					Usage:    "Relative path to the Terraform module you want to perform actions on.",
					Required: true,
				},
				&cli.BoolFlag{
					Name:     "terraform",
					Aliases:  []string{"tf"},
					Usage:    "Specify if you want to generate documentation for Terraform files.",
					Required: false,
				},
			},
			Action: func(c *cli.Context) error {
				holgerdocs(absolutePath(c.String("module-path")), c.Bool("terraform"))
				return nil
			},
		}}}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
