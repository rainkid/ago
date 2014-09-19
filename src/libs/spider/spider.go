package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

var (
	Server *Spider
	loger  *log.Logger = log.New(os.Stdout, "[SPIDER] ", log.Ldate|log.Ltime)
	smsUrl string      = "http://gou.3gtest.gionee.com/api/wifi/sms"
)

type Spider struct {
	qstart  chan *Item
	qfinish chan *Item
	qerror  chan *Item
}

type Item struct {
	id       string
	url      string
	callback string
	data     map[string]string
	tag      string
	err      error
}

func NewSpider() *Spider {
	Server = &Spider{
		qstart:  make(chan *Item),
		qfinish: make(chan *Item),
		qerror:  make(chan *Item),
	}
	return Server
}

func Start() *Spider {
	if Server == nil {
		Server = NewSpider()
		go Server.Listen()
	}
	return Server
}

func (spider *Spider) Do(item *Item) {
	switch item.tag {
	case "TmallItem":
		ti := &Tmall{item: item}
		go ti.Item()
		break
	case "TaobaoItem":
		ti := &Taobao{item: item}
		go ti.Item()
		break
	case "MmbItem":
		ti := &MMB{item: item}
		go ti.Item()
		break
	case "TmallShop":
		ti := &Tmall{item: item}
		go ti.Shop()
		break
	case "Other":
		ti := &Other{item: item}
		go ti.Get()
		break
	}
	return
}

func (spider *Spider) Error(item *Item) {
	if item.err != nil {
		loger.Println("[ERROR]", item.url, item.err.Error())
		content := fmt.Sprintf("%s %s", item.url, item.err.Error())
		url := fmt.Sprintf("%s?mobile=13809886150&content=%s&token=8153fa24b617b0165740211f4965dd2f", smsUrl, content)

		loader := NewLoader(url, "Get")
		_, err := loader.Send(nil)
		if err != nil {
			loger.Panicln("[ERROR] send sms error.")
		}
		item.err = nil
	}
	return
}

func (spider *Spider) Finish(item *Item) {
	loger.Println("[SUCCESS]", item.url)
	output, err := json.Marshal(item.data)
	if err != nil {
		loger.Println("error with json output")
		return
	}
	v := url.Values{}
	v.Add("id", item.id)
	v.Add("data", fmt.Sprintf("%s", output))

	loader := NewLoader(item.callback, "Post")
	_, err = loader.Send(v)
	if err != nil {
		loger.Println("callback with", item.tag, item.id)
	}
	return
}

func (spider *Spider) Daemon() {
	go spider.Listen()
	for {
		time.Sleep(time.Second * 5)
	}
}

func (spider *Spider) Add(tag, id, callback string) {
	item := &Item{
		tag:      tag,
		id:       id,
		callback: callback,
		data:     make(map[string]string),
		err:      nil,
	}
	spider.qstart <- item
}

func (spider *Spider) Listen() {
	for {
		select {
		case item := <-spider.qstart:
			go spider.Do(item)
			break
		case item := <-spider.qfinish:
			go spider.Finish(item)
			break
		case item := <-spider.qerror:
			go spider.Error(item)
			break
		}
	}
}
