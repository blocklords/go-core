package notify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/blocklords/go-core/fn"
	"github.com/go-resty/resty/v2"
	"net/url"
)

type (
	MarkDownModel struct {
		Title string `json:"title,omitempty"`
		Text  string `json:"text,omitempty"`
	}
	DingRequest struct {
		MsgType  string        `json:"msgtype,omitempty"`
		Markdown MarkDownModel `json:"markdown,omitempty"`
	}

	Ding struct {
		hookUrl string
		token   string
		secret  string
		body    DingRequest
	}

	DingFn func(d *Ding)
)

func DingHook(url string) DingFn {
	return func(d *Ding) {
		d.hookUrl = url
	}
}
func DingToken(token string) DingFn {
	return func(d *Ding) {
		d.token = token
	}
}
func DingSecret(secret string) DingFn {
	return func(d *Ding) {
		d.secret = secret
	}
}
func DingBody(body DingRequest) DingFn {
	return func(d *Ding) {
		d.body = body
	}
}

func NewDing(fns ...DingFn) *Ding {
	d := &Ding{}
	for _, fn := range fns {
		fn(d)
	}
	return d
}

func (d *Ding) Notify() (*resty.Response, error) {
	timestamp := fn.NowMs()

	// sign
	strToHash := fmt.Sprintf("%d\n%s", timestamp, d.secret)
	hmac256 := hmac.New(sha256.New, []byte(d.secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)

	return resty.New().R().SetHeader("Content-type", "application/json").
		SetBody(d.body).
		Post(fmt.Sprintf("%s?access_token=%s&timestamp=%d&sign=%s", d.hookUrl, d.token, timestamp, url.QueryEscape(base64.StdEncoding.EncodeToString(data))))
}
