package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/jacobhands/multiweb/cmd/flag"
)

// Router contains all http methods
type Router struct {
	Ctx *cli.Context
}

// New creates a new router containing cli context.
func New(ctx *cli.Context) Router {
	return Router{Ctx: ctx}
}

// GET is the router for /
func (r Router) GET(res http.ResponseWriter, req *http.Request) {
	baseURL := r.Ctx.String(flag.BaseURL)
	// resp and req are injected by Macaron
	// resp.WriteHeader(200) // HTTP 200
	res.Write([]byte("URL: " + fmt.Sprintf("%v", req.Host) + "\n"))
	res.Write([]byte("Subdomain: " + getSubDomain(req.Host, baseURL) + "\n"))
	res.Write([]byte(req.RequestURI + "\n\n"))
}

func getSubDomain(domain, baseURL string) string {
	subDomain := strings.Replace(domain, "."+baseURL, "", 1)
	subDomain = strings.Split(subDomain, ":")[0]
	return subDomain
}
