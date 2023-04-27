package server

import (
	"EasyUse/server/initialization"
)

func Start() {
	initialization.InitHttp()
}
