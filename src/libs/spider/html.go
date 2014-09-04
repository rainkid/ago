package spider

import (
	"fmt"
	iconv "github.com/djimenez/iconv-go"
	"regexp"
)

type HtmlParse struct {
	url      string
	content  string
	partten  string
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
			{`<(i?frame.*?)>(.*?)<(\/i?frame.*?)>`, ""},   //过滤frame标签
			{`<(noscript.*?)>(.*?)<(\/noscript.*?)>`, ""}, //过滤noframes标签
			{`on([a-z]+)\s*="(.*?)"`, ""},                 //过滤dom事件
			{`on([a-z]+)\s*='(.*?)'`, ""},
		},
	}
}

func (hp *HtmlParse) CleanScript() *HtmlParse {
	hp.replaces = append(hp.replaces, []string{`<(script.*?)>(.*?)<(\/script.*?)>`, ""})
	hp.replaces = append(hp.replaces, []string{`<(\/?script.*?)>`, ""})
	return hp
}

func (hp *HtmlParse) Convert(from, to string) *HtmlParse {
	converter, _ := iconv.NewConverter(from, to)
	hp.content, _ = converter.ConvertString(hp.content)
	defer converter.Close()
	return hp
}

func (hp *HtmlParse) LoadData(content string) *HtmlParse {
	hp.content = content
	return hp
}

func (hp *HtmlParse) Replace() *HtmlParse {
	length := len(hp.replaces)
	for i := 0; i < length; i++ {
		if l := len(hp.replaces[i]); l > 0 {
			p, r := hp.replaces[i][:1], hp.replaces[i][1:2]
			hp.content = regexp.MustCompile(p[0]).ReplaceAllString(hp.content, r[0])
		}
	}
	return hp
}

func (hp *HtmlParse) Partten(p string) *HtmlParse {
	hp.partten = p
	return hp
}

func (hp *HtmlParse) FindStringSubmatch() []string {
	re := regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return re.FindStringSubmatch(hp.content)
}

func (hp *HtmlParse) FindAllSubmatch() [][][]byte {
	re := regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return re.FindAllSubmatch([]byte(hp.content), -1)
}

func (hp *HtmlParse) FindByAttr(tagName, attr, value string) [][]string {
	hp.partten = fmt.Sprintf(`((?U)<%s+.*%s=['"]%s['"]+.*>(.*)</%s>).*?`, tagName, attr, value, tagName)
	re := regexp.MustCompile(hp.partten)
	//fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}

func (hp *HtmlParse) FindByTagName(tagName string) [][]string {
	hp.partten = fmt.Sprintf(`((?U)<%s+.*>(.*)</%s>).*?`, tagName, tagName)
	re := regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}

func (hp *HtmlParse) FindJson(nodeName string) [][]string {
	hp.partten = fmt.Sprintf(`(?U)"%s":"(.*)"`, nodeName)
	re := regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return re.FindAllStringSubmatch(hp.content, -1)
}
