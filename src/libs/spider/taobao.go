package spider

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Taobao struct {
	item    *Item
	content string
}

func (ti *Taobao) Item() {
	url := fmt.Sprintf("http://hws.m.taobao.com/cache/wdetail/5.0/?id=%s", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Send(nil)
	ti.item.err = err
	ti.CheckError()

	ti.content = strings.Replace(fmt.Sprintf("%s", content), `\"`, `"`, -1)
	if ti.GetItemTitle().CheckError() {
		return
	}
	//check price
	if ti.GetItemPrice().CheckError() {
		return
	}
	if ti.GetItemImg().CheckError() {
		return
	}

	Server.qfinish <- ti.item
}

func (ti *Taobao) GetItemTitle() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.Partten(`(?U)"itemId":"\d+","title":"(.*)"`).FindStringSubmatch()

	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}
	ti.item.data["title"] = title[1]
	return ti
}

func (ti *Taobao) GetItemPrice() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	price := hp.Partten(`(?U)"rangePrice":".*","price":"(.*)"`).FindStringSubmatch()

	if price == nil {
		ti.item.err = errors.New(`get price error`)
		return ti
	}
	iprice, _ := strconv.ParseFloat(price[1], 64)
	ti.item.data["price"] = fmt.Sprintf("%.2f", iprice)
	return ti
}

func (ti *Taobao) GetItemImg() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	img := hp.Partten(`(?U)"picsPath":\["(.*)"`).FindStringSubmatch()

	if img == nil {
		ti.item.err = errors.New(`get img error`)
		return ti
	}
	ti.item.data["img"] = img[1]
	return ti
}

func (ti *Taobao) Shop() {
	url := fmt.Sprintf("http://shop%s.taobao.com/?search=y&orderType=hotsell_desc", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	loader.SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36")
	content, err := loader.Send(nil)
	if err != nil {
		ti.item.err = err
		Server.qerror <- ti.item
		return
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(fmt.Sprintf("%s", content)).Convert("gbk", "utf-8")
	ti.content = hp.content

	if ti.GetShopTitle().CheckError() {
		return
	}

	if ti.GetShopRank().CheckError() {
		return
	}

	if ti.GetShopImgs().CheckError() {
		return
	}
	Server.qfinish <- ti.item
}

func (ti *Taobao) GetShopTitle() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.Partten(`<title>店内搜索页-(.*)-淘宝网</title>`).FindStringSubmatch()
	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}
	ti.item.data["title"] = title[1]
	return ti
}

func (ti *Taobao) GetShopImgs() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	imgs := hp.Partten(`(?U)src="(.*/uploaded/.*)"`).FindAllSubmatch()
	if imgs == nil {
		ti.item.err = errors.New(`get imgs error`)
		return ti
	}
	ti.item.data["imgs"] = fmt.Sprintf("%s,%s,%s", imgs[0][1], imgs[1][1], imgs[2][1])
	return ti
}

func (ti *Taobao) GetShopRank() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	rank := hp.Partten(`(?U)http://pics.taobaocdn.com/newrank/s_(.*).gif`).FindStringSubmatch()
	if rank == nil {
		ti.item.err = errors.New(`get imgs error`)
		return ti
	}
	ti.item.data["rank"] = rank[1]
	return ti
}

func (ti *Taobao) CheckError() bool {
	if ti.item.err != nil {
		Server.qerror <- ti.item
		return true
	}
	return false
}
