package handle

import (
	"EasyUse/server/logic"

	"github.com/gin-gonic/gin"
	respx "github.com/liuxiaobopro/gobox/resp"
	"github.com/zeromicro/go-zero/core/logx"
)

// UploadBook 上传书籍
func (th *handle) UploadBook(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logx.Error(err)
		th.ReturnErr(c, respx.ParamErrT)
		return
	}
	data, err1 := logic.Logic.UploadBook(c, file)
	if err != nil {
		logx.Error(err1)
		th.ReturnErr(c, err1)
		return
	}
	th.RetuenOk(c, data)
}
