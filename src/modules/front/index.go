package front

import (
	"fmt"
	"io"
	spider "libs/spider"
	// "math/rand"
	"os"
	// "time"
)

var page int = 1
var total int = 0
var f *os.File

type Index struct {
	FrontBase
}

func (c *Index) Index() {

}

func (c *Index) Shop() {
	url := "http://shop61757191.m.taobao.com/"
	ld := spider.NewLoader(url, "GET")
	_, err := ld.Send(nil)
	if err != nil {
		fmt.Println("...")
	}
	fmt.Println(ld.GetHeader())
	c.Json(0, "aa", "data")
}

func (c *Index) Test() {
	c.DisableView = true
	params := c.GetPosts([]string{"queryType", "startTime", "endTime", "cookie"})
	go doLoadData(params)
	c.Json(0, "request is submit, please wait.", "")
}

func doLoadData(params map[string]string) {
	var flag bool = true
	var filename string = fmt.Sprintf("%s-%s-%s.csv", params["startTime"], params["endTime"], params["queryType"])
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
	url := fmt.Sprintf("http://pub.alimama.com/report/getTbkPaymentDetails.json?startTime=%s&endTime=%s&queryType=%s&toPage=%d&perPageSize=20", params["startTime"], params["endTime"], params["queryType"], page)
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
	if fmt.Sprintf("%s", hp.FindJsonInt("lastPage")[0][1]) == "true" {
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
			output += fmt.Sprintf("%s,%s,%s, %s\n",
				hp.FindJsonStr("auctionTitle")[0][1],
				hp.FindJsonInt("auctionNum")[0][1],
				hp.FindJsonStr("createTime")[0][1],
				hp.FindJsonInt("payPrice")[0][1])
			*total++
		}
	}
	return output
}
