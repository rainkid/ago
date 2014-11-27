package api

import (
	spider "libs/spider"
	// "net/url"
	// "fmt"
)

type Tmall struct {
	ApiBase
}

func (c *Tmall) Test() {
	/*ret := &spider.PingResult{}
	err := spider.Ping(ret, "111.161.126.101")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ret.Average)*/
	// proxy := spider.NewSpiderProxy()
	// proxy.Load()
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
