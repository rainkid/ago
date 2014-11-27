package api

import (
	"fmt"
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

func (c *Taobao) Samestyle() {
	id := c.GetInput("id")
	callback := c.GetInput("callback")
	title := c.GetInput("title")
	if id == "" {
		c.Json(-1, "with empty id", "")
		return
	}
	if title == "" {
		c.Json(-1, "with empty title", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}

	surl := fmt.Sprintf("http://s.taobao.com/search?q=%s", title)
	fmt.Println(surl)
	sloader := spider.NewLoader(surl, "Get").WithPcAgent()
	scontent, _ := sloader.Send(nil)

	shp := spider.NewHtmlParse().LoadData(scontent).Replace().Convert()
	sret := shp.Partten(`"nid":"` + id + `","pid":"-(\d+)"`).FindStringSubmatch()
	if sret == nil || len(sret) == 0 {
		c.Json(-1, "fail", "")
	}
	pid := sret[1]

	var result []map[string]string
	url := fmt.Sprintf("http://s.taobao.com/list?tab=all&type=samestyle&uniqpid=-%s&app=i2i&nid=%s", pid, id)
	fmt.Println(url)
	loader := spider.NewLoader(url, "Get").WithPcAgent()
	content, _ := loader.Send(nil)

	hp := spider.NewHtmlParse().LoadData(content).Replace().Convert()
	ret := hp.FindByAttr("div", "class", "row item icon-datalink")

	l := len(ret) - 1
	if l <= 0 {
		c.Json(-1, "fail", result)
	}
	for i := 1; i < l; i++ {
		data := map[string]string{"comment_num": "0", "pay_num": "0"}
		val := ret[i][1]
		hp1 := spider.NewHtmlParse().LoadData(val)

		id := hp1.Partten(`(?U)data-item="(\d+)"`).FindStringSubmatch()
		data["id"] = fmt.Sprintf("%s", id[1])

		price := hp1.Partten(`(?U)<i>￥</i>(.*)</span>`).FindStringSubmatch()
		data["price"] = fmt.Sprintf("%s", price[1])

		imgs := hp1.Partten(`(?U)data-ks-lazyload="(.*)"`).FindStringSubmatch()
		data["img"] = fmt.Sprintf("%s", imgs[1])

		title := hp1.Partten(`(?U)title="(.*)"`).FindStringSubmatch()
		data["title"] = fmt.Sprintf("%s", title[1])

		address := hp1.Partten(`(?U)<div class="seller-loc">(.*)</div>`).FindStringSubmatch()
		data["address"] = fmt.Sprintf("%s", address[1])

		pay_num := hp1.Partten(`(?U)(\d+) 人付款`).FindStringSubmatch()
		if pay_num != nil {
			data["pay_num"] = fmt.Sprintf("%s", pay_num[1])
		}

		comment_num := hp1.Partten(`(?U)(\d+) 条评论`).FindStringSubmatch()
		if comment_num != nil {
			data["comment_num"] = fmt.Sprintf("%s", comment_num[1])
		}

		score := hp1.Partten(`(?U)<span class="feature-dsr-num">(.*)</span>`).FindStringSubmatch()
		if comment_num != nil {
			data["score"] = fmt.Sprintf("%s", score[1])
		}
		if i == 6 {
			break
		}
		result = append(result, data)
	}
	c.Json(0, "success", result)
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
