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

	app.Commands = []cli.Command {
		cli.Command{
			Name: "dive",
			Usage: "resets repo to the oldest commit",
			Action: handlers.DiveHandler,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}