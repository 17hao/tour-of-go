package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var showCommand = &cli.Command{
	Name:      "show",
	Usage:     "print options or args",
	UsageText: "my-cli show [--local[=<value>]] [<args>]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "local",
			Value: "default value",
		},
	},
	Action: func(context *cli.Context) error {
		switch context.Args().First() {
		case "options":
			fmt.Printf("options: %s=%s, %s=%s, %s=%s\n", "verbose", context.String("verbose"), "global", context.String("global"), "local", context.String("local"))
		case "args":
			fmt.Printf("args: %+v\n", context.Args().Slice())
		default:
			fmt.Printf("usage: my-cli show [--local[=<value>]] [<args>]\n")
		}
		return nil
	},
}

// my-cli show global options or sub-command options or sub-command arguments
func main() {
	app := &cli.App{
		Name:      "my-cli",
		Usage:     "show global options, sub-command options, sub-command arguments",
		UsageText: "my-cli [-v | --version] [--global=<value>] <command> [--local[=<value>]] [<args>]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"V"},
				Usage:   "verbose flag",
			},
			&cli.StringFlag{
				Name: "global",
			},
		},
		Commands: []*cli.Command{
			showCommand,
		},
		Version: "1.0.0",
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
