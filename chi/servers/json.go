package servers

import (
	"fmt"
	"net/http"

	"github.com/thinkgos/render"
)

// Code code interface
type Code interface {
	fmt.Stringer
	Value() int
}

// Response 信息回复基本格式
type Response struct {
	Code   int         `json:"code"`             // 码
	Msg    string      `json:"msg,omitempty"`    // 信息
	Detail string      `json:"detail,omitempty"` // 错误携带的信息, 用于开发者调试
	Data   interface{} `json:"data"`             // 数据域
}

// Option option
type Option func(r *Response)

// WithData data
func WithData(data interface{}) Option {
	return func(r *Response) {
		r.Data = data
	}
}

// WithCode code
func WithCode(code int) Option {
	return func(r *Response) {
		r.Code = code
	}
}

// WithMsg message
func WithMsg(msg string) Option {
	return func(r *Response) {
		r.Msg = msg
	}
}

// WithICode Code interface 使应答修改code和msg,用于显示
func WithICode(code Code) Option {
	return func(r *Response) {
		r.Code = code.Value()
		r.Msg = code.String()
	}
}

// WithDetail detail 开发调试使用
func WithDetail(detail string) Option {
	return func(r *Response) {
		r.Detail = detail
	}
}

// WithError err detail为err的stringer
func WithError(err error) Option {
	return func(r *Response) {
		r.Detail = err.Error()
	}
}

// JSON 返回json信息,带标准回复
func JSON(w http.ResponseWriter, httpCode int, opts ...Option) {
	rsp := Response{
		Code: httpCode,
		Msg:  http.StatusText(httpCode),
		Data: "{}",
	}

	for _, opt := range opts {
		opt(&rsp)
	}
	render.JSON(w, httpCode, rsp)
}

// JSONCustom http.StatusBadRequest式应答,自定义code,提供给前端
func JSONCustom(w http.ResponseWriter, code Code, opts ...Option) {
	rsp := Response{
		Code: code.Value(),
		Msg:  code.String(),
		Data: "{}",
	}
	for _, opt := range opts {
		opt(&rsp)
	}
	render.JSON(w, http.StatusBadRequest, rsp)
}

// JSONDetail http.StatusBadRequest式应答,含detail字段,用于debug
func JSONDetail(w http.ResponseWriter, err error, opts ...Option) {
	rsp := Response{
		Code:   http.StatusBadRequest,
		Msg:    http.StatusText(http.StatusBadRequest),
		Data:   "{}",
		Detail: err.Error(),
	}
	for _, opt := range opts {
		opt(&rsp)
	}
	render.JSON(w, http.StatusBadRequest, rsp)
}
