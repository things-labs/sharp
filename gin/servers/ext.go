package servers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Attachment(c *gin.Context, filename string, data []byte) {
	c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=\"%s\"", filename))
	c.Data(http.StatusOK, "application/octet-stream;charset=utf-8", data)
}
