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
