package spider

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Loader struct {
	client  *http.Client
	req     *http.Request
	rheader http.Header
	url     string
	method  string
	mheader map[string]string
}

func NewLoader(url, method string) *Loader {
	return &Loader{
		client: &http.Client{},
		url:    url,
		method: method,
		mheader: map[string]string{
			"User-Agent":   "Mozilla/5.0 (Linux; U; Android 2.4; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
			"Content-Type": "application/x-www-form-urlencoded",
		},
	}
}

func (l *Loader) Send(v url.Values) ([]byte, error) {
	req, _ := http.NewRequest(strings.ToUpper(l.method), l.url, strings.NewReader(v.Encode()))

	//set headers
	l.header(req)
	req.Header.Add("Content-Length", strconv.Itoa(len(v.Encode())))

	resp, err := l.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	l.rheader = resp.Header
	return body, nil
}

func (l *Loader) GetHeader() http.Header {
	return l.rheader
}

func (l *Loader) SetHeader(key, value string) {
	l.mheader[key] = value
}

func (l *Loader) header(req *http.Request) {
	for h, v := range l.mheader {
		req.Header.Set(h, v)
	}
}
