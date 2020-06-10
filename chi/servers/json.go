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

// Response 回复基本格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Detail  string      `json:"detail,omitempty"`
	Data    interface{} `json:"data"`
}

func dataField(data ...interface{}) interface{} {
	if len(data) > 0 {
		return data[0]
	}
	return "{}"
}

// JSON 返回json信息,带标准回复
func JSON(w http.ResponseWriter, statusCode int, data ...interface{}) {
	render.JSON(w, statusCode, &Response{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
		Data:    dataField(data...),
	})
}

// JSONCustom http.StatusBadRequest式应答,自定义code,提供给前端
func JSONCustom(w http.ResponseWriter, code Code, data ...interface{}) {
	render.JSON(w, http.StatusBadRequest, &Response{
		Code:    code.Value(),
		Message: code.String(),
		Data:    dataField(data...),
	})
}

// JSONDetail http.StatusBadRequest式应答,含detail字段,用于debug
func JSONDetail(w http.ResponseWriter, err error, data ...interface{}) {
	render.JSON(w, http.StatusBadRequest, &Response{
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		err.Error(),
		dataField(data...),
	})
}
