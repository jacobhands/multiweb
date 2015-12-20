package cmd

import (
	"net/http"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/go-macaron/switcher"
	"gopkg.in/macaron.v1"
)

const (
	// Flag names
	flagFolder = "folder"
	flagSSLPub = "ssl.pub"
	flagSSLKey = "ssl.key"
	flagDomain = "domain"
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
	cli.StringFlag{
		Name:  flagDomain,
		Value: "sites.example.com",
		Usage: "The base domain to route with.",
	},
}

func runCmdServer(ctx *cli.Context) {
	println("Serving files!", ctx.String(flagFolder), ctx.IsSet(flagFolder))
	baseURL := ctx.String(flagDomain)

	m := macaron.Classic()
	hs := switcher.NewHostSwitcher()
	// Set instance corresponding to host address.
	hs.Set("*."+baseURL, m)

	m.Get("/",
		func(resp http.ResponseWriter, req *http.Request) {
			// resp and req are injected by Macaron
			// resp.WriteHeader(200) // HTTP 200
			resp.Write([]byte(req.Host + "\n"))
			subDomain := strings.Replace(req.Host, "."+baseURL, "", 1)
			subDomain = strings.Split(subDomain, ":")[0]
			resp.Write([]byte("Subdomain: " + subDomain))
		},
	)
	hs.Run()
}
