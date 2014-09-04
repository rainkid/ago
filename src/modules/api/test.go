package api

import (
	spider "libs/spider"
)

type Test struct {
	ApiBase
}

func (c *Test) Index() {
	sp := spider.Start()
	sp.Add("tmall", 21827332489)
	c.Json(0, "aaa", "Content")
}
