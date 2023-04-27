package router

import (
	"EasyUse/server/handle"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/", handle.Handle.Hello)                 // hello
	r.POST("/uploadBook", handle.Handle.UploadBook) // 上传书籍
}
