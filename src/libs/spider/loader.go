package spider

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Loader struct {
	client  *http.Client
	url     string
	method  string
	mheader [][]string
}

func NewLoader(url, method string) *Loader {
	return &Loader{
		client: &http.Client{},
		url:    url,
		method: method,
		mheader: [][]string{
			{"User-Agent", "Mozilla/5.0 (Linux; U; Android 2.4; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"},
			{"Content-Type", "text/html;charset=GBK"},
		},
	}
}

func (l *Loader) Get() ([]byte, error) {
	req, _ := http.NewRequest(strings.ToUpper(l.method), l.url, nil)

	//set headers
	length := len(l.mheader)
	for i := 0; i < length; i++ {
		h, v := l.mheader[i][:1], l.mheader[i][1:2]
		req.Header.Set(h[0], v[0])
	}
	resp, err := l.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (l *Loader) SetHeader(key, value string) {
	l.mheader = append(l.mheader, []string{key, value})
}

func (l *Loader) header(req *http.Request) {
	length := len(l.mheader)
	for i := 0; i < length; i++ {
		h, v := l.mheader[i][:1], l.mheader[i][1:2]
		req.Header.Set(h[0], v[0])
	}
}
