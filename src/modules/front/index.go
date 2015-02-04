package front

import (
	"fmt"
	spider "github.com/rainkid/spider"
	"io"
	// "math/rand"
	pserver "libs/pserver"
	ws "libs/websock"
	"os"
	// "time"
)

var page int = 614
var total int = 0
var f *os.File

type Index struct {
	FrontBase
}

func (c *Index) Index() {

}

func (c *Index) Shop() {
	url := "http://shop61757191.m.taobao.com/"
	ld := spider.NewLoader(url, "GET").WithProxy(false)
	_, err := ld.Send(nil)
	if err != nil {
		fmt.Println("...")
	}
	c.Json(0, "aa", "data")
}

func (c *Index) Ppt() {

}

func (c *Index) Ppt1() {

}


func (c *Index) Next() {
	ws.Server.SendData([]byte("next"))
	c.Json(0,"success","")
}

func (c *Index) Server() {
	code := c.GetInput("code")
	pserver.Server.SendData([]byte(code))
}

func (c *Index) Test() {
	c.DisableView = true
	params := c.GetPosts([]string{"payStatus", "queryType", "startTime", "endTime", "cookie"})
	fmt.Println(params)
	go doLoadData(params)
	c.Json(0, "request is submit, please wait.", "")
}

func doLoadData(params map[string]string) {
	var flag bool = true
	var filename string = fmt.Sprintf("%s-%s-%s-%s.csv", params["startTime"], params["endTime"], params["queryType"], params["payStatus"])
	fmt.Println(filename)
	if f == nil {
		fd, err := os.Create(filename)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("create file ", filename)
		f = fd
	}
	for {
		content, err := load(params, page)
		if err != nil {
			fmt.Println("get data error.")
		}
		if flag == false {
			fmt.Println("load data finish", filename)
			f.Close()
			break
		}

		d := dataToString(content, &flag, &total)
		fmt.Println("get data with page-", page, "total-", total)
		if d == "" {
			continue
		}
		n, err := io.WriteString(f, d)
		if err != nil {
			fmt.Println(n, err)
		}

		page++
		// time.Sleep(time.Second)
	}
}

func load(params map[string]string, page int) ([]byte, error) {
	url := fmt.Sprintf("http://pub.alimama.com/report/getTbkPaymentDetails.json?startTime=%s&endTime=%s&payStatus=%s&queryType=%s&toPage=%d&perPageSize=20", params["startTime"], params["endTime"], params["payStatus"], params["queryType"], page)
	//get content
	loader := spider.NewLoader(url, "Get")
	loader.SetHeader("Cookie", params["cookie"])
	loader.SetHeader("Refer", "http://pub.alimama.com/index.htm")
	return loader.Send(nil)
}

func dataToString(d []byte, flag *bool, total *int) string {
	hp := spider.NewHtmlParse()
	hp = hp.LoadData(d)
	o := hp.Partten(`(?U){"status".*}`).FindAllSubmatch()
	var l int = len(o)
	if l == 0 {
		fmt.Println("match data error,please input cookie ")
		return ""
	}
	hp.LoadData(o[0][0])
	if fmt.Sprintf("%s", hp.Partten(`"lastPage":(.*),"nextPage"`).FindStringSubmatch()[1]) == "true" {
		*flag = true
	} else {
		*flag = false
	}
	var output string
	if l > 0 {
		for i := 1; i < l; i++ {
			hp.LoadData(o[i][0])
			/*output += fmt.Sprintf("%s,%s,%s, %s\n",
			hp.FindJsonStr("auctionTitle")[0][1],
			hp.FindJsonInt("auctionNum")[0][1],
			hp.FindJsonStr("createTime")[0][1],
			hp.FindJsonInt("payPrice")[0][1])*/
			var createTimeStr, auctionTitleStr, exShopTitleStr, exNickNameStr, auctionNumStr, payPriceStr []byte

			createTime := hp.FindJsonStr("createTime")
			if createTime != nil && len(createTime) > 0 {
				createTimeStr = createTime[0][1]
			}
			auctionTitle := hp.FindJsonStr("auctionTitle")
			if createTime != nil && len(auctionTitle) > 0 {
				auctionTitleStr = auctionTitle[0][1]
			}
			exShopTitle := hp.FindJsonStr("exShopTitle")
			if exShopTitle != nil && len(exShopTitle) > 0 {
				exShopTitleStr = exShopTitle[0][1]
			}
			exNickName := hp.FindJsonStr("exNickName")
			if exNickName != nil && len(exNickName) > 0 {
				exNickNameStr = exNickName[0][1]
			}
			auctionNum := hp.FindJsonInt("auctionNum")
			if auctionNum != nil && len(auctionNum) > 0 {
				auctionNumStr = auctionNum[0][1]
			}
			payPrice := hp.FindJsonInt("payPrice")
			if payPrice != nil && len(payPrice) > 0 {
				payPriceStr = payPrice[0][1]
			}
			output += fmt.Sprintf("%s,%s,%s, %s,%s,%s\n",
				createTimeStr,
				auctionTitleStr,
				exShopTitleStr,
				exNickNameStr,
				auctionNumStr,
				payPriceStr,
			)
			*total++
		}
	}
	return output
}
