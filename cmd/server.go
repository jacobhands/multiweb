package cmd

import "github.com/codegangsta/cli"

// CmdServer will serve the website
var CmdServer = cli.Command{
	Name:        "server",
	ShortName:   "s",
	Description: "Serves up websites",
	Action:      runServer,
	Flags: []cli.Flag{
		stringFlag("folder", "./sites", "Folder containing websites. Eg. foo.com, bar.com"),
	},
}

func runServer(c *cli.Context) {
	println("Serving files!", c.String("folder"))
}
