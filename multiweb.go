package main

import (
	"net/http"
	"os"
	"strings"

	"gstx.co/jh/multiweb/cmd"

	"github.com/codegangsta/cli"
	"github.com/go-macaron/switcher"
	"gopkg.in/macaron.v1"
)

var baseURL = "sites.example.com"

func main() {
	app := cli.NewApp()
	app.Name = "multiweb"
	app.Usage = "Serve multiple websites easily using DNS magic!"
	app.Version = "0.0.1-alpha"
	app.Action = func(c *cli.Context) {
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
				resp.Write([]byte("Subdomain: " + subDomain))
			},
		)
		hs.Run()
	}
	app.Commands = []cli.Command{
		cmd.CmdServer,
	}

	app.Run(os.Args)
}
