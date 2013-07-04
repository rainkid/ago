package debug

import (
	"github.com/rainkid/dogo"
	"net/http/pprof"
)

type Pprof struct {
	dogo.Controller
}

func (c *Pprof) Index() {
	pprof.Index(c.GetResponse(), c.GetRequest())
}

func (c *Pprof) Cmdline() {
	pprof.Cmdline(c.GetResponse(), c.GetRequest())
}

func (c *Pprof) Profile() {
	pprof.Profile(c.GetResponse(), c.GetRequest())
}

func (c *Pprof) Symbol() {
	pprof.Symbol(c.GetResponse(), c.GetRequest())
}
