package servers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Code custom code interface
type Code interface {
	fmt.Stringer
	Value() int
}

// Response 错误信息回复基本格式
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

// JSON 标准http status code应答
func JSON(c *gin.Context, code int, data ...interface{}) {
	c.JSON(code, &Response{
		Code:    code,
		Message: http.StatusText(code),
		Data:    dataField(data...),
	})
}

// JSONCustom http.StatusBadRequest式应答,自定义code码应答,一般给前端判断使用
func JSONCustom(c *gin.Context, code Code, data ...interface{}) {
	c.JSON(http.StatusBadRequest, &Response{
		Code:    code.Value(),
		Message: code.String(),
		Data:    dataField(data...),
	})
}

// JSONDetail http.StatusBadRequest式应答,含detail字段,调试使用
func JSONDetail(c *gin.Context, err error, data ...interface{}) {
	c.JSON(http.StatusBadRequest, &Response{
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		err.Error(),
		dataField(data...),
	})
}
