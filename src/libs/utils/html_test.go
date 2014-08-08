package utils

import (
	"fmt"
	"testing"
)

func Test_LoadUrl(t *testing.T) {
	hp := NewHtmlParse()
	err := hp.LoadUrl("http://qudao.ebinf.com/mmb-union/cpsStat.jsp?uuuu=jinliquan&pppp=maiM_17Gou&startDat=2014-07-10")
	if err != nil {
		fmt.Println(err)
	}
	// hp.Clear()
	// fmt.Println(hp.content)
	m := hp.FindByTagName("tr")

	fmt.Println(len(m[1][2]))
	fmt.Println(m[1][2])
	// fmt.Println(m)
}

func Test_FindByAttr(t *testing.T) {
	hp := NewHtmlParse()
	html := "<div><tr class='a'><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr></div>"
	hp.LoadStr(html)
	m := hp.FindByAttr("tr", "class", "a")
	fmt.Println(len(m[:1][0][2]))
	fmt.Println(m[:1][0][2])
}

func Test_FindByAttr1(t *testing.T) {
	hp := NewHtmlParse()
	html := "<tr id='a' class=aaa><td>aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>"
	hp.LoadStr(html)
	m := hp.FindByAttr("tr", "id", "a")
	fmt.Println(m)
}

func Test_FindByAttr2(t *testing.T) {
	hp := NewHtmlParse()
	html := `<tr><td class="a" a="rr">aaaaa</td></tr><tr><td>bbbbbbbb</td></tr>`
	hp.LoadStr(html)
	m := hp.FindByAttr("td", "class", "a")
	fmt.Println(m)
}