package notify

import "github.com/go-resty/resty/v2"

type (
	MarkDownModel struct {
		Title string `json:"title,omitempty"`
		Text  string `json:"text,omitempty"`
	}
	DingBody struct {
		MsgType  string        `json:"msgtype,omitempty"`
		Markdown MarkDownModel `json:"markdown,omitempty"`
	}
)

type (
	Body interface {
		DingBody
	}
	INotify[B Body] interface {
		Notify(body DingBody) (*resty.Response, error)
	}
)
