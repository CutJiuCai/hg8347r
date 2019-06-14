package hg8347r

import (
	"net/http"
	"net/url"
	"testing"
)

func TestEscape(t *testing.T) {
	raw := `\x20\x7b\x7d\x5b\x2e\x5f\x22\x3a\x2cA`
	if escape(raw) != " {}[._\":,A" {
		t.Error("escape is broken")
	}
}

func TestReqGet(t *testing.T) {
	client := newReq()
	header := http.Header{"name": {"hansnow"}}
	s := client.Get("https://httpbin.org/get", header).String()
	t.Log(s)
}

func TestReqPost(t *testing.T) {
	client := newReq()
	body := url.Values{}
	body.Add("key", "value")
	body.Add("empty", "")
	s := client.Post("https://httpbin.org/post", body).String()
	t.Log(s)
}

func TestClientCookie(t *testing.T) {
	client := newReq()
	client.Get("https://httpbin.org/cookies/set/name/hansnow")
	s := client.Get("https://httpbin.org/cookies")
	t.Log(s)
}
