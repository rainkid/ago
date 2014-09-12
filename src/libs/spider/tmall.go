package spider

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TmallItem struct {
	item    *Item
	content string
}

func (ti *TmallItem) Start() {
	url := fmt.Sprintf("http://detail.m.tmall.com/item.htm?id=%s", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Get()
	ti.item.err = err
	ti.CheckError()

	hp := NewHtmlParse()
	hp = hp.LoadData(fmt.Sprintf("%s", content)).Convert("gbk", "utf-8").Replace()
	ti.content = hp.content

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

func (ti *TmallItem) GetTitle() *TmallItem {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.FindJson("title")
	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}
	ti.item.data["title"] = title[0][1]
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
		ti.item.err = errors.New(`get prices jsondata error`)
		return ti
	}
	hp.LoadData(jsonData[0])
	prices := hp.FindJson("price")
	lp := len(prices)
	if prices == nil {
		ti.item.err = errors.New(`get prices error`)
		return ti
	}
	for i := 0; i < lp; i++ {
		p, _ := strconv.ParseFloat(prices[i][1], 64)
		if p > 0 {
			if p < price {
				price = p
			}
		}
	}
	ti.item.data["price"] = fmt.Sprintf("%.2f", price)
	return ti
}

func (ti *TmallItem) GetImg() *TmallItem {
	hp := NewHtmlParse().LoadData(ti.content)
	data := hp.FindByAttr("section", "id", "s-showcase")
	if data == nil {
		ti.item.err = errors.New(`get imgs error`)
		return ti
	}
	pdata := hp.LoadData(data[0][2]).Partten(`(?U)src="(.*)"`).FindStringSubmatch()
	if pdata == nil {
		ti.item.err = errors.New(`get imgs error`)
		return ti
	}
	ti.item.data["img"] = pdata[1]
	return ti
}

func (ti *TmallItem) CheckError() bool {
	if ti.item.err != nil {
		Server.qerror <- ti.item
		return true
	}
	return false
}
