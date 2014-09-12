package api

import (
	// "fmt"
	spider "libs/spider"
)

type Test struct {
	ApiBase
}

func (c *Test) Index() {
	params := c.GetInputs([]string{"tag", "id"})
	if params["tag"] == "" {
		c.Json(-1, "with empty tag", "")
		return
	}
	if params["id"] == "" {
		c.Json(-1, "with empty id", "")
		return
	}
	sp := spider.Start()
	// sp.Add("tmall", "21827332489")
	// sp.Add("taobao", "41040031908")
	// sp.Add("mmb", "212127")
	// sp.Add("shop", "mbaobao")
	// sp.Add("shop", "xiaomi")
	sp.Add(params["tag"], params["id"])
	c.Json(0, "success", "success")
}
