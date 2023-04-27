package handle

import (
	"github.com/gin-gonic/gin"
	httpx "github.com/liuxiaobopro/gobox/http"
)

type handle struct {
	httpx.GinHandle
}

var Handle = new(handle)

func (th *handle) Hello(c *gin.Context) {
	th.RetuenOk(c, "Hello World")
}
