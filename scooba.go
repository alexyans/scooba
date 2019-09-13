package main

import (
	"os"

	"github.com/alexyans/scooba/handlers"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Scooba"
	app.Usage = "Dive into a git-tracked codebase"

	app.Commands = []cli.Command{
		{
			Name: "dive",
			Usage: "resets repo to the oldest commit",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "commit, c",
					Usage: "sets the commit hash to start from (default: oldest chronological commit)",
				},
			},
			Action: handlers.DiveHandler,
		},
		{
			Name: "forward",
			Usage: "moves ahead by as many commits as specified (Default: 1)",
			Action: handlers.ForwardHandler,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}