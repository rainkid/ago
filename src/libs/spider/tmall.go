package spider

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Tmall struct {
	item    *Item
	content []byte
}

func (ti *Tmall) Item() {
	url := fmt.Sprintf("http://detail.m.tmall.com/item.htm?id=%s", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Send(nil)
	ti.item.err = err
	ti.CheckError()

	hp := NewHtmlParse()
	hp = hp.LoadData(content).Convert().Replace()
	ti.content = hp.content

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
	fmt.Println(ti.item.data)
	Server.qfinish <- ti.item
}

func (ti *Tmall) GetItemTitle() *Tmall {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.FindJsonStr("title")

	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}
	ti.item.data["title"] = fmt.Sprintf("%s", title[0][1])
	return ti
}

func (ti *Tmall) GetItemPrice() *Tmall {
	hp := NewHtmlParse().LoadData(ti.content)

	defaultPriceArr := hp.FindByAttr("b", "class", "ui-yen")
	defaultPriceStr := bytes.Replace(defaultPriceArr[0][2], []byte("&yen;"), []byte(""), -1)

	var price float64
	if bytes.Contains(defaultPriceStr, []byte("-")) {
		defaultPrices := bytes.Split(defaultPriceStr, []byte(" - "))
		price, _ = strconv.ParseFloat(fmt.Sprintf("%s", defaultPrices[0]), 64)
	} else {
		price, _ = strconv.ParseFloat(fmt.Sprintf("%s", defaultPriceStr), 64)
	}

	jsonData := hp.Partten(`{"isSuccess":true.*"serviceDO"`).FindStringSubmatch()

	if jsonData != nil {
		hp.LoadData(jsonData[0])
		prices := hp.FindJsonStr("price")

		lp := len(prices)
		if prices != nil {
			for i := 0; i < lp; i++ {
				p, _ := strconv.ParseFloat(fmt.Sprintf("%s", prices[i][1]), 64)
				if p > 0 {
					if p < price {
						price = p
					}
				}
			}
		}
	}
	ti.item.data["price"] = fmt.Sprintf("%.2f", price)
	return ti
}

func (ti *Tmall) GetItemImg() *Tmall {
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
	ti.item.data["img"] = fmt.Sprintf("%s", pdata[1])
	return ti
}

func (ti *Tmall) Shop() {
	if ti.GetShopTitle().CheckError() {
		return
	}
	url := fmt.Sprintf("http://s.taobao.com/search?q=%s&app=shopsearch", ti.item.data["title"])
	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	loader.SetHeader("User-Agent", "Mozilla/5.0 (Linux; Android 4.3; Nexus 7 Build/JSS15Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Safari/537.36")
	content, err := loader.Send(nil)
	if err != nil {
		ti.item.err = err
		Server.qerror <- ti.item
		return
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(content).CleanScript().Replace().Convert()
	ti.content = hp.content

	if ti.GetShopLogo().CheckError() {
		return
	}

	if ti.GetShopImgs().CheckError() {
		return
	}
	// fmt.Println(ti.item.data)
	Server.qfinish <- ti.item
}

func (ti *Tmall) GetShopTitle() *Tmall {
	url := fmt.Sprintf("http://shop.m.tmall.com/?shop_id=%s", ti.item.id)
	fmt.Println(url)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	shop, err := loader.Send(nil)
	if err != nil {
		ti.item.err = err
		Server.qerror <- ti.item
		return ti
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(shop)
	shopname := hp.FindByTagName("title")
	if shopname == nil {
		ti.item.err = errors.New("get shop title error")
		Server.qerror <- ti.item
		return ti

	}
	uid := hp.Partten(`G_msp_userId = "(.*)"`).FindStringSubmatch()
	ti.item.data["uid"] = fmt.Sprintf("%s", uid[1])
	title := bytes.Replace(shopname[0][2], []byte("-"), []byte(""), -1)
	title = bytes.Replace(title, []byte("天猫触屏版"), []byte(""), -1)
	title = bytes.Trim(title, " ")
	ti.item.data["title"] = fmt.Sprintf("%s", title)
	// fmt.Println(ti.item.data)
	return ti
}

func (ti *Tmall) GetShopLogo() *Tmall {
	hp := NewHtmlParse().LoadData(ti.content)
	img := hp.Partten(`(?U)<a.*data-uid="` + ti.item.data["uid"] + `".*> <img src="(.*)" .*/> </a>`).FindStringSubmatch()
	if img == nil {
		ti.item.err = errors.New(`get shop img error`)
		return ti
	}
	ti.item.data["img"] = fmt.Sprintf("%s", img[1])
	return ti
}

func (ti *Tmall) GetShopImgs() *Tmall {
	hp := NewHtmlParse().LoadData(ti.content)
	imgs := hp.Partten(`(?U)<a trace="auction".*data-uid="` + ti.item.data["uid"] + `".*> <img src="(.*)";"> </a>`).FindAllSubmatch()
	if imgs == nil {
		ti.item.err = errors.New(`get shop imgs error`)
		return ti
	}

	var imglist []string
	l := len(imgs)
	if l > 3 {
		l = 3
	}
	for i := 1; i <= l; i++ {
		imglist = append(imglist, fmt.Sprintf("%s", imgs[i][1]))
	}

	ti.item.data["imgs"] = strings.Join(imglist, ",")
	return ti
}

func (ti *Tmall) CheckError() bool {
	if ti.item.err != nil {
		Server.qerror <- ti.item
		return true
	}
	return false
}
