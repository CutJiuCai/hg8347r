package hg8347r

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// escape transform hex digits to character
func escape(r string) string {
	p := regexp.MustCompile(`\\x\w{2}`)
	s := p.ReplaceAllStringFunc(r, func(match string) string {
		m := strings.Replace(match, "\\", "0", 1)
		i, _ := strconv.ParseInt(m, 0, 0)
		return string(i)
	})
	return s
}

// req a simple HTTP client
// mainly inspired by [imroc/req](https://github.com/imroc/req)
type req struct {
	Client *http.Client
}

func newReq() *req {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	r := &req{Client: client}
	return r
}

type response struct {
	Resp *http.Response
}

func (r *response) String() string {
	b, _ := ioutil.ReadAll(r.Resp.Body)
	r.Resp.Body.Close()
	return string(b)
}

func (r *response) EscapeString() string {
	return escape(r.String())
}

func (r *req) Do(method string, rawURL string, vs ...interface{}) *response {
	request, _ := http.NewRequest(method, rawURL, nil)
	for _, v := range vs {
		switch vv := v.(type) {
		case http.Header:
			// set header
			for key, values := range vv {
				for _, value := range values {
					request.Header.Add(key, value)
				}
			}
		case url.Values:
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
			data := []byte(vv.Encode())
			request.Body = ioutil.NopCloser(bytes.NewReader(data))
			// very important to HG8347R
			request.ContentLength = int64(len(data))
		}
	}
	resp, _ := r.Client.Do(request)
	response := &response{Resp: resp}
	return response
}

func (r *req) Get(url string, v ...interface{}) *response {
	return r.Do("GET", url, v...)
}

func (r *req) Post(url string, v ...interface{}) *response {
	return r.Do("POST", url, v...)
}
