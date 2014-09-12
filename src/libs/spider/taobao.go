package spider

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TaobaoItem struct {
	item    *Item
	content string
}

func (ti *TaobaoItem) Start() {
	url := fmt.Sprintf("http://hws.m.taobao.com/cache/wdetail/5.0/?id=%s", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Get()
	ti.item.err = err
	ti.CheckError()

	ti.content = strings.Replace(fmt.Sprintf("%s", content), `\"`, `"`, -1)
	if ti.GetTitle().CheckError() {
		return
	}
	//check price
	if ti.GetPrice().CheckError() {
		return
	}
	if ti.GetImg().CheckError() {
		return
	}

	Server.qfinish <- ti.item
}

func (ti *TaobaoItem) GetTitle() *TaobaoItem {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.Partten(`(?U)"itemId":"\d+","title":"(.*)"`).FindStringSubmatch()

	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}
	ti.item.data["title"] = title[1]
	return ti
}

func (ti *TaobaoItem) GetPrice() *TaobaoItem {
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

func (ti *TaobaoItem) GetImg() *TaobaoItem {
	hp := NewHtmlParse().LoadData(ti.content)
	img := hp.Partten(`(?U)"picsPath":\["(.*)"`).FindStringSubmatch()

	if img == nil {
		ti.item.err = errors.New(`get img error`)
		return ti
	}
	ti.item.data["img"] = img[1]
	return ti
}

func (ti *TaobaoItem) CheckError() bool {
	if ti.item.err != nil {
		Server.qerror <- ti.item
		return true
	}
	return false
}
