package spider

import (
	"fmt"
	"strconv"
	"strings"
)

type TaobaoItem struct {
	item    *Item
	content string
}

func (ti *TaobaoItem) Start() {
	url := fmt.Sprintf("http://hws.m.taobao.com/cache/wdetail/5.0/?id=%d", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Get()
	if err != nil {
		ti.item.err = err.Error()
		Server.qerror <- ti.item
		return
	}

	ti.content = strings.Replace(fmt.Sprintf("%s", content), `\"`, `"`, -1)
	ti.GetTitle().GetPrice().GetImg()

	Server.qfinish <- ti.item
}

func (ti *TaobaoItem) GetTitle() *TaobaoItem {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.Partten(`(?U)"itemId":"\d+","title":"(.*)"`).FindStringSubmatch()

	if title == nil {
		return ti.SError(`get title error.`);
	}
	ti.item.title = title[1]
	return ti
}

func (ti *TaobaoItem) GetPrice() *TaobaoItem {
	hp := NewHtmlParse().LoadData(ti.content)
	price := hp.Partten(`(?U)"rangePrice":".*","price":"(.*)"`).FindStringSubmatch()

	if price == nil {
		return ti.SError(`get price error.`);
	}
	ti.item.price, _ = strconv.ParseFloat(price[1], 64)
	return ti
}

func (ti *TaobaoItem) GetImg() *TaobaoItem {
	hp := NewHtmlParse().LoadData(ti.content)
	img := hp.Partten(`(?U)"picsPath":\["(.*)"`).FindStringSubmatch()

	if img == nil {
		return ti.SError(`get img error.`);
	}
	ti.item.img = img[1]
	return ti
}

func (ti *TaobaoItem) SError(msg string) *TaobaoItem{
	ti.item.err = msg
	Server.qerror<-ti.item
	return ti
}
