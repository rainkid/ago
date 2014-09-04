package spider

import (
	"fmt"
	"strconv"
)

type MMBItem struct {
	item    *Item
	content string
}

func (ti *MMBItem) Start() {
	url := fmt.Sprintf("http://mmb.cn/wap/shop/product.do?id=%d", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Get()
	if err != nil {
		ti.item.err = err.Error()
		Server.qerror <- ti.item
		return
	}
	ti.content = fmt.Sprintf("%s", content)

	ti.GetTitle().GetPrice().GetImg()

	Server.qfinish <- ti.item
}

func (ti *MMBItem) GetTitle() *MMBItem {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.Partten(`<div class="class169">([[:^ascii:]]+)<br/>`).FindStringSubmatch()
	
	if title == nil {
		return ti.SError(`get title error.`)
	}
	return ti
}

func (ti *MMBItem) GetPrice() *MMBItem {
	hp := NewHtmlParse().LoadData(ti.content)
	price := hp.Partten(`(?U)<div class="class131">.*,买卖宝价:(.*)元<br/>`).FindStringSubmatch()

	if price == nil {
		return ti.SError(`get price error.`)
	}
	ti.item.price, _ = strconv.ParseFloat(price[1], 64)
	return ti
}

func (ti *MMBItem) GetImg() *MMBItem {
	hp := NewHtmlParse().LoadData(ti.content)
	img := hp.Partten(`(?U)<div class="class169">[[:^ascii:]]+<br/><img src="(.*)"`).FindStringSubmatch()
	if img == nil {
		return ti.SError(`get img error.`)
	}
	ti.item.img = img[1]
	return ti
}

func (ti *MMBItem) SError(msg string) *MMBItem{
	ti.item.err = msg
	Server.qerror<-ti.item
	return ti
}