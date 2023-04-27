package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func Start() {
	http.HandleFunc("/", greet)
	logx.Info("Starting server: http://localhost:50001")
	if err := http.ListenAndServe(":50001", nil); err != nil {
		panic(err)
	}
}
