package spider

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Server *Spider
	loger  *log.Logger = log.New(os.Stdout, "[spider] ", log.Ldate|log.Ltime)
)

type Spider struct {
	qstart  chan *Item
	qfinish chan *Item
	qerror  chan *Item
}

type Item struct {
	id    int
	url   string
	title string
	price float64
	img   string
	tag   string
	err string
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
	}

}

func (spider *Spider) Error(item *Item) {
	loger.Println("error with", item.tag," ", item.err)
}

func (spider *Spider) Finish(item *Item) {
	loger.Println("finished with", item.url)
	fmt.Println(item)
}

func (spider *Spider) Daemon() {
	go spider.Listen()
	for {
		time.Sleep(time.Second * 5)
	}
}

func (spider *Spider) Add(tag string, id int) {
	item := &Item{
		tag: tag,
		id:  id,
	}
	spider.qstart <- item
}

func (spider *Spider) Listen() {
	for {
		select {
		case item := <-spider.qstart:
			spider.Do(item)
			break
		case item := <-spider.qfinish:
			spider.Finish(item)
			break
		case item := <-spider.qerror:
			spider.Error(item)
			break
		}
	}
}
