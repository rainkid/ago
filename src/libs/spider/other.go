package spider

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

type Other struct {
	item    *Item
	content []byte
}

func (ti *Other) Get() {
	//get content
	ti.item.url = ti.item.id
	fmt.Println(ti.item.url)

	var content []byte
	var err error

	loader := NewLoader(ti.item.id, "Get")
	content, err = loader.Send(nil)

	/*if bytes.Index([]byte(ti.item.url), []byte("m.jd.com")) > 0 {
		content, err = ti.GetContent(ti.item.url)
	} else {
		content, err = loader.Send(nil)
	}*/

	if loader.redirects > 20 {
		content, err = ti.GetContent(ti.item.url)
	}

	ti.item.err = err
	if ti.CheckError() {
		return
	}

	hp := NewHtmlParse()

	hp = hp.LoadData(content).CleanScript().Replace()
	ct := []byte(loader.rheader.Get("Content-Type"))
	ct = bytes.ToLower(ct)

	var needconv bool = true
	if bytes.Index(ct, []byte("utf-8")) > 0 {
		needconv = false
	}

	if bytes.Index(ct, []byte("gbk")) > 0 {
		hp.Convert()
		needconv = false
	}
	if needconv && hp.IsGbk() {
		hp.Convert()
	}

	ti.content = hp.content

	//get title and check
	if ti.GetOtherTitle().CheckError() {
		return
	}
	// fmt.Println(ti.item.data)
	Server.qfinish <- ti.item
}

func (ti *Other) GetContent(s string) ([]byte, error) {
	url := fmt.Sprintf("http://rainkid.sinaapp.com/spider.php?url=%s", s)
	cmd := exec.Command("curl", url)
	return cmd.Output()
}

func (ti *Other) GetOtherTitle() *Other {
	hp := NewHtmlParse().LoadData(ti.content)
	title := hp.FindByTagName("title")
	if title == nil {
		ti.item.err = errors.New(`get title error`)
		return ti
	}

	ti.item.data["title"] = fmt.Sprintf("%s", title[0][2])
	return ti
}

func (ti *Other) CheckError() bool {
	if ti.item.err != nil {
		Server.qerror <- ti.item
		return true
	}
	return false
}
