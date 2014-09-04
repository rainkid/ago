package spider

import (
	// "fmt"
	"testing"
	// "strings"
	// "strconv"
)

func Test_Tmall(t *testing.T) {
	/*url := "http://detail.m.tmall.com/item.htm?id=21827332489"
	loader := NewLoader(url, "Get")
	content, err := loader.Get()
	if err != nil {
		fmt.Println(err)
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(fmt.Sprintf("%s", content)).Convert("gbk","utf-8").Replace()


	//get title
	title := hp.FindJson("title")
	if title != nil {
		fmt.Println("title:"+title[0][1])
	}
	defaultPriceArr :=hp.FindByAttr("b", "class", "ui-yen")
	defaultPriceStr := strings.Replace(defaultPriceArr[0][2], "&yen;", "", -1)

	var price float64
	if strings.Contains(defaultPriceStr, "-") {
		defaultPrices := strings.Split(defaultPriceStr, " - ")
		price,_ = strconv.ParseFloat(defaultPrices[0], 64)
	} else {
		price,_ = strconv.ParseFloat(defaultPriceStr, 64)
	}

	jsonData:= hp.Partten(`{"isSuccess":true.*}}}`).FindStringSubmatch()
	hp.LoadData(jsonData[0])
	prices := hp.FindJson("price")

	lp := len(prices)
	for i:=0;  i<lp;i++  {
		p, _ := strconv.ParseFloat(prices[i][1], 64)
		if p>0 {
			if p<price {
				price = p
			}
		}
	}
	fmt.Println(fmt.Sprintf("price: %.2f\n",price))*/
}

func Test_Spider(t *testing.T) {
	sp = Serv

}
