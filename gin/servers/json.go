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
	Data    interface{} `json:"data"`
}

func dataField(data ...interface{}) interface{} {
	if len(data) > 0 {
		return data[0]
	}
	return "{}"
}

// JSONs 标准http status code应答
func JSONs(c *gin.Context, httpCode int, code Code, data ...interface{}) {
	c.JSON(httpCode, &Response{
		code.Value(),
		code.String(),
		dataField(data...),
	})
}

// JSON 标准http status code应答
func JSON(c *gin.Context, httpCode int, data ...interface{}) {
	c.JSON(httpCode, &Response{
		httpCode,
		http.StatusText(httpCode),
		dataField(data...),
	})
}

// JSONCustom http.StatusBadRequest式应答,自定义code码应答,一般给前端判断使用
func JSONCustom(c *gin.Context, code Code, data ...interface{}) {
	c.JSON(http.StatusBadRequest, &Response{
		code.Value(),
		code.String(),
		dataField(data...),
	})
}

// JSONDetail http.StatusBadRequest式应答,message为err的stringer
func JSONDetail(c *gin.Context, err error, data ...interface{}) {
	c.JSON(http.StatusBadRequest, &Response{
		http.StatusBadRequest,
		err.Error(),
		dataField(data...),
	})
}
