package router

import (
	"net/http"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/jacobhands/multiweb/cmd/flag"
)

// Router contains all http methods
type Router struct {
	Ctx *cli.Context
}

// GET is the router for /
func (r Router) GET(res http.ResponseWriter, req *http.Request) {
	baseURL := r.Ctx.String(flag.BaseURL)
	// resp and req are injected by Macaron
	// resp.WriteHeader(200) // HTTP 200
	res.Write([]byte(req.Host + "\n"))
	subDomain := strings.Replace(req.Host, "."+baseURL, "", 1)
	subDomain = strings.Split(subDomain, ":")[0]
	res.Write([]byte("Subdomain: " + subDomain + "\n"))
}
