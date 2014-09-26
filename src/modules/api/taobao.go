package api

import (
	spider "libs/spider"
)

type Taobao struct {
	ApiBase
}

func (c *Taobao) Item() {
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

	sp.Add("TaobaoItem", id, callback)
	c.Json(0, "success", "success")
}

func (c *Taobao) Shop() {
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

	sp.Add("TaobaoShop", id, callback)
	c.Json(0, "success", "success")
}
