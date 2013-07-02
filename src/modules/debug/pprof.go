package debug

import (
	"fmt"
	"github.com/rainkid/dogo"
	utils "libs/utils"
	"net/http/pprof"
)

type Pprof struct {
	dogo.Controller
}

func (c *Pprof) Index() {
	param := c.GetInput("param")
	switch utils.ItoString(param) {
	case "":
		pprof.Index(c.GetResponse(), c.GetRequest())
	case "cmdline":
		pprof.Cmdline(c.GetResponse(), c.GetRequest())
	case "profile":
		pprof.Profile(c.GetResponse(), c.GetRequest())
	case "symbol":
		pprof.Symbol(c.GetResponse(), c.GetRequest())
	default:
		pprof.Index(c.GetResponse(), c.GetRequest())
	}
	c.GetResponse().WriteHeader(200)
}
