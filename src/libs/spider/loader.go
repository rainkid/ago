package spider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Loader struct {
	client    *http.Client
	req       *http.Request
	resp      *http.Response
	data      url.Values
	redirects int64
	rheader   http.Header
	url       string
	method    string
	mheader   map[string]string
}

func NewLoader(url, method string) *Loader {
	return &Loader{
		redirects: 0,
		url:       url,
		method:    strings.ToUpper(method),
		mheader: map[string]string{
			"User-Agent":   "Mozilla/5.0 (Linux; Android 4.3; Nexus 7 Build/JSS15Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Safari/537.36",
			"Content-Type": "application/x-www-form-urlencoded",
		},
	}
}

func (l *Loader) CheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 20 {
		return errors.New("stopped after 20 redirects")
	}
	l.redirects++
	return nil
}

func (l *Loader) Sample() ([]byte, error) {
	resp, err := http.Get(l.url)
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

func (l *Loader) GetResp() (*http.Response, error) {
	if l.method == "POST" {
		l.req, _ = http.NewRequest(l.method, l.url, strings.NewReader(l.data.Encode()))
	} else {
		l.req, _ = http.NewRequest(l.method, l.url, nil)
	}
	l.req.Close = true

	//set headers
	l.header()
	return l.client.Do(l.req)
}

func (l *Loader) Send(v url.Values) ([]byte, error) {
	l.data = v
	l.client = &http.Client{
		CheckRedirect: l.CheckRedirect,
	}

	resp, err := l.GetResp()
	if err != nil {
		return nil, err
	}
	l.resp = resp

	defer l.resp.Body.Close()
	body, err := ioutil.ReadAll(l.resp.Body)
	if err != nil {
		return nil, err
	}
	l.rheader = l.resp.Header
	return body, nil
}

func (l *Loader) GetHeader() http.Header {
	return l.rheader
}

func (l *Loader) SetHeader(key, value string) {
	l.mheader[key] = value
}

func (l *Loader) header() {
	l.req.Close = true
	if l.method == "POST" {
		l.req.Header.Add("Content-Length", strconv.Itoa(len(l.data.Encode())))
	}
	for h, v := range l.mheader {
		l.req.Header.Set(h, v)
	}
}
