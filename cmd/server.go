package cmd

import (
	"io/ioutil"
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
	fss := fileServerSet{}
	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {
		if f.IsDir() {
			domain := f.Name()
			folderRoot := folder + "/" + f.Name()
			println("NEW FOLDER ROOT: " + folderRoot)
			newFss := fileServer{
				Domain:     domain,
				FolderRoot: folderRoot,
				ServeHTTP: func(w http.ResponseWriter, r *http.Request) {
					println(r.RequestURI)
					dir := folderRoot + "/www/"
					http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
				},
			}
			fss.FileServers = append(fss.FileServers, newFss)
		}
	}
	for i, v := range fss.FileServers {
		fss.Index[v.Domain] = i
	}
	m.Get("/*", fss.ServeHTTP)
	hs.RunOnAddr(":" + strconv.Itoa(ctx.Int(flag.Port)))
}

type fileServer struct {
	Domain     string
	FolderRoot string
	ServeHTTP  func(http.ResponseWriter, *http.Request)
}

type fileServerSet struct {
	FileServers []fileServer
	Index       map[string]int
}

func (f fileServerSet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domain := strings.Split(r.Host, ":")[0]
	for _, f := range f.FileServers {
		// println(f.Domain + " | " + r.Host)
		if f.Domain == domain {
			f.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func getSubDomain(domain, baseURL string) string {
	subDomain := strings.Replace(domain, "."+baseURL, "", 1)
	subDomain = strings.Split(subDomain, ":")[0]
	return subDomain
}
