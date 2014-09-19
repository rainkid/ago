package spider

import (
	"errors"
	"fmt"
)

type Other struct {
	item    *Item
	content string
}

func (ti *Other) Get() {
	//get content
	loader := NewLoader(ti.item.id, "Get")
	content, err := loader.Send(nil)
	ti.item.err = err
	if ti.CheckError() {
		return
	}

	hp := NewHtmlParse()
	hp = hp.LoadData(fmt.Sprintf("%s", content))
	ti.content = hp.content
	// ti.content = fmt.Sprintf("%s", content)
	//get title and check
	if ti.GetOtherTitle().CheckError() {
		return
	}

	Server.qfinish <- ti.item
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
