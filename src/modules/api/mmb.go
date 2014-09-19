package api

import (
	spider "libs/spider"
)

type Mmb struct {
	ApiBase
}

func (c *Mmb) Item() {
	id := c.GetInput("id")
	callback := c.GetInput("callback")
	if id == "" {
		c.Json(-1, "with empty id", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}
	sp := spider.Start()
	sp.Add("MmbItem", id, callback)
	c.Json(0, "success", "success")
}
