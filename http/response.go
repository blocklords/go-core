package http

import (
	"github.com/kataras/iris/v12"
	"strings"
)

type (
	HttpStatus string
)

var (
	HttpStatusOk    HttpStatus = "ok"
	HttpStatusWarn  HttpStatus = "warn"
	HttpStatusError HttpStatus = "error"
)

type (
	HttpResponse[T any] struct {
		Status HttpStatus `json:"status"` // ok/warn
		Error  string     `json:"error"`  // 具体的报错信息，用于开发排查错误
		Body   T          `json:"body"`   // ok 时返回数据， warn 返回错误
	}

	HRFn[T any] func(hr *HttpResponse[T])
)

func DefaultResponse[T any]() *HttpResponse[T] {
	return &HttpResponse[T]{
		Status: HttpStatusOk,
		Error:  "",
	}
}

// 链式操作
func (response *HttpResponse[T]) WithStatus(status HttpStatus) *HttpResponse[T] {
	response.Status = status
	return response
}

func (response *HttpResponse[T]) WithError(err string) *HttpResponse[T] {
	response.Error = err
	return response
}

func (response *HttpResponse[T]) WithBody(body T) *HttpResponse[T] {
	response.Body = body
	return response
}

// 方法注入
func ResponseStatus[T any](status HttpStatus) HRFn[T] {
	return func(hr *HttpResponse[T]) {
		hr.Status = status
	}
}

func ResponseError[T any](error string) HRFn[T] {
	return func(hr *HttpResponse[T]) {
		hr.Error = error
	}
}

func ResponseBody[T any](body T) HRFn[T] {
	return func(hr *HttpResponse[T]) {
		hr.Body = body
	}
}

func NewHttpResponse[T any](fns ...HRFn[T]) *HttpResponse[T] {
	response := DefaultResponse[T]()
	for _, fn := range fns {
		fn(response)
	}
	return response
}

// 定义集中状态类型
func HttpOk[T any](body T) *HttpResponse[T] {
	return NewHttpResponse[T](
		ResponseStatus[T](HttpStatusOk),
		ResponseError[T](""),
		ResponseBody[T](body),
	)
}

func HttpWarn[T any](title T, err ...string) *HttpResponse[T] {
	return NewHttpResponse[T](
		ResponseStatus[T](HttpStatusWarn),
		ResponseBody[T](title),
		ResponseError[T](strings.Join(err, ";")),
	)
}

func HttpError[T any](title T, err ...string) *HttpResponse[T] {
	return NewHttpResponse[T](
		ResponseStatus[T](HttpStatusError),
		ResponseBody[T](title),
		ResponseError[T](strings.Join(err, ";")),
	)
}

func Response[T any](ctx iris.Context, response *HttpResponse[T]) {
	ctx.JSON(response)
}
