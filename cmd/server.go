package cmd

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/go-macaron/switcher"
	"github.com/jacobhands/multiweb/cmd/flag"
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
	cli.IntFlag{
		Name:  flag.Port,
		Value: 8080,
		Usage: "The port to listen to http requests on.",
	},
}

func runCmdServer(ctx *cli.Context) {
	println("Serving files!", ctx.String(flag.Folder), ctx.IsSet(flag.Folder))

	baseURL := ctx.String(flag.BaseURL)
	folder := ctx.String(flag.Folder)
	m := macaron.Classic()
	hs := switcher.NewHostSwitcher()
	// Set instance corresponding to host address.
	hs.Set("*."+baseURL, m)
	// http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc")))
	m.Get("/*",
		func(w http.ResponseWriter, r *http.Request) {
			subDomain := getSubDomain(r.Host, ctx.String(flag.BaseURL))
			println(r.RequestURI)
			dir := folder + "/" + subDomain + "/www/"
			http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
		},
	)
	hs.RunOnAddr(":" + strconv.Itoa(ctx.Int(flag.Port)))
}
func getSubDomain(domain, baseURL string) string {
	subDomain := strings.Replace(domain, "."+baseURL, "", 1)
	subDomain = strings.Split(subDomain, ":")[0]
	return subDomain
}
