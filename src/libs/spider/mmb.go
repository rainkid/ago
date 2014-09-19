package spider

import (
	"errors"
	"fmt"
	"strconv"
)

type MMB struct {
	item    *Item
	content string
}

func (ti *MMB) Item() {
	url := fmt.Sprintf("http://mmb.cn/wap/shop/product.do?id=%s", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Send(nil)
	ti.item.err = err
	if ti.CheckError() {
		return
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(fmt.Sprintf("%s", content)).Replace()
	ti.content = hp.content
	// ti.content = fmt.Sprintf("%s", content)
	//get title and check
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

func (ti *MMB) GetItemTitle() *MMB {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.Partten(`(?U)<div class="class169">([[:^ascii:]]+)<img`).FindStringSubmatch()
	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}
	ti.item.data["title"] = fmt.Sprintf("%s", title[1])
	return ti
}

func (ti *MMB) GetItemPrice() *MMB {
	hp := NewHtmlParse().LoadData(ti.content)
	price := hp.Partten(`(?U)<span style="color:#F6310A;">(.*)</span>`).FindStringSubmatch()

	if price == nil {
		ti.item.err = errors.New(`get price error`)
		return ti
	}
	iprice, _ := strconv.ParseFloat(price[1], 64)
	ti.item.data["price"] = fmt.Sprintf("%.2f", iprice)
	return ti
}

func (ti *MMB) GetItemImg() *MMB {
	hp := NewHtmlParse().LoadData(ti.content)
	img := hp.Partten(`(?U)"(http://rep.mmb.cn/wap/upload/productImage/+.*)"`).FindStringSubmatch()
	if img == nil {
		ti.item.err = errors.New(`get img error`)
		return ti
	}
	ti.item.data["img"] = img[1]
	return ti
}

func (ti *MMB) CheckError() bool {
	if ti.item.err != nil {
		Server.qerror <- ti.item
		return true
	}
	return false
}
