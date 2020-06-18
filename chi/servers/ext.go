package servers

import (
	"fmt"
	"net/http"

	"github.com/thinkgos/render"
)

func Attachment(w http.ResponseWriter, filename string, data []byte) {
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=\"%s\"", filename))
	render.Data(w, http.StatusOK, "application/octet-stream;charset=utf-8", data)
}
