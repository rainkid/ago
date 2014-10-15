package spider

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Taobao struct {
	item    *Item
	content []byte
}

func (ti *Taobao) Item() {
	url := fmt.Sprintf("http://hws.m.taobao.com/cache/wdetail/5.0/?id=%s", ti.item.id)

	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	content, err := loader.Send(nil)
	ti.item.err = err
	ti.CheckError()

	ti.content = bytes.Replace(content, []byte(`\"`), []byte(`"`), -1)
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
	// fmt.Println(ti.item.data)
	Server.qfinish <- ti.item
}

func (ti *Taobao) GetItemTitle() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.Partten(`(?U)"itemId":"\d+","title":"(.*)"`).FindStringSubmatch()

	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}
	ti.item.data["title"] = fmt.Sprintf("%s", title[1])
	return ti
}

func (ti *Taobao) GetItemPrice() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)
	price := hp.Partten(`(?U)"rangePrice":".*","price":"(.*)"`).FindStringSubmatch()

	if price == nil {
		price = hp.Partten(`(?U)"price":"(.*)"`).FindStringSubmatch()
	}
	if price == nil {
		ti.item.err = errors.New(`get price error`)
		return ti
	}

	var iprice float64
	if bytes.Index(price[1], []byte("-")) > 0 {
		price = bytes.Split(price[1], []byte("-"))
		iprice, _ = strconv.ParseFloat(fmt.Sprintf("%s", price[0]), 64)
	} else {
		iprice, _ = strconv.ParseFloat(fmt.Sprintf("%s", price[1]), 64)
	}

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
	ti.item.data["img"] = fmt.Sprintf("%s", img[1])
	return ti
}

func (ti *Taobao) Shop() {
	if ti.GetShopTitle().CheckError() {
		return
	}
	url := fmt.Sprintf("http://s.taobao.com/search?q=%s&app=shopsearch", ti.item.data["title"])
	fmt.Println(url)
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
	fmt.Println(ti.item.data)
	Server.qfinish <- ti.item
}

func (ti *Taobao) GetShopTitle() *Taobao {
	url := fmt.Sprintf("http://shop%s.taobao.com", ti.item.id)
	fmt.Println(url)
	ti.item.url = url
	//get content
	loader := NewLoader(url, "Get")
	loader.SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36")
	shop, err := loader.Send(nil)
	// fmt.Println(fmt.Sprintf("%s", shop))
	if err != nil {
		ti.item.err = err
		Server.qerror <- ti.item
		return ti
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(shop).Convert().Replace()
	shopname := hp.FindByTagName("title")
	uid := hp.FindJsonStr("userId")

	if shopname == nil {
		ti.item.err = errors.New("get shop title error")
		Server.qerror <- ti.item
		return ti

	}
	ti.item.data["uid"] = fmt.Sprintf("%s", uid[0][1])
	title := bytes.Replace(shopname[0][2], []byte("首页"), []byte(""), -1)
	title = bytes.Replace(title, []byte("淘宝网"), []byte(""), -1)
	title = bytes.Replace(title, []byte("-"), []byte(" "), -1)
	title = bytes.Trim(title, " ")
	ti.item.data["title"] = fmt.Sprintf("%s", title)
	return ti
}

func (ti *Taobao) GetShopImgs() *Taobao {
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

func (ti *Taobao) GetShopLogo() *Taobao {
	hp := NewHtmlParse().LoadData(ti.content)

	img := hp.Partten(`(?U)<a.*data-uid="` + ti.item.data["uid"] + `".*> <img src="(.*)" .*/> </a>`).FindStringSubmatch()
	if img == nil {
		ti.item.err = errors.New(`get shop img error`)
		return ti
	}
	ti.item.data["img"] = fmt.Sprintf("%s", img[1])
	return ti
}

func (ti *Taobao) CheckError() bool {
	if ti.item.err != nil {
		Server.qerror <- ti.item
		return true
	}
	return false
}
