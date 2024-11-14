package notify

import "github.com/go-resty/resty/v2"

type (
	INotify interface {
		Notify() (*resty.Response, error)
	}
)
