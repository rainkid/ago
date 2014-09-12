package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Server      *Spider
	loger       *log.Logger = log.New(os.Stdout, "[SPIDER] ", log.Ldate|log.Ltime)
	smsUrl      string      = "http://gou.3gtest.gionee.com/api/wifi/sms"
	callbackUrl string      = "http://gou.3gtest.gionee.com/api/apk/"
)

type Spider struct {
	qstart  chan *Item
	qfinish chan *Item
	qerror  chan *Item
}

type Item struct {
	id   string
	url  string
	data map[string]string
	tag  string
	err  error
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
	case "tmall":
		ti := &TmallItem{item: item}
		go ti.Start()
		break
	case "taobao":
		ti := &TaobaoItem{item: item}
		go ti.Start()
		break
	case "mmb":
		ti := &MMBItem{item: item}
		go ti.Start()
		break
	case "shop":
		ti := &Shop{item: item}
		go ti.Start()
		break
	}
	return
}

func (spider *Spider) Error(item *Item) {
	if item.err != nil {
		loger.Println("[ERROR]", item.url, item.err.Error())
		content := fmt.Sprintf("%s %s", item.url, item.err.Error())
		url := fmt.Sprintf("%s?mobile=13809886150&content=%s&token=8153fa24b617b0165740211f4965dd2f", smsUrl, content)
		_, err := http.Get(url)
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
		fmt.Println("error with json output")
		return
	}
	fmt.Println(fmt.Sprintf("%s", output))
	return
}

func (spider *Spider) Daemon() {
	go spider.Listen()
	for {
		time.Sleep(time.Second * 5)
	}
}

func (spider *Spider) Add(tag, id string) {
	item := &Item{
		tag:  tag,
		id:   id,
		data: make(map[string]string),
		err:  nil,
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
