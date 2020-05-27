package servers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

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
func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data ...interface{}) {
	render.Status(r, statusCode)
	render.JSON(w, r, &Response{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
		Data:    dataField(data...),
	})
}

// JSONCustom http.StatusBadRequest式应答,自定义code,提供给前端
func JSONCustom(w http.ResponseWriter, r *http.Request, code Code, data ...interface{}) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, &Response{
		Code:    code.Value(),
		Message: code.String(),
		Data:    dataField(data...),
	})
}

// JSONDetail http.StatusBadRequest式应答,含detail字段,用于debug
func JSONDetail(w http.ResponseWriter, r *http.Request, err error, data ...interface{}) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, &Response{
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		err.Error(),
		dataField(data...),
	})
}
