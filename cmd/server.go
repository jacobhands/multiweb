package cmd

import "github.com/codegangsta/cli"

const (
	// Flag names
	flagFolder = "folder"
	flagSSLPub = "ssl.pub"
	flagSSLKey = "ssl.key"
)

// CmdServer will serve the website
var CmdServer = cli.Command{
	Name:        "server",
	ShortName:   "s",
	Description: "Serves up websites",
	Action:      runServer,
	Flags:       cmdServerFlags,
}

// Flags
var cmdServerFlags = []cli.Flag{
	cli.StringFlag{
		Name:  flagFolder,
		Value: "./sites",
		Usage: "Folder containing websites. Eg. foo.com and bar.com",
	},
	cli.StringFlag{
		Name:  flagSSLPub,
		Value: "",
		Usage: "SSL public key for server",
	},
	cli.StringFlag{
		Name:  flagSSLKey,
		Value: "",
		Usage: "SSL private key for server",
	},
}

func runServer(c *cli.Context) {
	println("Serving files!", c.String(flagFolder), c.IsSet(flagFolder))
}
