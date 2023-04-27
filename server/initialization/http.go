package initialization

import (
	"EasyUse/server/router"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

func InitHttp() {
	r := gin.Default()
	router.Init(r)

	logx.Info("http server start: http://localhost:50001")
	if err := r.Run(":50001"); err != nil {
		panic(err)
	}
}
