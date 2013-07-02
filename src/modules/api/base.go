package api

import (
	"github.com/rainkid/dogo"
)

type ApiBase struct {
	dogo.Controller
}

func (c *ApiBase) Init() {
	c.DisableView = true
}
