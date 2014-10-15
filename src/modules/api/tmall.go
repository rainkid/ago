package api

import (
	"fmt"
	spider "libs/spider"
	// "net/url"
)

type Tmall struct {
	ApiBase
}

func (c *Tmall) Test() {
	d := c.GetPost("data")
	fmt.Println("get post data:=", d)
}

func (c *Tmall) Item() {
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
	sp.Add("TmallItem", id, callback)
	c.Json(0, "success", "success")
}

func (c *Tmall) Shop() {
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

	sp.Add("TmallShop", id, callback)
	c.Json(0, "success", "success")
}
