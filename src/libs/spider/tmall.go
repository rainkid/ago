package spider

import (
	"fmt"
	"strconv"
	"strings"
)

type TmallItem struct {
	item    *Item
	content string
}

func (ti *TmallItem) Start() {
	url := fmt.Sprintf("http://detail.m.tmall.com/item.htm?id=%d", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Get()
	if err != nil {
		ti.item.err = err.Error()
		Server.qerror <- ti.item
		return
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(fmt.Sprintf("%s", content)).Convert("gbk", "utf-8").Replace()
	ti.content = hp.content

	ti.GetTitle().GetPrice().GetImg()

	Server.qfinish <- ti.item
}

func (ti *TmallItem) GetTitle() *TmallItem {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.FindJson("title")
	if title == nil {
		return ti.SError(`get title error.`)
	}
	ti.item.title = title[0][1]
	return ti
}

func (ti *TmallItem) GetPrice() *TmallItem {
	hp := NewHtmlParse().LoadData(ti.content)

	defaultPriceArr := hp.FindByAttr("b", "class", "ui-yen")
	defaultPriceStr := strings.Replace(defaultPriceArr[0][2], "&yen;", "", -1)

	var price float64
	if strings.Contains(defaultPriceStr, "-") {
		defaultPrices := strings.Split(defaultPriceStr, " - ")
		price, _ = strconv.ParseFloat(defaultPrices[0], 64)
	} else {
		price, _ = strconv.ParseFloat(defaultPriceStr, 64)
	}

	jsonData := hp.Partten(`{"isSuccess":true.*}}}`).FindStringSubmatch()
	if jsonData == nil {
		return ti.SError(`get prices jsondata error.`)
	}
	hp.LoadData(jsonData[0])
	prices := hp.FindJson("price")
	lp := len(prices)
	if prices == nil {
		return ti.SError(`get prices error.`)
	}
	for i := 0; i < lp; i++ {
		p, _ := strconv.ParseFloat(prices[i][1], 64)
		if p > 0 {
			if p < price {
				price = p
			}
		}
	}
	ti.item.price = price
	return ti
}

func (ti *TmallItem) GetImg() *TmallItem {
	hp := NewHtmlParse().LoadData(ti.content)
	d := hp.FindByAttr("section", "id", "s-showcase")
	if d == nil {
		return ti.SError(`get imgs error.`)
	}
	p := hp.LoadData(d[0][2]).Partten(`(?U)src="(.*)"`).FindStringSubmatch()
	if p == nil {
		return ti.SError(`get imgs error.`)
	}
	ti.item.img = p[1]
	return ti
}

func (ti *TmallItem) SError(msg string) *TmallItem{
	ti.item.err = msg
	Server.qerror<-ti.item
	return ti
}
