package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type HtmlParse struct {
	url      string
	content  string
	replaces [][]string
}

func NewHtmlParse() *HtmlParse {
	return &HtmlParse{
		replaces: [][]string{
			{`\s+`, " "},                              //过滤多余回车
			{`<[ ]+`, "<"},                            //过滤<__("<"号后面带空格)
			{`<\!–.*?–>`, ""},                         // //注释
			{`<(\!.*?)>`, ""},                         //过滤DOCTYPE
			{`<(\/?html.*?)>`, ""},                    //过滤html标签
			{`<(\/?br.*?)>`, ""},                      //过滤br标签
			{`<(\/?head.*?)>`, ""},                    //过滤head标签
			{`<(\/?meta.*?)>`, ""},                    //过滤meta标签
			{`<(\/?body.*?)>`, ""},                    //过滤body标签
			{`<(\/?link.*?)>`, ""},                    //过滤link标签
			{`<(\/?form.*?)>`, ""},                    //过滤form标签
			{`<(applet.*?)>(.*?)<(\/applet.*?)>`, ""}, //过滤applet标签
			{`<(\/?applet.*?)>`, ""},
			{`<(style.*?)>(.*?)<(\/style.*?)>`, ""}, //过滤style标签
			{`<(\/?style.*?)>`, ""},
			{`<(title.*?)>(.*?)<(\/title.*?)>`, ""}, //过滤title标签
			{`<(\/?title.*?)>`, ""},
			{`<(object.*?)>(.*?)<(\/object.*?)>`, ""}, //过滤object标签
			{`<(\/?objec.*?)>`, ""},
			{`<(noframes.*?)>(.*?)<(\/noframes.*?)>`, ""}, //过滤noframes标签
			{`<(\/?noframes.*?)>`, ""},
			{`<(i?frame.*?)>(.*?)<(\/i?frame.*?)>`, ""}, //过滤frame标签
			{`<(script.*?)>(.*?)<(\/script.*?)>`, ""},   //过滤script标签
			{`<(\/?script.*?)>`, ""},
			{`<(noscript.*?)>(.*?)<(\/noscript.*?)>`, ""}, //过滤noframes标签
			{`on([a-z]+)\s*="(.*?)"`, ""},                 //过滤dom事件
			{`on([a-z]+)\s*='(.*?)'`, ""},
		},
	}
}

func (hp *HtmlParse) LoadUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	hp.content = fmt.Sprintf("%s", body)
	hp.Clear()
	return nil
}

func (hp *HtmlParse) LoadStr(content string) {
	hp.content = content
}

func (hp *HtmlParse) Clear() {
	length := len(hp.replaces)
	for i := 0; i < length; i++ {
		if l := len(hp.replaces[i]); l > 0 {
			p, r := hp.replaces[i][:1], hp.replaces[i][1:2]
			hp.content = regexp.MustCompile(p[0]).ReplaceAllString(hp.content, r[0])
		}
	}
}

func (hp *HtmlParse) FindByTagName(tagName string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s+.*>(.*)</%s>).*?`, tagName, tagName))
	// fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}

func (hp *HtmlParse) FindByAttr(tagName, attr, value string) [][]string {
	re := regexp.MustCompile(fmt.Sprintf(`((?U)<%s+.*%s=['"]%s['"]+.*>(.*)</%s>).*?`, tagName, attr, value, tagName))
	//fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}
