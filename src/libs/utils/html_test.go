package utils

import (
	"fmt"
	"testing"
	"strings"
	"strconv"
)

func Test_LoadUrl(t *testing.T) {
	hp := NewHtmlParse()
	err := hp.LoadUrl("http://detail.m.tmall.com/item.htm?id=36721128966")
	if err != nil {
		fmt.Println(err)
	}
	// hp.Clear()
	// fmt.Println(hp.content)
	// title := hp.FindByTagName("title")
	// fmt.Println(title[0][2])

	// price := hp.FindByAttr("b", "class","ui-yen")
	// fmt.Println(price[0][2])
}

func Test_FindJson(t *testing.T) {
	url := "http://detail.m.tmall.com/item.htm?id=21827332489"
	fmt.Println(url)
	hp := NewHtmlParse()
	err := hp.LoadUrl(url)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(hp.content)
	title := hp.FindJson("title")
	fmt.Println(title[0][1])
	defaultPriceArr :=hp.FindByAttr("b", "class", "ui-yen")
	defaultPriceStr := strings.Replace(defaultPriceArr[0][2], "&yen;", "", -1)

	var price float64
	if strings.Contains(defaultPriceStr, "-") {
		defaultPrices := strings.Split(defaultPriceStr, " - ")
		price,_ = strconv.ParseFloat(defaultPrices[0], 64)
	} else {
		price,_ = strconv.ParseFloat(defaultPriceStr, 64)
	}

	jsonData:= hp.Search(`{"isSuccess":true.*}}}`)
	hp.LoadStr(jsonData[0][0])
	prices := hp.FindJson("price")
	fmt.Println(prices)

	lp := len(prices)
	for i:=0;  i<lp;i++  {
		p, _ := strconv.ParseFloat(prices[i][1], 64)
		if p>0 {
			if p<price {
				price = p
			}
		}
	}
	fmt.Println(fmt.Sprintf("price: %.2f\n",price))
	// fmt.Println(hp.content)
}

func Test_FindByAttr(t *testing.T) {
	/*hp := NewHtmlParse()
	html := "<div><tr class='a'><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr></div>"
	hp.LoadStr(html)
	m := hp.FindByAttr("div", "class", "main")
	fmt.Println(len(m))
}

func Test_FindByAttr1(t *testing.T) {
	/*hp := NewHtmlParse()
	html := "<tr id='a' class=aaa><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>"
	hp.LoadStr(html)
	m := hp.FindByAttr("tr", "id", "a")
	fmt.Println(m)*/
}

func Test_FindByAttr2(t *testing.T) {
	/*hp := NewHtmlParse()
	html := `<tr><td class="a" a="rr">aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>`
	hp.LoadStr(html)
	m := hp.FindByAttr("td", "class", "a")
	fmt.Println(m)*/
}
