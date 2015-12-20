package main

import (
	"os"

	"gstx.co/jh/multiweb/cmd"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "multiweb"
	app.Usage = "Serve multiple websites easily using DNS magic!"
	app.Version = "0.0.1-alpha"
	app.Action = func(c *cli.Context) {

	}
	app.Commands = []cli.Command{
		cmd.CmdServer,
	}

	app.Run(os.Args)
}
