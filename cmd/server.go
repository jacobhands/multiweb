package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/go-macaron/switcher"
	"github.com/jacobhands/multiweb/cmd/flag"
	"github.com/jacobhands/multiweb/router"
	"gopkg.in/macaron.v1"
)

// CmdServer will serve the website
var CmdServer = cli.Command{
	Name:        "server",
	ShortName:   "s",
	Description: "Serves up websites",
	Action:      runCmdServer,
	Flags:       cmdServerFlags,
}

// Flags
var cmdServerFlags = []cli.Flag{
	cli.StringFlag{
		Name:  flag.Folder,
		Value: "./sites",
		Usage: "Folder containing websites. Eg. foo.com and bar.com",
	},
	cli.StringFlag{
		Name:  flag.BaseURL,
		Value: "sites.example.com",
		Usage: "The base domain to route with.",
	},
}

func runCmdServer(ctx *cli.Context) {
	println("Serving files!", ctx.String(flag.Folder), ctx.IsSet(flag.Folder))

	baseURL := ctx.String(flag.BaseURL)
	r := router.New(ctx)
	m := macaron.Classic()
	hs := switcher.NewHostSwitcher()

	// Set instance corresponding to host address.
	hs.Set("*."+baseURL, m)

	m.Get("/*", r.GET)
	hs.Run()
}
